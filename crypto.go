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
// CryptSuite defines common interface for different crypto implementations.
// Currently Hyperledger Fabric supports only Elliptic curves.
type CryptoSuite interface {
	// GenerateKey returns PrivateKey.
	GenerateKey() (interface{}, error)
	// CreateCertificateRequest will create CSR request. It takes enrolmentId and Private key
	CreateCertificateRequest(enrollmentId string, key interface{}, hosts []string) ([]byte, error)
	// Sign signs message. It takes message to sign and Private key
	Sign(msg []byte, k interface{}) ([]byte, error)
	// Hash computes Hash value of provided data. Hash function will be different in different crypto implementations.
	Hash(data []byte) []byte
}