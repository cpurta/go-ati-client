package playerinfo

import (
	"net/http"

	"github.com/cpurta/go-ati-client/config"
)

func NewClient(config *config.Config, httpClient *http.Client) *Client {
	return &Client{
		Validator: NewValidator(config, httpClient),
	}
}

type Client struct {
	Validator
}
