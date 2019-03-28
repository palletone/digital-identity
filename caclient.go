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
	"cloudflare/cfssl/csr"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"net/http"
)

type PalletCAClient struct {
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

var CA *PalletCAClient

type ServerInfo struct {
	CAName string
	CACert *x509.Certificate
}


// RegistrationRequest holds all data needed for new registration of new user in Certificate Authority
type CARegistrationRequest struct {
	// EnrolmentId is unique name that identifies identity
	EnrolmentId string `json:"id"`
	// Type defines type of this identity (user,client, auditor etc...)
	Type string `json:"type"`
	// Secret is password that will be used for enrollment. If not provided random password will be generated
	Secert string `json:"secert,omitempty" mask:"password" help:"The enrollment secret for the identity being registered"`
	// MaxEnrollments define maximum number of times that identity can enroll. If not provided or is 0 there is no limit
	MaxEnrollments int `json:"max_enrollments,omitempty"`
	// Affiliation associates identity with particular organisation.
	Affiliation string `json:"affiliation" help:"The identity's affiliation"`
	// Attrs are attributes associated with this identity
	Attrs []CaRegisterAttribute `json:"attrs,omitempty"`
	// CAName is the name of the CA that should be used.
	CAName string `json:"caname,omitempty" skip:"true"`
}

// CaRegisterAttribute holds user attribute used for registration
// for example user may have attr `accountType` with value `premium`
// this attributes can be accessed in chainCode and build business logic on top of them
type CaRegisterAttribute struct {
	// Name is the name of the attribute.
	Name string `json:"name"`
	// Value is the value of the attribute. Can be empty string
	Value string `json:"value"`
	// ECert define how this attribute will be included in ECert. If this value is true this attribute will be
	// added to ECert automatically on Enrollment if no attributes are requested on Enrollment request.
	ECert bool `json:"ecert,omitempty"`
}

type certificateRequest struct {
	CaEnrollmentRequest
	CR string `json:"certificate_request"`
}

// enrollmentResponse is response from fabric-ca server for enrolment that contains created Ecert
type enrollmentResponse struct {
	caResponse
	Result enrollmentResponseResult `json:"result"`
}

type enrollmentResponseResult struct {
	Cert       string
	ServerInfo enrollmentResponseServerInfo
	Version    string
}

type enrollmentResponseServerInfo struct {
	CAName  string
	CAChain string
}

// CAResponse represents response message from fabric-ca server
type caResponse struct {
	Success  bool            `json:"success"`
	Errors   []caResponseErr `json:"errors"`
	Messages []string        `json:"messages"`
}

type caResponseErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CaEnrollAttribute struct {
	// Name is the name of the attribute
	Name string `json:"name"`
	// Optional define behaviour when required attribute is not available to user. If `true` then request will continue,
	// but attribute will not be included in ECert. If `false` and attribute is missing, request will fail.
	// If false and attribute is available, request will continue and attribute will be added in ECert
	Optional bool `json:"optional,omitempty"`
}

// CaEnrollmentRequest holds data needed for getting ECert (enrollment) from CA server
type CaEnrollmentRequest struct {
	EnrollmentId string `json:"name" skip:"true"`
	// Secret is the password for this identity
	Secret string `json:"secret,omitempty" skip:"true" mask:"password"`
	Profile string `json:"profile,omitempty"`
	//Label is used for hardware secure modules.
	Label string `json:"label,omitempty"`
	CAName string `json:"caname,omitempty" skip:"true"`
	// Host is the list of valid host names for this certificate. If empty default hosts will be used
	Hosts []string `json:"hosts"`
	// Attrs are the attributes that must be included in ECert. This is subset of the attributes used in registration.
	Attrs []CaEnrollAttribute `json:"attr_reqs,omitempty"`
	// CSR is Certificate Signing Request info
	CSR *CSRInfo `json:"csr,omitempty" help:"Certificate Signing Request info"`
}

type CSRInfo struct {
	CN           string           `json:"CN"`
	Names        []csr.Name      `json:"names,omitempty"`
	Hosts        []string         `json:"hosts,omitempty"`
	KeyRequest   *BasicKeyRequest `json:"key,omitempty"`
	CA           *csr.CAConfig    `json:"ca,omitempty"`
	SerialNumber string           `json:"serial_number,omitempty"`
}

// BasicKeyRequest encapsulates size and algorithm for the key to be generated
type BasicKeyRequest struct {
	Algo string `json:"algo" yaml:"algo"`
	Size int    `json:"size" yaml:"size"`
}

// CARegisterCredentialResponse credentials from palletone-ca server registration request
type caRegisterCredentialResponse struct {
	Secret string `json:"secret"`
}

type caRegisterResponse struct {
	caResponse
	Result caRegisterCredentialResponse `json:"result"`
}

// CAGetCertsResponse holds response from `GetCaCertificateChain`
type CAGetCertResponse struct {
	// RootCertificates is list of pem encoded certificates
	RootCertificates []*x509.Certificate
	// IntermediateCertificates is list of pem encoded intermediate certificates
	IntermediateCertificates []*pem.Block
	// CAName is the name of the CA server that returns this certificates
	CAName string
	// Version is the version of server that returns this certificates
	Version string
}

type caInfoRequest struct {
	CaName string `json:"caname,omitempty"`
}

type caInfoResponse struct {
	caResponse
	Result caResponseResult `json:"result"`
}

type caResponseResult struct {
	CAName  string `json:"CAName"`
	CAChain string `json:"CAChain"`
	Version string `json:"Version"`
}

type CARevocationRequest struct {
	EnrollmentId string `json:"id,omitempty"`
	Serial string `json:"serial,omitempty"`
	AKI string `json:"aki,omitempty"`
	Reason string `json:"reason,omitempty"`
	CAName string `json:"caname,omitempty"`
	// GenCRL specifies whether to generate a CRL. CRL will be returned only when AKI and Serial are provided.
	GenCRL bool `json:"gencrl,omitempty"`
}

// CaRevokeResultCertificate identify revoked certificate
type CaRevokeResultCertificate struct {
	Serial string `json:"Serial"`
	AKI string `json:"AKI"`
}

type CARevokeResult struct {
	RevokedCertificates []CaRevokeResultCertificate `json:"RevokedCerts"`
	CRL string `json:"CRL"`
}

type caRevokeResponse struct {
	caResponse
	Result CARevokeResult
}

type CaIdentityResponse struct {
	ID             string                `json:"id"`
	Type           string                `json:"type"`
	Affiliation    string                `json:"affiliation"`
	Attributes     []CaRegisterAttribute `json:"attrs" mapstructure:"attrs"`
	MaxEnrollments int                   `json:"max_enrollments" mapstructure:"max_enrollments"`
}

type CAGetIdentityResponse struct {
	CaIdentityResponse
	CAName string `json:"caname"`
}

type caGetIdentity struct {
	caResponse
	Result CAGetIdentityResponse `json:"result"`
}

type CAListAllIdentitesResponse struct {
	CAName     string               `json:"caname"`
	Identities []CaIdentityResponse `json:"identities,omitempty"`
}

type caListAllIdentities struct {
	caResponse
	Result CAListAllIdentitesResponse `json:"result"`
}

func NewCAClient(path string, transport *http.Transport) (*PalletCAClient, error) {
	config, err := NewCAConfig(path)
	if err != nil {
		return nil, err
	}
	return NewCaClientFromConfig(*config, transport)
}

func concatErrors(errs []caResponseErr) error {
	errors := ""
	for _, e := range errs {
		errors += e.Message + ":"
	}
	return fmt.Errorf(errors)
}

func NewCaClientFromConfig(config CAConfig, transport *http.Transport) (*PalletCAClient, error) {
	var crypto CryptoSuite
	var err error

	switch config.CryptoConfig.Family {
	case "ecdsa":
		crypto, err = NewECCryptoSuiteFromConfig(config.CryptoConfig)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrInvalidAlgorithmFamily
	}
	CA = &PalletCAClient{SkipTLSVerification: config.SkipTLSValidation,
		Url:       config.Url,
		Crypto:    crypto,
		Transport: transport,
		MspId:     config.MspId,
		FilePath:  config.FilePath}
	return CA, nil
}