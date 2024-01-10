package google

import (
	"fmt"
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

type pkg struct{}

// http://localhost:8080/search?q=golang

const (
	PkgPath        = "github.com/advanced-go/search/google"
	searchLocation = PkgPath + ":Search"

	repZero = "search"
	repOne  = "http://localhost:8080/search"
	real    = "http://www.google.com/search"
)

// Search - search handler
// Uses : https://www.google.com/search
func Search(h http.Header, values url.Values) ([]byte, runtime.Status) {
	return search[runtime.Log](h, values)
}

func search[E runtime.ErrorHandler](h http.Header, values url.Values) ([]byte, runtime.Status) {
	if values == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E
	//requestId := "invalid-change"
	h = runtime.AddRequestId(h)
	newUrl := resolver.Build(searchTag, values)
	req, err := http.NewRequest(http.MethodGet, newUrl, nil)
	if err != nil {
		return nil, e.Handle(runtime.NewStatusError(http.StatusInternalServerError, searchLocation, err), runtime.RequestId(h), "")
	}
	// exchange.Do() will always return a non nil *http.Response
	resp, status := exchange.Do(req)
	if !status.OK() {
		return nil, e.Handle(status, runtime.RequestId(h), searchLocation)
	}
	var buf []byte
	buf, status = runtime.NewBytes(resp)
	if !status.OK() {
		return nil, e.Handle(status, runtime.RequestId(h), searchLocation)
	}
	status = runtime.NewStatusOK()
	status.ContentHeader().Set(http2.ContentType, resp.Header.Get(http2.ContentType))
	status.ContentHeader().Set(http2.ContentLength, fmt.Sprintf("%v", len(buf)))
	return buf, status
}
