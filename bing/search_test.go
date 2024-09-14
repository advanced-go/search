package bing

import (
	"bytes"
	"context"
	"fmt"
	"github.com/advanced-go/search/module"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
	"net/http"
	"time"
)

func ExampleSearch_Error() {
	resp, status := Search[core.Output](nil)
	fmt.Printf("test: Search() -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	//Output:
	//test: Search() -> [status:Bad Request] [status-code:400]

}

func ExampleSearch_Success() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5000)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080"+"/"+module.Authority+":bing?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := Search[core.Output](req)
	buf, _ := io2.ReadAll(resp.Body, resp.Header)

	fmt.Printf("test: Search(%v) -> [status:%v] [status-code:%v] [content:%v]\n", req.URL.String(), status, status.Code, len(buf) > 0)

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search:bing?q=golang) -> [status:OK] [status-code:200] [content:true]

}

func ExampleSearch_Deadline_Exceeded() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080"+"/"+module.Authority+":bing?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := Search[core.Output](req)
	buf, _ := io2.ReadAll(resp.Body, resp.Header)

	fmt.Printf("test: Search(%v) -> [status:%v] [status-code:%v] [content:%v]\n", req.URL.String(), status, status.Code, len(buf))

	//Output:
	//test: Search(http://localhost:8080/github/advanced-go/search:bing?q=golang) -> [status:Timeout [Get "https://www.bing.com/search?q=golang": context deadline exceeded]] [status-code:504] [content:0]

}

func ExampleSearch_Text() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+module.Authority+":bing?q=golang", nil)
	resp, status := Search[core.Output](req)
	buf, _ := io2.ReadAll(resp.Body, resp.Header)
	ct := http.DetectContentType(buf)
	fmt.Printf("test: Search-Text() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", resp.StatusCode, status, ct)

	//Output:
	//test: Search-Text() -> [status-code:200] [read-all:OK] [content-type:text/html; charset=utf-8]

}

func _ExampleSearch_Gzip() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+module.Authority+":bing?q=golang", nil)
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
