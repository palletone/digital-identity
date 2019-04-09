package client

import (
	"testing"
)

func newCaGenInfo() *CaGenInfo {
	cainfo := NewCaGenInfo("10", "zk", "Hi palletOne", true, "user", "gptn.mediator1",)
	return cainfo
}
func TestEnrollAdmin(t *testing.T) {
	cainfo := newCaGenInfo()

	err := cainfo.EnrollAdmin()
	if err != nil {
		t.Log(err)
	}

}

func TestEnrolluser(t *testing.T) {

	cainfo := newCaGenInfo()
	err := cainfo.Enrolluser()
	if err != nil {
		t.Log(err)
	}
}

func TestRevoke(t *testing.T) {
	cainfo := newCaGenInfo()
	err := cainfo.Revoke("07", "aacompromise")
	if err != nil {
		t.Log(err)
	}
}

func TestGetIndentity(t *testing.T) {
	cainfo := newCaGenInfo()
	idresp := cainfo.GetIndentity("06", "ca1")
	t.Log(idresp)
}

func TestGetIndentities(t *testing.T) {
	cainfo := newCaGenInfo()
	idresps := cainfo.GetIndentities()
	t.Log(idresps)
}

func TestGetCaCertificateChain(t *testing.T) {
	cainfo := newCaGenInfo()
	certChain, err := cainfo.GetCaCertificateChain("ca1")
	if err != nil {
		t.Log(err)
	}
	t.Log(certChain)
}
