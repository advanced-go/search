package google

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"reflect"
)

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()
	fmt.Printf("test: PkgPath = \"%v\"\n", pkgUri2)

	//Output:
	//test: PkgPath = "github.com/advanced-go/search/google"

}

func Example_Search() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+"/search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	_, status := search[runtime.Output](req)
	fmt.Printf("test: search(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: search(http://localhost:8080/github.com/advanced-go/search/google/search?q=golang) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:115289]

}

func searchOverrideFail(id string) (string, string) {
	switch id {
	case searchTag:
		return "file://[cwd]/resource/query-result.txt", ""
	}
	return "", ""
}

func Example_Search_OverrideFail() {
	resolver.SetOverride(searchOverrideFail, "")
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+"/search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	_, status := search[runtime.Output](req)
	fmt.Printf("test: search(%v) -> [status:%v]\n", req.URL.String(), status)

	//Output:
	//{ "code":91, "status":"I/O Failure", "request-id":"invalid-change", "trace" : [ "github.com/advanced-go/search/tree/main/google:Search","github.com/advanced-go/core/tree/main/exchange:Do","github.com/advanced-go/core/tree/main/exchange:readResponse" ], "errors" : [ "open C:\Users\markb\GitHub\search\google\resource\query-result.txt: The system cannot find the path specified." ] }
	//test: search(http://localhost:8080/github.com/advanced-go/search/google/search?q=golang) -> [status:I/O Failure [open C:\Users\markb\GitHub\search\google\resource\query-result.txt: The system cannot find the path specified.]]

}

func searchOverrideSuccess(id string) (string, string) {
	switch id {
	case searchTag:
		return "file://[cwd]/googletest/resource/query-result.txt", ""
	}
	return "", ""
}

func Example_Search_OverrideSuccess() {
	resolver.SetOverride(searchOverrideSuccess, "")
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+"/search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	_, status := search[runtime.Output](req)
	fmt.Printf("test: search(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: search(http://localhost:8080/github.com/advanced-go/search/google/search?q=golang) -> [status:OK] [content-type:text/plain] [content-length:49]

}
