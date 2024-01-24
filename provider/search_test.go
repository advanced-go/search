package provider

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func ExampleSearch() {
	resolver.SetOverrides([]runtime.Pair{{searchPath, "https://www.google.com/search?q=golang"}})
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := search[runtime.Output](nil, req.URL.Query())
	//s := string(buf)
	//if len(s) > 0 {
	//}

	fmt.Printf("test: Search(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, resp.Header.Get(http2.ContentType), 0)

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:115289]

}

func ExampleSearch_Override() {
	resolver.SetOverrides([]runtime.Pair{{searchPath, resultUri}})
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := search[runtime.Output](nil, req.URL.Query())
	buf, _ := runtime.ReadAll(resp.Body, nil)
	s := string(buf)
	s = s[:len(s)-2]
	fmt.Printf("test: Search(%v) -> [status:%v] [content:%v] [content-type:%v]\n", req.URL.String(), status, s, resp.Header.Get(http2.ContentType))

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:OK] [content:This is an alternate result for a Google query.] [content-type:text/plain]

}
