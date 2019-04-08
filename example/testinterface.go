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
	"github.com/palletone/digital-identity"

)

func EnrollAdmin() error {
	err := digital_identity.InitCASDK("./", "caconfig.yaml")
	if err != nil {
		return err
	}

	enrollRequest := digital_identity.CaEnrollmentRequest{EnrollmentId: "lk", Secret: "123"}

	_, _, err = digital_identity.Enroll(digital_identity.CA, enrollRequest)
	if err != nil {
		return err

	}
	return nil
}

func Enrolluser() error {
	EnrollAdmin()
	attr := []digital_identity.CaRegisterAttribute{{
		Name: "zt",
		Value: "Hello palletone",
		ECert: true,
	},
	}
	rr := digital_identity.CARegistrationRequest{
		EnrolmentId: "05",
		Affiliation: "gptn.mediator1",
		Type: "user",
		Attrs: attr,
	}
	err := digital_identity.Register(digital_identity.CA, digital_identity.ID, &rr)

	if err != nil {
		return err
	}
	return nil

}

func Revoke(enrollmentid ,reason string) error{
	EnrollAdmin()
	req := digital_identity.CARevocationRequest{EnrollmentId:enrollmentid,Reason:reason,GenCRL:true}
	err := digital_identity.Revoke(digital_identity.CA,digital_identity.ID,&req)
	if err != nil {
		return err
	}
	return nil
}

func GetIndentity(enrollmentid ,caname string) digital_identity.CAGetIdentityResponse {
	EnrollAdmin()
	var idresp digital_identity.CAGetIdentityResponse
	idresp,err := digital_identity.GetIndentity(digital_identity.CA,digital_identity.ID,enrollmentid,caname)
	if err != nil {
		return digital_identity.CAGetIdentityResponse{}
	}
	return idresp

}

func GetIndentities() digital_identity.CAListAllIdentitesResponse {
	EnrollAdmin()
	var idresps digital_identity.CAListAllIdentitesResponse
	idresps,err := digital_identity.GetIndentities(digital_identity.CA,digital_identity.ID)
	if err != nil {
		return digital_identity.CAListAllIdentitesResponse{}
	}
	return idresps

}

func GetCaCertificateChain(caName string)(digital_identity.CAGetCertResponse,error) {
	EnrollAdmin()
	var certChain digital_identity.CAGetCertResponse
	certChain,err := digital_identity.GetCertificateChain(digital_identity.CA,digital_identity.ID,caName)
	if err != nil {
		return digital_identity.CAGetCertResponse{},err
	}
	return certChain,nil
}