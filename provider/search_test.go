package provider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	"github.com/advanced-go/stdlib/uri"
	"io"
	"net/http"
	"time"
)

func ExampleSearch_Success() {
	resolver.SetTemplates([]uri.Attr{{searchPath, "https://www.google.com/search?q=golang"}})
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5000)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := Search[core.Output](req)
	buf, _ := io2.ReadAll(resp.Body, resp.Header)

	fmt.Printf("test: Search(%v) -> [status:%v] [status-code:%v] [content:%v]\n", req.URL.String(), status, status.Code, len(buf) > 0)

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:OK] [status-code:200] [content:true]

}

func ExampleSearch_Deadline_Exceeded() {
	resolver.SetTemplates([]uri.Attr{{searchPath, "https://www.google.com/search?q=golang"}})
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := Search[core.Output](req)
	buf, _ := io2.ReadAll(resp.Body, resp.Header)

	fmt.Printf("test: Search(%v) -> [status:%v] [status-code:%v] [content:%v]\n", req.URL.String(), status, status.Code, len(buf))

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search/provider:search?q=golang) -> [status:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]] [status-code:504] [content:0]

}

func ExampleSearch_Text() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	resp, status := Search[core.Output](req)
	buf, _ := io2.ReadAll(resp.Body, resp.Header)
	ct := http.DetectContentType(buf)
	fmt.Printf("test: Search-Text() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", resp.StatusCode, status, ct)

	//Output:
	//test: Search-Text() -> [status-code:200] [read-all:OK] [content-type:text/html; charset=utf-8]

}

func ExampleSearch_Gzip() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	req.Header.Add(io2.AcceptEncoding, io2.GzipEncoding)

	resp, status := Search[core.Output](req)
	buf, _ := io.ReadAll(resp.Body)
	ct := http.DetectContentType(buf)
	fmt.Printf("test: Search-Gzip() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", resp.StatusCode, status, ct)

	buf, status = io2.ReadAll(bytes.NewReader(buf), resp.Header)
	ct = http.DetectContentType(buf)
	fmt.Printf("test: Search-Gzip-Decoded() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", resp.StatusCode, status, ct)

	//Output:
	//test: Search-Gzip() -> [status-code:200] [read-all:OK] [content-type:application/x-gzip]
	//test: Search-Gzip-Decoded() -> [status-code:200] [read-all:OK] [content-type:text/html; charset=utf-8]

}
