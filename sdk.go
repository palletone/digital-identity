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
package digital_identity


func InitCASDK(configPth string, configFile string) error {
	//caconfigFilePath := path.Join(configPth, configFile)
	//_, err := NewCAClient(caconfigFilePath, nil)
	//if err != nil {
	//	return err
	//}
	return nil
}
//todo sdk for caclient
func Enroll(ca *FabricCAClient, req CaEnrollmentRequest) (*Identity, []byte, error) {
	return nil, nil, nil
}

func Register(ca *FabricCAClient, identity *Identity, req *CARegistrationRequest) error {


	return nil
}

func getCaCerts(ca *FabricCAClient) (*CAGetCertResponse, error) {

	return nil, nil
}

func Revoke(ca *FabricCAClient, identity *Identity, req *CARevocationRequest) error {

	return nil
}

func GetIndentity(ca *FabricCAClient, identity *Identity, id string, caName string) (CAGetIdentityResponse, error) {


	return CAGetIdentityResponse{}, nil
}

func GetIndentities(ca *FabricCAClient, identity *Identity) (CAListAllIdentitesResponse, error) {


	return CAListAllIdentitesResponse{}, nil
}

