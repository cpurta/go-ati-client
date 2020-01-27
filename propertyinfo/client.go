package propertyinfo

import (
	"net/http"

	"github.com/cpurta/go-ati-client/config"
)

func NewClient(config *config.Config, httpClient *http.Client) *Client {
	return &Client{
		config:     config,
		httpClient: httpClient,
	}
}

type Client struct {
	config     *config.Config
	httpClient *http.Client
}
