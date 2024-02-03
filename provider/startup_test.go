package provider

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
	"github.com/advanced-go/core/uri"
	"net/http"
)

func ExampleStartupPing() {
	//w := httptest.NewRecorder()
	r, _ := http.NewRequest("", "github/advanced-go/search/provider:ping", nil)
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	status := messaging.Ping(nil, nid)
	//buf, status := io2.ReadAll(w.Result().Body, nil)
	//if !status.OK() {
	//	fmt.Printf("test: NewBytes() -> [status:%v]\n", status)
	//}
	fmt.Printf("test: Ping() -> [nid:%v] [nss:%v] [ok:%v] [status:%v]\n", nid, rsc, ok, status)

	//Output:
	//test: Ping() -> [nid:github/advanced-go/search/provider] [nss:ping] [ok:true] [status:200]

}
