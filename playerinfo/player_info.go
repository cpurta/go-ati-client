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

package playerinfo

type PlayerInfo struct {
	PlayerID              string                  `json:"PlayerID"`
	PlayerProfile         *PlayerProfile          `json:"PlayerProfile"`
	Addresses             []interface{}           `json:"Address"`
	PhoneNumbers          []interface{}           `json:"Phone"`
	Emails                []interface{}           `json:"Email"`
	Idents                []interface{}           `json:"Ident"`
	Images                []interface{}           `json:"Image"`
	Comments              []interface{}           `json:"Comment"`
	Stops                 []interface{}           `json:"Stop"`
	Links                 []interface{}           `json:"Link"`
	Groups                []interface{}           `json:"Group"`
	PlayerExternalIDs     []interface{}           `json:"PlayerExternalIDs"`
	KeyPairs              []interface{}           `json:"KeyPair"`
	Cards                 []interface{}           `json:"Cards"`
	PlayerBalances        []interface{}           `json:"PlayerBalance"`
	PlayerAccountBalances []*PlayerAccountBalance `json:"PlayerAccountBalance"`
	Interests             []interface{}           `json:"Interest"`
	PlayerPins            []interface{}           `json:"PlayerPins"`
	PlayerStatistics      []interface{}           `json:"PlayerStatistics"`
	MarketingOffers       []interface{}           `json:"marketingOffers"`
	PlayerRankHistories   []interface{}           `json:"PlayerRankHistory"`
}

type PlayerProfile struct {
	AnniversaryDate     string `json:"AnniversaryDate"`
	BirthDate           string `json:"BirthDate"`
	CtrComplete         bool   `json:"CtrComplete"`
	EyeColor            string `json:"EyeColor"`
	FirstName           string `json:"FirstName"`
	FullName            string `json:"FullName"`
	HairColor           string `json:"HairColor"`
	HeightUnitOfMeasure string `json:"HeightUom"`
	GamePlayProtected   bool   `json:"GamePlayProtected"`
	GenderType          string `json:"GenderType"`
	Generation          string `json:"Generation"`
	GeographicId        string `json:"GeographicId"`
	JunketRepId         string `json:"JunketRepId"`
	LastName1           string `json:"LastName1"`
	LastName2           string `json:"LastName2"`
	LocaleCode          string `json:"LocaleCode"`
	MaidenName          string `json:"MaidenName"`
	MaritalState        string `json:"MaritalState"`
	MarketingMessages   bool   `json:"MarketingMessages"`
	MiddleName          string `json:"MiddleName"`
	PinProtected        bool   `json:"PinProtected"`
	PlayerActive        bool   `json:"PlayerActive"`
	PlayerHeight        int    `json:"PlayerHeight"`
	PlayerPrivacy       bool   `json:"PlayerPrivacy"`
	PlayerProperty      string `json:"PlayerProperty"`
	PlayerRank          string `json:"PlayerRank"`
	PlayerVIP           bool   `json:"PlayerVIP"`
	PlayerWeight        int    `json:"PlayerWeight"`
	PreferredName       string `json:"PreferredName"`
	RankName            string `json:"RankName"`
	RegisteredDateTime  string `json:"RegisteredDateTime"`
	RegisteredEmployee  string `json:"RegisteredEmployee"`
	RegisteredProperty  string `json:"RegisteredProperty"`
	SkinColor           string `json:"SkinColor"`
	Suffix              string `json:"Suffix"`
	Title               string `json:"Title"`
	WeightUnitOfMeasure string `json:"WeightUom"`
}

type PlayerAccountBalance struct{}
