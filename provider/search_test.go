package provider

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	uri2 "github.com/advanced-go/core/uri"
	"net/http"
)

func ExampleSearch() {
	resolver.SetOverrides([]uri2.Pair{{searchPath, "https://www.google.com/search?q=golang"}})
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := search[runtime.Output](nil, req.Header, req.URL.Query())
	buf, _ := io2.ReadAll(resp.Body, resp.Header)

	fmt.Printf("test: Search(%v) -> [status:%v] [content-type:%v] [content-encoding:%v] [content-length:%v]\n", req.URL.String(), status, resp.Header.Get(http2.ContentType), resp.Header.Get(io2.ContentEncoding), len(buf))

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-encoding:] [content-length:116450]

}

func ExampleSearch_Timeout() {
	resolver.SetOverrides([]uri2.Pair{{searchPath, "https://www.google.com/search?q=golang"}})
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := search[runtime.Output](nil, req.Header, req.URL.Query())
	buf, _ := io2.ReadAll(resp.Body, resp.Header)

	fmt.Printf("test: Search(%v) -> [status:%v] [content-type:%v] [content-encoding:%v] [content-length:%v]\n", req.URL.String(), status, resp.Header.Get(http2.ContentType), resp.Header.Get(io2.ContentEncoding), len(buf))

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-encoding:] [content-length:116450]

}
