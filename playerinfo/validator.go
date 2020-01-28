package playerinfo

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/cpurta/go-ati-client/config"
	"github.com/cpurta/go-ati-client/util"
)

// NewValidator returns an interface that is able to validate a player and their
// PIN by sending a POST request to the PlayerInfo resource of the ATI REST service.
func NewValidator(config *config.Config, httpClient *http.Client) Validator {
	return &defaultValidator{
		baseClient: &util.BaseClient{
			Config:     config,
			HTTPClient: httpClient,
		},
	}
}

var _ Validator = &defaultValidator{}

// Validator allows for validating a players PIN.
type Validator interface {
	Validate(playerID int, pin string, ctx context.Context) error
}

type defaultValidator struct {
	baseClient *util.BaseClient
}

type validatePlayerPIN struct {
	PlayerID   int         `json:"PlayerID"`
	PlayerPINs []playerPIN `json:"PlayerPins"`
}

type playerPIN struct {
	PinNumber string `json:"PinNumber"`
}

// Validate will send a POST request to the configured ATI REST service that will
// determine if the players PIN is valid. If the response returns anything other
// than HTTP 200 it will return an error stating that the players PIN in invalid.
func (validator *defaultValidator) Validate(playerID int, pin string, ctx context.Context) error {
	var (
		err error

		h               = sha1.New()
		hashedPlayerPIN string

		validatePlayerPINRequest *validatePlayerPIN

		requestBody []byte

		requestURL *url.URL
		request    *http.Request
		response   *http.Response
	)

	if requestURL, err = validator.getRequestURL(); err != nil {
		return err
	}

	io.WriteString(h, pin)

	hashedPlayerPIN = fmt.Sprintf("%x", h.Sum(nil))

	validatePlayerPINRequest = &validatePlayerPIN{
		PlayerID: playerID,
		PlayerPINs: []playerPIN{
			{
				PinNumber: hashedPlayerPIN,
			},
		},
	}

	if requestBody, err = json.Marshal(validatePlayerPINRequest); err != nil {
		return err
	}

	if request, err = http.NewRequest(http.MethodPost, requestURL.String(), strings.NewReader(string(requestBody))); err != nil {
		return err
	}

	request = request.WithContext(ctx)

	if response, err = validator.baseClient.HTTPClient.Do(request); err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to validate player PIN: %d", response.StatusCode)
	}

	return nil
}

func (validator *defaultValidator) getRequestURL() (*url.URL, error) {
	var (
		err        error
		endpoint   = fmt.Sprintf("%s/api/%s/PlayerInfo/Validation", validator.baseClient.Config.Host, validator.baseClient.Config.APIVersion)
		requestURL *url.URL
	)

	if requestURL, err = url.Parse(endpoint); err != nil {
		return nil, err
	}

	return requestURL, nil
}
