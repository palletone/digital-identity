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
	"path"
)

func InitCASDK(configPth string, configFile string) (*PalletCAClient,error) {
	caconfigFilePath := path.Join(configPth, configFile)

	cacli, err := NewCAClient(caconfigFilePath, nil)
	if err != nil {
		return nil,err
	}
	return cacli,nil
}

func Enroll(ca *PalletCAClient, req CaEnrollmentRequest) (*Identity, []byte, error) {
	id, csr, err := ca.Enroll(req)
	if err != nil {
		return nil, nil, err
	}
	resp, err := getCaCerts(ca)
	if req.Profile == "tls" {
		id.SaveTLScert(ca, resp)
	} else {
		id.SaveCert(ca, nil, resp)
	}
	return id, csr, nil
}

func Register(ca *PalletCAClient, identity *Identity, req *CARegistrationRequest) error {
	resp, err := ca.Register(identity, req)
	if err != nil {
		return err
	}

	enrollRequest := CaEnrollmentRequest{EnrollmentId: req.EnrolmentId, Secret: resp}
	id, _, err := ca.Enroll(enrollRequest)
	if err != nil {
		return err
	}

	cainfo, err := getCaCerts(ca)
	if err != nil {
		return err
	}
	err = id.SaveCert(ca, &enrollRequest, cainfo)
	if err != nil {
		return err
	}
	return nil
}

func getCaCerts(ca *PalletCAClient) (*CAGetCertResponse, error) {
	resp, err := ca.GetCaCertificateChain(ca.ServerInfo.CAName)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Revoke(ca *PalletCAClient, identity *Identity, req *CARevocationRequest) error {
	r, err := ca.Revoke(identity, req)
	if err != nil {
		return err
	}
	err = SaveCrl(ca, req, r)
	if err != nil {
		return err
	}
	return nil
}

func GetIndentity(ca *PalletCAClient, identity *Identity, id string, caName string) (CAGetIdentityResponse, error) {
	resp, err := ca.GetIndentity(identity, id, caName)
	if err != nil {
		return CAGetIdentityResponse{}, err
	}

	return *resp, nil
}

func GetIndentities(ca *PalletCAClient, identity *Identity) (CAListAllIdentitesResponse, error) {
	resp, err := ca.GetIdentities(identity, "")
	if err != nil {
		return CAListAllIdentitesResponse{}, nil
	}

	return *resp, nil
}

func GetCertificateChain(ca *PalletCAClient, identity *Identity,caName string) (CAGetCertResponse,error) {
	resp, err := ca.GetCaCertificateChain(caName)
	if err != nil {
		return CAGetCertResponse{},err
	}
	return *resp,nil
}