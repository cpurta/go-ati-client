// Copyright (C) 2020  Chris Purta
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package goaticlient

import (
	"net/http"

	"github.com/cpurta/go-ati-client/config"
	"github.com/cpurta/go-ati-client/playerinfo"
	"github.com/cpurta/go-ati-client/propertyinfo"
	"golang.org/x/oauth2"
)

// NewClient returns an ATI client that allows for access to all sublcients that
// access an individueal resources provided by the ATI REST service.
func NewClient(config *config.Config, httpClient *http.Client, token *oauth2.Token) *Client {
	return &Client{
		PropertyInfoClient: propertyinfo.NewClient(config, httpClient),
		PlayerInfoClient:   playerinfo.NewClient(config, httpClient, token),
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
