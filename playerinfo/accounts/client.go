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

package accounts

import (
	"net/http"

	"github.com/cpurta/go-ati-client/config"
	"golang.org/x/oauth2"
)

func NewClient(config *config.Config, httpClient *http.Client, token *oauth2.Token) *Client {
	return &Client{
		Getter: NewGetter(config, httpClient, token),
	}
}

type Client struct {
	Getter
}
