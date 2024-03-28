package provider

import (
	"context"
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	uri2 "github.com/advanced-go/core/uri"
	"net/http"
	"time"
)

func ExampleSearch() {
	resolver.SetTemplates([]uri2.Pair{{searchPath, "https://www.google.com/search?q=golang"}})
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	buf, h, status := search[runtime.Output](nil, req.Header, req.URL.Query())
	//buf, _ := io2.ReadAll(resp.Body, resp.Header)

	fmt.Printf("test: Search(%v) -> [status:%v] [content-type:%v] [content-encoding:%v] [content-length:%v]\n", req.URL.String(), status, h.Get(http2.ContentType), h.Get(io2.ContentEncoding), len(buf))

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-encoding:] [content-length:116450]

}

func ExampleSearch_Timeout() {
	resolver.SetTemplates([]uri2.Pair{{searchPath, "https://www.google.com/search?q=golang"}})
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancel()
	_, _, status := search[runtime.Output](ctx, req.Header, req.URL.Query())

	fmt.Printf("test: Search(%v) -> [status:%v] [status-code:%v]\n", req.URL.String(), status, status.Code)

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]] [status-code:504]

}
