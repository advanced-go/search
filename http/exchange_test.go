package http

import (
	"fmt"
	"github.com/advanced-go/search/google"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

const (
	getResp = "file://[cwd]/httptest/get-resp.txt"
)

func ExampleExchange_Invalid() {
	resp, status := Exchange(nil)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	req, _ := http.NewRequest("", "http://www.google.com/search?q=golang", nil)
	resp, status = Exchange(req)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	req, _ = http.NewRequest("", "http://www.google.com/github/advanced-go/search", nil)
	resp, status = Exchange(req)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	//Output:
	//test: Exchange(nil) -> [status:Bad Request] [status-code:400]
	//test: Exchange(nil) -> [status:Bad Request] [status-code:400]
	//test: Exchange(nil) -> [status:Bad Request] [status-code:400]

}

func ExampleExchange_Authority() {
	req, _ := http.NewRequest(http.MethodGet, core.AuthorityRootPath, nil)
	resp, status := Exchange(req)
	fmt.Printf("test: Exchange(\"/authority\") -> [status:%v] [status-code:%v] [auth:%v]\n", status, resp.StatusCode, resp.Header.Get(core.XAuthority))

	//Output:
	//test: Exchange("/authority") -> [status:OK] [status-code:200] [auth:github/advanced-go/search]

}

func _ExampleExchange_Version() {
	req, _ := http.NewRequest("", "http://locahhost:8081/github/advanced-go/search:version", nil)
	resp, status := Exchange(req)
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange(\"/version\") -> [status:%v] [status-code:%v] [content:%v]\n", status, resp.StatusCode, string(buf))

	//Output:
	//test: Exchange("/version") -> [status:OK] [status-code:200] [content:{
	// "version": "1.1.1"
	//  }]

}

func _ExampleExchange_Health() {
	req, _ := http.NewRequest("", "http://locahhost:8081/github/advanced-go/search:health/readiness", nil)
	resp, status := Exchange(req)
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange(\"/health/readiness\") -> [status:%v] [status-code:%v] [content:%v]\n", status, resp.StatusCode, string(buf))

	req, _ = http.NewRequest("", "http://locahhost:8081/github/advanced-go/search:health/liveness", nil)
	resp, status = Exchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange(\"/health/liveness\") -> [status:%v] [status-code:%v] [content:%v]\n", status, resp.StatusCode, string(buf))

	//Output:
	//test: Exchange("/health/readiness") -> [status:OK] [status-code:200] [content:{
	// "status": "up"
	//}]
	//test: Exchange("/health/liveness") -> [status:OK] [status-code:200] [content:{
	// "status": "up"
	//}]

}

func ExampleExchange_Google() {
	var buf []byte

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/github/advanced-go/search:google?q=golang", nil)
	resp, _ := Exchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, len(buf) > 0)

	//Output:
	//test: Exchange() -> [status-code:200] [content:true]

}

func ExampleExchange_Google_Override() {
	var buf []byte
	query := "q=golang"

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/github/advanced-go/search:google?q=golang", nil)
	path := uri.BuildPath(google.SearchPath, query)
	req.Header.Add(path, getResp)
	resp, _ := Exchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange() -> [status-code:%v] [content:%v] [buf:%v]\n", resp.StatusCode, len(buf) > 0, len(buf)) //string(buf))

	//Output:
	//test: Exchange() -> [status-code:200] [content:true] [buf:279]

}

func ExampleExchange_Yahoo() {
	var buf []byte

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/github/advanced-go/search:yahoo?q=golang", nil)
	resp, _ := Exchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, len(buf) > 0)

	//Output:
	//test: Exchange() -> [status-code:200] [content:true]

}

func ExampleExchange_Bing() {
	var buf []byte

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/github/advanced-go/search:bing?q=golang", nil)
	resp, _ := Exchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, len(buf) > 0)

	//Output:
	//test: Exchange() -> [status-code:200] [content:true]

}

func ExampleExchange_Duck() {
	var buf []byte

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/github/advanced-go/search:duck?q=golang", nil)
	resp, _ := Exchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, len(buf) > 0)

	//Output:
	//test: Exchange() -> [status-code:200] [content:true]

}
