package client

//func newCaGenInfo() *CaGenInfo {
//	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
//	cainfo := NewCaGenInfo("108", "lkk", "Hi palletOne", true, "user", "gptn.mediator1",key)
//	return cainfo
//}
//func TestEnrollAdmin(t *testing.T) {
//	cainfo := newCaGenInfo()
//
//	err := cainfo.EnrollAdmin()
//	if err != nil {
//		t.Log(err)
//	}
//
//}
//
//func TestEnrolluser(t *testing.T) {
//	cainfo := newCaGenInfo()
//	certpem,err := cainfo.Enrolluser()
//	if err != nil {
//		t.Log(err)
//	}
//	t.Log(certpem)
//}
//
//func TestRevoke(t *testing.T) {
//	cainfo := newCaGenInfo()
//	crlPem,err := cainfo.Revoke("102", "aacompromise")
//	if err != nil {
//		t.Log(err)
//	}
//	t.Log(crlPem)
//}
//
//func TestGetIndentity(t *testing.T) {
//	cainfo := newCaGenInfo()
//	idresp,err := cainfo.GetIndentity("104", "")
//	if err != nil {
//		t.Log(err)
//	}
//	t.Log(idresp)
//}
//
//func TestGetIndentities(t *testing.T) {
//	cainfo := newCaGenInfo()
//	idresps,err := cainfo.GetIndentities()
//	if err != nil {
//		t.Log(err)
//	}
//	t.Log(idresps)
//}
//
//func TestGetCaCertificateChain(t *testing.T) {
//	cainfo := newCaGenInfo()
//	certChain, err := cainfo.GetCaCertificateChain("")
//	if err != nil {
//		t.Log(err)
//	}
//	t.Log(certChain.RootCertificates)
//}
