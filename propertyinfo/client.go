package propertyinfo

import (
	"net/http"

	"github.com/cpurta/go-ati-client/config"
)

// NewClient returns a client that is able to access property infomation.
func NewClient(config *config.Config, httpClient *http.Client) *Client {
	return &Client{
		config:     config,
		httpClient: httpClient,
	}
}

// Client allows for HTTP requests to be sent to the PropertyInfo resource of the
// ATI REST service.
type Client struct {
	config     *config.Config
	httpClient *http.Client
}
