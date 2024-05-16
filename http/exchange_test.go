package http

import (
	"fmt"
	"github.com/advanced-go/stdlib/io"
	"net/http"
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
	//test: Exchange(nil) -> [status:Invalid Argument [error: request is nil]] [status-code:500]
	//test: Exchange(nil) -> [status:Bad Request [error: invalid URI, authority does not match: "/search" "github/advanced-go/search"]] [status-code:400]
	//test: Exchange(nil) -> [status:Bad Request [error: invalid URI, path only contains an authority: "/github/advanced-go/search"]] [status-code:400]

}

func _ExampleExchange_Google() {
	var buf []byte

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/github/advanced-go/search:google?q=golang", nil)
	resp, _ := Exchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, len(buf))

	//Output:
	//test: Exchange() -> [status-code:200] [content:93444]

}

func _ExampleExchange_Yahoo() {
	var buf []byte

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/github/advanced-go/search:yahoo?q=golang", nil)
	resp, _ := Exchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, len(buf))

	//Output:
	//test: Exchange() -> [status-code:200] [content:93444]

}
