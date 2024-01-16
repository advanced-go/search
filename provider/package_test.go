package provider

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
)

const (
	resultUri = "file://[cwd]/providertest/resource/query-result.txt"
)

func Example_PkgPath() {
	dotCom := ".com"
	pkgPath := reflect.TypeOf(any(pkg{})).PkgPath()
	fmt.Printf("test: PkgPath = \"%v\"\n", strings.Replace(pkgPath, dotCom, "", len(dotCom)))

	//Output:
	//test: PkgPath = "github/advanced-go/search/provider"

}

func Example_Search() {
	resolver.SetOverrides([]runtime.Pair{{searchPath, "https://www.google.com/search?q=slo"}})
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	buf, status := Search[runtime.Output](nil, req.URL.Query(), false)
	s := string(buf)
	if len(s) > 0 {
	}
	fmt.Printf("test: search(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:115289]

}

func Example_Search_Override() {
	resolver.SetOverrides([]runtime.Pair{{searchPath, resultUri}})
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	buf, status := Search[runtime.Output](nil, req.URL.Query(), false)
	//buf, _ := runtime.NewBytes(r)
	s := string(buf)
	s = s[:len(s)-2]
	fmt.Printf("test: search(%v) -> [status:%v] [content:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, s, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:OK] [content:This is an alternate result for a Google query.] [content-type:text/plain] [content-length:49]

}

func Example_HttpHandler() {
	resolver.SetOverrides([]runtime.Pair{{searchPath, resultUri}})
	rec := httptest.NewRecorder()

	HttpHandler(rec, nil)
	fmt.Printf("test: HttpHandler() -> [status-code:%v]\n", rec.Result().StatusCode)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+"invalid-path"+":search?q=golang", nil)
	rec = httptest.NewRecorder()
	HttpHandler(rec, req)
	buf, _ := runtime.New[[]byte](rec.Result())
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":searchBad?q=golang", nil)
	rec = httptest.NewRecorder()
	HttpHandler(rec, req)
	buf, _ = runtime.New[[]byte](rec.Result())
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:500]
	//test: HttpHandler() -> [status-code:400] [content:error invalid URI, NID does not match: "/invalid-path:search" "github/advanced-go/search/provider"]
	//test: HttpHandler() -> [status-code:404] [content:error invalid URI, resource was not found: [searchBad]]

}

func ExampleHttpHandler_Search() {
	resolver.SetOverrides([]runtime.Pair{{searchPath, "https://www.google.com/search?q=golang"}})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)

	HttpHandler(rec, req)
	//_, _ := runtime.New[[]byte](rec.Result())
	fmt.Printf("test: HttpHandler() -> [status-code:%v]\n", rec.Result().StatusCode)

	//Output:
	//test: HttpHandler() -> [status-code:200]

}
