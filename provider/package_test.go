package provider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/advanced-go/core/io2"
	uri2 "github.com/advanced-go/core/uri"
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

func ExampleHttpHandler_Error() {
	// Need to set overrides, the returned func will reset original values
	defer resolver.SetTemplates([]uri2.Pair{{searchPath, resultUri}})()

	rec := httptest.NewRecorder()
	HttpHandler(rec, nil)
	fmt.Printf("test: HttpHandler() -> [status-code:%v]\n", rec.Result().StatusCode)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+"invalid-path"+":search?q=golang", nil)
	rec = httptest.NewRecorder()
	HttpHandler(rec, req)
	buf, _ := io2.ReadAll(rec.Result().Body, nil)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":searchBad?q=golang", nil)
	rec = httptest.NewRecorder()
	HttpHandler(rec, req)
	buf, _ = io2.ReadAll(rec.Result().Body, nil)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:500]
	//test: HttpHandler() -> [status-code:400] [content:error: invalid URI, NID does not match: "/invalid-path:search" "github/advanced-go/search/provider"]
	//test: HttpHandler() -> [status-code:404] [content:error invalid URI, resource was not found: [searchBad]]

}

func ExampleHttpHandler_Text() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	rec := httptest.NewRecorder()
	HttpHandler(rec, req)
	buf, status := io2.ReadAll(rec.Result().Body, nil)
	ct := http.DetectContentType(buf)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", rec.Result().StatusCode, status, ct)

	//Output:
	//test: HttpHandler() -> [status-code:200] [read-all:OK] [content-type:text/html; charset=utf-8]

}

func ExampleHttpHandler_Gzip() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	req.Header.Add(io2.AcceptEncoding, io2.GzipEncoding)
	rec := httptest.NewRecorder()
	HttpHandler(rec, req)
	buf, status := io2.ReadAll(rec.Result().Body, nil)
	ct := http.DetectContentType(buf)
	fmt.Printf("test: HttpHandler-Gzip() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", rec.Result().StatusCode, status, ct)

	buf, status = io2.ReadAll(bytes.NewReader(buf), rec.Result().Header)
	ct = http.DetectContentType(buf)
	fmt.Printf("test: HttpHandler-Gzip-Decoded() -> [status-code:%v] [read-all:%v] [content-type:%v]\n", rec.Result().StatusCode, status, ct)

	//Output:
	//test: HttpHandler-Gzip() -> [status-code:200] [read-all:OK] [content-type:application/x-gzip]
	//test: HttpHandler-Gzip-Decoded() -> [status-code:200] [read-all:OK] [content-type:text/html; charset=utf-8]

}

func ExampleHttpHandler_Timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080"+"/"+PkgPath+":search?q=golang", nil)
	rec := httptest.NewRecorder()
	HttpHandler(rec, req)
	buf, status := io2.ReadAll(rec.Result().Body, nil)
	ct := http.DetectContentType(buf)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [read-all:%v] [content-type:%v] [buf:%v]\n", rec.Result().StatusCode, status, ct, buf)

	//Output:
	//test: HttpHandler() -> [status-code:504] [read-all:OK] [content-type:text/plain; charset=utf-8] [buf:[]]

}
