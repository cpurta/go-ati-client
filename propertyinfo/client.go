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
