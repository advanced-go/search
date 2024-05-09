package provider

import (
	"fmt"
	"github.com/advanced-go/stdlib/io"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

func ExampleHttpExchanger_Error() {
	// Need to set overrides, the returned func will reset original values
	defer resolver.SetTemplates([]uri.Attr{{searchPath, resultUri}})()

	var buf []byte

	resp, _ := HttpExchange(nil)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+"invalid-path"+":search?q=golang", nil)
	resp, _ = HttpExchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":searchBad?q=golang", nil)
	resp, _ = HttpExchange(req)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: HttpExchange() -> [status-code:500] [content:error: Request is nil]
	//test: HttpExchange() -> [status-code:400] [content:error: invalid URI, NID does not match: "/invalid-path:search" "github/advanced-go/search/provider"]
	//test: HttpExchange() -> [status-code:404] [content:error invalid URI, resource not found: [searchBad]]

}
