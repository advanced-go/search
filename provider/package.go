package provider

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"strings"
)

type pkg struct{}

const (
	PkgPath = "github/advanced-go/search/provider"

	searchLocation = PkgPath + ":search"
)

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	path, status0 := http2.ValidateRequest(r, PkgPath)
	if !status0.OK() {
		http2.WriteResponse[runtime.Log](w, nil, status0, nil)
		return
	}
	runtime.AddRequestId(r)
	switch strings.ToLower(path) {
	case searchResource:
		buf, status := Search[runtime.Log](r.Header, r.URL.Query())
		http2.WriteResponse[runtime.Log](w, buf, status, status.ContentHeader())
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: [%v]", path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
	}
}

func Search[E runtime.ErrorHandler](h http.Header, values url.Values) ([]byte, runtime.Status) {
	if values == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E

	//return []byte("this is a search result"), runtime.StatusOK()
	newUrl := resolver.Build(searchPath, values.Encode())
	newUrl = "https://search.yahoo.com/search?p=golang"
	resp, status := exchange.Get(newUrl, h)
	if !status.OK() {
		return nil, e.Handle(status, runtime.RequestId(h), searchLocation)
	}
	var buf []byte
	buf, status = runtime.NewBytes(resp)
	if !status.OK() {
		return nil, e.Handle(status, runtime.RequestId(h), searchLocation)
	}
	status = runtime.NewStatusOK()
	for name, _ := range resp.Header {
		status.ContentHeader().Add(name, resp.Header.Get(name))
	}
	//status.ContentHeader().Set(http2.ContentType, resp.Header.Get(http2.ContentType))
	//status.ContentHeader().Set(http2.ContentLength, fmt.Sprintf("%v", len(buf)))
	return buf, status
}
