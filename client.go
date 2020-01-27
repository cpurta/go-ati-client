package goaticlient

import (
	"net/http"

	"github.com/cpurta/go-ati-client/config"
	"github.com/cpurta/go-ati-client/playerinfo"
	"github.com/cpurta/go-ati-client/propertyinfo"
)

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

func (client *Client) PropertyInfo() *propertyinfo.Client {
	return client.PropertyInfoClient
}

func (client *Client) PlayerInfo() *playerinfo.Client {
	return client.PlayerInfoClient
}
