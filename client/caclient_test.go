package client

//func newCaGenInfo() *CaGenInfo {
//	cainfo := NewCaGenInfo("21", "zk", "Hi palletOne", true, "user", "gptn.mediator1",)
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
//
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
//	crlPem,err := cainfo.Revoke("07", "aacompromise")
//	if err != nil {
//		t.Log(err)
//	}
//	t.Log(crlPem)
//}
//
//func TestGetIndentity(t *testing.T) {
//	cainfo := newCaGenInfo()
//	idresp,err := cainfo.GetIndentity("06", "ca1")
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
//	certChain, err := cainfo.GetCaCertificateChain("ca1")
//	if err != nil {
//		t.Log(err)
//	}
//	t.Log(certChain)
//}
