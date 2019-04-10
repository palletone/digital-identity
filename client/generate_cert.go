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

package client

import (
	"os"
	"fmt"
)

type CaGenInfo struct {
	EnrolmentId string `json:"enrolmentid"`
	Name        string `json:"name"`
	Data        string `json:"data"`
	ECert       bool   `json:"ecert"`
	Type        string `json:"type"`
	Affiliation string `json:"affiliation"`

}

func NewCaGenInfo(address string, name string, data string, ecert bool, ty string, affiliation string,) *CaGenInfo {
	return &CaGenInfo{
		EnrolmentId: address,
		Name:        name,
		Data:        data,
		ECert:       ecert,
		Type:        ty,
		Affiliation: affiliation,

	}
}


//You must register your administrator certificate first
func (c *CaGenInfo) EnrollAdmin() error {
	gopath := os.Getenv("GOPATH")
	path := gopath+"/src/github.com/palletone/digital-identity/config"
	fmt.Println(path)
	cacli,err := InitCASDK(path, "caconfig.yaml")
	if err != nil {
		return err
	}

	enrollRequest := CaEnrollmentRequest{EnrollmentId: cacli.Admin, Secret: cacli.Adminpw}

	_, _, err = Enroll(CA, enrollRequest)
	if err != nil {
		return err

	}
	return nil
}

func (c *CaGenInfo) Enrolluser() error {
	c.EnrollAdmin()
	attr := []CaRegisterAttribute{{
		Name:  c.Name,
		Value: c.Data,
		ECert: c.ECert,
	},
	}
	rr := CARegistrationRequest{
		EnrolmentId: c.EnrolmentId,
		Affiliation: c.Affiliation,
		Type:        c.Type,
		Attrs:       attr,
	}
	err := Register(CA, ID, &rr)

	if err != nil {
		return err
	}
	return nil

}

func (c *CaGenInfo) Revoke(enrollmentid, reason string) error {
	c.EnrollAdmin()
	req := CARevocationRequest{EnrollmentId: enrollmentid, Reason: reason, GenCRL: true}
	err := Revoke(CA, ID, &req)
	if err != nil {
		return err
	}
	return nil
}

func (c *CaGenInfo) GetIndentity(enrollmentid, caname string) *CAGetIdentityResponse {
	c.EnrollAdmin()
	var idresp CAGetIdentityResponse
	idresp, err := GetIndentity(CA, ID, enrollmentid, caname)
	if err != nil {
		return &CAGetIdentityResponse{}
	}
	return &idresp

}

func (c *CaGenInfo) GetIndentities() *CAListAllIdentitesResponse {
	c.EnrollAdmin()
	var idresps CAListAllIdentitesResponse
	idresps, err := GetIndentities(CA, ID)
	if err != nil {
		return &CAListAllIdentitesResponse{}
	}
	return &idresps

}

func (c *CaGenInfo) GetCaCertificateChain(caName string) (*CAGetCertResponse, error) {
	c.EnrollAdmin()
	var certChain CAGetCertResponse
	certChain, err := GetCertificateChain(CA, ID, caName)
	if err != nil {
		return &CAGetCertResponse{}, err
	}
	return &certChain, nil
}
