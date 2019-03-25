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

import (
	"crypto/x509"
	"net/http"
)

type FabricCAClient struct {
	// Uri is access point for palletone-ca server. Port number and scheme must be provided.
	// for example http://127.0.0.1:7054
	Url string
	// SkipTLSVerification define how connection must handle invalid TLC certificates.
	// if true, all verifications are skipped. This value is overwritten by Transport property, if provided
	SkipTLSVerification bool
	// Crypto is CryptSuite implementation used to sign request for palletone-ca server
	Crypto CryptoSuite
	// Transport define transport rules for communication with palletone-ca server. If nil, default Go setting will be used
	// It is responsibility of the user to provide proper TLS/certificate setting in TLS communication.
	Transport *http.Transport
	// MspId value will be added to Identity in Enrollment and ReEnrollment invocations.
	// This value is not used anywhere in CA implementation, but is need in every call to palletone and is added here
	// for convenience, because (in general case) palletoneCA is serving one MSP
	// User can overwrite this value at any time.
	MspId string

	FilePath string

	ServerInfo ServerInfo
}

var CA *FabricCAClient

type ServerInfo struct {
	CAName string
	CACert *x509.Certificate
}