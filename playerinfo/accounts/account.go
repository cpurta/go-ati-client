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

type PlayerAccount struct {
	AccountActiveField     bool   `json:"AccountActiveField"`
	AccountExpiry          string `json:"AccountExpiry"`
	AccountExpirySpecified bool   `json:"AccountExpirySpecified"`
	AccountId              string `json:"AccountId"`
	AccountName            string `json:"AccountName"`
	AccountOwnership       bool   `json:"AccountOwnership"`
	AccountType            int    `json:"AccountType"`
	CurrencyCode           string `json:"CurrencyCode"`
	CreditType             string `json:"CreditType"`
	DepositOk              bool   `json:"DepositOk"`
	ForeignCode            string `json:"ForeignCode"`
	ForeignDefaultAmount   int    `json:"ForeignDefaultAmount"`
	ForeignDepositMax      int    `json:"ForeignDepositMax"`
	ForeignDepositMin      int    `json:"ForeignDepositMin"`
	ForeignExpiry          string `json:"ForeignExpiry"`
	ForeignExpirySpecified bool   `json:"ForeignExpirySpecified"`
	ForeignRate            int    `json:"ForeignRate"`
	ForeignWithdrawInc     int    `json:"ForeignWithdrawInc"`
	ForeignWithdrawMax     int    `json:"ForeignWithdrawMax"`
	ForeignWithdrawMin     int    `json:"ForeignWithdrawMin"`
	LocalDefaultAmount     int    `json:"LocalDefaultAmount"`
	LocalDepositMax        int    `json:"LocalDepositMax"`
	LocalDepositMin        int    `json:"LocalDepositMin"`
	LocalWithdrawInc       int    `json:"LocalWithdrawInc"`
	LocalWithdrawMax       int    `json:"LocalWithdrawMax"`
	LocalWithdrawMin       int    `json:"LocalWithdrawMin"`
	OwnershipAttributes    bool   `json:"OwnershipAttributes"`
	PinProtected           bool   `json:"PinProtected"`
	SelectAmount           bool   `json:"SelectAmount"`
	WithdrawOk             bool   `json:"WithdrawOk"`
}
