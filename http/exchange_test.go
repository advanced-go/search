package http

import (
	"fmt"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

func ExampleExchang() {
	var buf []byte

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/github/advanced-go/search:google?q=golang", nil)
	resp, _ := Exchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: Exchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, len(buf))

	//Output:
	//test: Exchange() -> [status-code:200] [content:93444]

}
