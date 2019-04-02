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

	"fmt"
)

func enrollAdmin() error {
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

func enrolluser() error {
	enrollAdmin()
	attr := []digital_identity.CaRegisterAttribute{{
		Name: "zz",
		Value: "true",
		ECert: true,
	},
	}
	rr := digital_identity.CARegistrationRequest{
		EnrolmentId: "01",
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

func revoke(enrollmentid ,reason string) error{
	enrollAdmin()
	req := digital_identity.CARevocationRequest{EnrollmentId:enrollmentid,Reason:reason,GenCRL:true}
	err := digital_identity.Revoke(digital_identity.CA,digital_identity.ID,&req)
	if err != nil {
		return err
	}
	return nil
}

func getIndentity(enrollmentid ,caname string) digital_identity.CAGetIdentityResponse {
	enrollAdmin()
	var idresp digital_identity.CAGetIdentityResponse
	idresp,err := digital_identity.GetIndentity(digital_identity.CA,digital_identity.ID,enrollmentid,caname)
	if err != nil {
		return digital_identity.CAGetIdentityResponse{}
	}
	return idresp

}

func getIndentities() digital_identity.CAListAllIdentitesResponse {
	enrollAdmin()
	var idresps digital_identity.CAListAllIdentitesResponse
	idresps,err := digital_identity.GetIndentities(digital_identity.CA,digital_identity.ID)
	if err != nil {
		return digital_identity.CAListAllIdentitesResponse{}
	}
	return idresps

}

func getCaCertificateChain(caName string)(digital_identity.CAGetCertResponse,error) {
	enrollAdmin()
	var certChain digital_identity.CAGetCertResponse
	certChain,err := digital_identity.GetCertificateChain(digital_identity.CA,digital_identity.ID,caName)
	if err != nil {
		return digital_identity.CAGetCertResponse{},err
	}
	return certChain,nil
}
func main()  {

	//err := enrollAdmin()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//err = enrolluser()
	//if err != nil {
	//	fmt.Println(err)
	//}

	idresps := getIndentity("01","palletone")
	fmt.Println(idresps)

	//certChain,err  := getCaCertificateChain("palletone")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(certChain.IntermediateCertificates)
}