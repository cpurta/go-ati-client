package playerinfo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cpurta/go-ati-client/config"
	"github.com/cpurta/go-ati-client/util"
)

func NewGetter(config *config.Config, httpClient *http.Client) Getter {
	return &defaultGetter{
		baseClient: &util.BaseClient{
			Config:     config,
			HTTPClient: httpClient,
		},
	}
}

type Getter interface {
	Get(playerID int, ctx context.Context) (*PlayerInfo, error)
}

var _ Getter = &defaultGetter{}

type defaultGetter struct {
	baseClient *util.BaseClient
}

func (getter *defaultGetter) Get(playerID int, ctx context.Context) (*PlayerInfo, error) {
	var (
		err error

		playerInfo = &PlayerInfo{}
		requestURL *url.URL
		request    *http.Request
		response   *http.Response
	)

	if requestURL, err = getter.getRequestURL(playerID); err != nil {
		return nil, err
	}

	if request, err = http.NewRequest(http.MethodGet, requestURL.String(), nil); err != nil {
		return nil, err
	}

	if response, err = getter.baseClient.HTTPClient.Do(request); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&playerInfo); err != nil {
		return nil, err
	}

	return playerInfo, nil
}

func (getter *defaultGetter) getRequestURL(playerID int) (*url.URL, error) {
	var (
		err        error
		endpoint   = fmt.Sprintf("%s/api/%s/PlayerInfo/%d", getter.baseClient.Config.Host, getter.baseClient.Config.APIVersion, playerID)
		requestURL *url.URL
	)

	if requestURL, err = url.Parse(endpoint); err != nil {
		return nil, err
	}

	return requestURL, nil
}
