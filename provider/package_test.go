package provider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/io"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"time"
)

const (
	resultUri = "file://[cwd]/providertest/resource/query-result.txt"
)

type pkg struct{}

func Example_PkgPath() {
	dotCom := ".com"
	pkgPath := reflect.TypeOf(any(pkg{})).PkgPath()
	fmt.Printf("test: PkgPath = \"%v\"\n", strings.Replace(pkgPath, dotCom, "", len(dotCom)))

	//Output:
	//test: PkgPath = "github/advanced-go/search/provider"

}

func ExampleHttpExchange_Error() {
	// Need to set overrides, the returned func will reset original values
	defer resolver.SetTemplates([]uri.Attr{{searchPath, resultUri}})()

	rec := httptest.NewRecorder()
	HttpExchange(rec, nil)
	buf, _ := io.ReadAll(rec.Result().Body, nil)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+"invalid-path"+":search?q=golang", nil)
	rec = httptest.NewRecorder()
	HttpExchange(rec, req)
	buf, _ = io.ReadAll(rec.Result().Body, nil)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":searchBad?q=golang", nil)
	rec = httptest.NewRecorder()
	HttpExchange(rec, req)
	buf, _ = io.ReadAll(rec.Result().Body, nil)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpExchange() -> [status-code:500] [content:error: Request is nil]
	//test: HttpExchange() -> [status-code:400] [content:error: invalid URI, NID does not match: "/invalid-path:search" "github/advanced-go/search/provider"]
	//test: HttpExchange() -> [status-code:404] [content:error invalid URI, resource not found: [searchBad]]

}

func ExampleHttpExchange_Text() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	rec := httptest.NewRecorder()
	HttpExchange(rec, req)
	buf, status := io.ReadAll(rec.Result().Body, nil)
	ct := http.DetectContentType(buf)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", rec.Result().StatusCode, status, ct)

	//Output:
	//test: HttpExchange() -> [status-code:200] [read-all:OK] [content-type:text/html; charset=utf-8]

}

func ExampleHttpExchange_Gzip() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	req.Header.Add(io.AcceptEncoding, io.GzipEncoding)
	rec := httptest.NewRecorder()
	HttpExchange(rec, req)
	buf, status := io.ReadAll(rec.Result().Body, nil)
	ct := http.DetectContentType(buf)
	fmt.Printf("test: HttpExchange-Gzip() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", rec.Result().StatusCode, status, ct)

	buf, status = io.ReadAll(bytes.NewReader(buf), rec.Result().Header)
	ct = http.DetectContentType(buf)
	fmt.Printf("test: HttpExchange-Gzip-Decoded() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", rec.Result().StatusCode, status, ct)

	//Output:
	//test: HttpExchange-Gzip() -> [status-code:200] [read-all:OK] [content-type:application/x-gzip]
	//test: HttpExchange-Gzip-Decoded() -> [status-code:200] [read-all:OK] [content-type:text/html; charset=utf-8]

}

func ExampleHttpExchange_Timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	rec := httptest.NewRecorder()
	HttpExchange(rec, req)
	buf, status := io.ReadAll(rec.Result().Body, nil)
	ct := http.DetectContentType(buf)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [read-all:%v] [content-type:%v] [buf:%v]\n", rec.Result().StatusCode, status, ct, buf)

	//Output:
	//test: HttpExchange() -> [status-code:504] [read-all:OK] [content-type:text/plain; charset=utf-8] [buf:[]]

}
