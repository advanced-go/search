package provider

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/uri"
	"net/http"
	"net/url"
	"strings"
)

type pkg struct{}

// http://localhost:8080/search?q=golang

const (
	PkgPath        = "github.com/advanced-go/search/provider"
	searchLocation = PkgPath + ":search"
	searchResource = "search"
)

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		http2.WriteResponse[runtime.Log](w, nil, runtime.NewStatus(http.StatusBadRequest), nil)
		return
	}
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	if !ok || nid != PkgPath {
		status := runtime.NewStatusWithContent(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid URI, path is not valid: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	runtime.AddRequestId(r)
	switch strings.ToLower(rsc) {
	case searchResource:
		buf, status := search[runtime.Log](r.Header, r.URL.Query())
		http2.WriteResponse[runtime.Log](w, buf, status, status.ContentHeader())
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: [%v]", rsc)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
	}

}

func search[E runtime.ErrorHandler](h http.Header, values url.Values) ([]byte, runtime.Status) {
	if values == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E

	newUrl := resolver.Build(searchTag, searchPath, newValues(values).Encode())
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
