package playerinfo

import (
	"net/http"

	"github.com/cpurta/go-ati-client/config"
)

// NewClient returns a client that has access to player infomation.
func NewClient(config *config.Config, httpClient *http.Client) *Client {
	return &Client{
		Validator: NewValidator(config, httpClient),
	}
}

// Client allows for operations to be performed on the PlayerInfo resource of an
// ATI REST service.
type Client struct {
	Validator
	Getter
}
