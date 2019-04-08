/*
 *
 *    This file is part of go-palletone.
 *    go-palletone is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *    go-palletone is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *    You should have received a copy of the GNU General Public License
 *    along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
 * /
 *
 *  * @author PalletOne core developer <dev@pallet.one>
 *  * @date 2018
 *
 */

package main

import (
	"testing"
	"fmt"
)

func TestEnrollAdmin(t *testing.T) {
	err := EnrollAdmin()
	if err != nil {
		t.Log(err)
	}
}

func TestEnrolluser(t *testing.T) {
	err := Enrolluser()
	if err != nil {
		fmt.Println(err)
	}
}

func TestRevoke(t *testing.T) {
	err := Revoke("01","aacompromise")
	if err != nil {
		fmt.Println(err)
	}
}
func TestGetIndentity(t *testing.T) {
	idresp := GetIndentity("05","")
	fmt.Println(idresp)
}

func TestGetIndentities(t *testing.T) {
	idresps := GetIndentities()
	fmt.Println(idresps)
}

func TestGetCaCertificateChain(t *testing.T) {
	certChain,err  := GetCaCertificateChain("ca1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(certChain)
}