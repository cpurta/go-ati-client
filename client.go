package goaticlient

import (
	"net/http"

	"github.com/cpurta/go-ati-client/config"
	"github.com/cpurta/go-ati-client/playerinfo"
	"github.com/cpurta/go-ati-client/propertyinfo"
)

// NewClient returns an ATI client that allows for access to all sublcients that
// access an individueal resources provided by the ATI REST service.
func NewClient(config *config.Config, httpClient *http.Client) *Client {
	return &Client{
		PropertyInfoClient: propertyinfo.NewClient(config, httpClient),
		PlayerInfoClient:   playerinfo.NewClient(config, httpClient),
	}
}

type Client struct {
	PropertyInfoClient *propertyinfo.Client
	PlayerInfoClient   *playerinfo.Client
}

// PropertyInfo returns a sub-client that has access to property infomation.
func (client *Client) PropertyInfo() *propertyinfo.Client {
	return client.PropertyInfoClient
}

// PlayerInfo returns a sub-client that has access to player infomation.
func (client *Client) PlayerInfo() *playerinfo.Client {
	return client.PlayerInfoClient
}
