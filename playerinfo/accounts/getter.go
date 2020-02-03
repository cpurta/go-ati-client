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
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cpurta/go-ati-client/config"
	"github.com/cpurta/go-ati-client/util"
	"golang.org/x/oauth2"
)

func NewGetter(config *config.Config, httpClient *http.Client, token *oauth2.Token) Getter {
	return &defaultGetter{
		baseClient: &util.BaseClient{
			Config:     config,
			HTTPClient: httpClient,
		},
		token: token,
	}
}

type Getter interface {
	Get(playerID int, ctx context.Context) (*PlayerAccount, error)
}

var _ Getter = &defaultGetter{}

type defaultGetter struct {
	baseClient *util.BaseClient
	token      *oauth2.Token
}

func (getter *defaultGetter) Get(playerID int, ctx context.Context) (*PlayerAccount, error) {
	var (
		err error

		playerAccount = &PlayerAccount{}
		requestURL    *url.URL
		request       *http.Request
		response      *http.Response
	)

	if requestURL, err = getter.getRequestURL(playerID); err != nil {
		return nil, err
	}

	if request, err = http.NewRequest(http.MethodGet, requestURL.String(), nil); err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+getter.token.AccessToken)

	if response, err = getter.baseClient.HTTPClient.Do(request); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&playerAccount); err != nil {
		return nil, err
	}

	return playerAccount, nil
}

func (getter *defaultGetter) getRequestURL(playerID int) (*url.URL, error) {
	var (
		err        error
		endpoint   = fmt.Sprintf("%s/api/%s/PlayerInfo/%d/Accounts", getter.baseClient.Config.Host, getter.baseClient.Config.APIVersion, playerID)
		requestURL *url.URL
	)

	if requestURL, err = url.Parse(endpoint); err != nil {
		return nil, err
	}

	return requestURL, nil
}
