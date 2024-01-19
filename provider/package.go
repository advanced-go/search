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
	httpHandlerLoc = PkgPath + ":HttpHander"
)

// Accept-Encoding :  gzip, deflate, br

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	path, status0 := http2.ValidateRequest(r, PkgPath)
	if !status0.OK() {
		http2.WriteResponse[runtime.Log](w, status0.Error(), status0, nil)
		return
	}
	runtime.AddRequestId(r)
	switch strings.ToLower(path) {
	case searchResource:
		resp, status := Search[runtime.Log](r.Header, r.URL.Query())
		if !status.OK() {
			http2.WriteResponse[runtime.Log](w, nil, status, nil)
		} else {
			http2.WriteResponse[runtime.Log](w, resp.Body, status, resp.Header)
		}
	default:
		status := runtime.NewStatusError(http.StatusNotFound, httpHandlerLoc, errors.New(fmt.Sprintf("error invalid URI, resource was not found: [%v]", path)))
		http2.WriteResponse[runtime.Log](w, status.Error(), status, nil)
	}
}

func Search[E runtime.ErrorHandler](h http.Header, values url.Values) (*http.Response, runtime.Status) {
	if values == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E

	newHeader := make(http.Header)
	if h != nil {
		accept := h.Get(runtime.AcceptEncoding)
		if len(accept) > 0 {
			newHeader.Add(runtime.AcceptEncoding, accept)
		}
	}
	uri := resolver.Build(searchPath, values.Encode())
	resp, status := exchange.Get(uri, newHeader)
	if !status.OK() {
		return nil, e.Handle(status, runtime.RequestId(h), searchLocation)
	}
	return resp, status
}

/*
	var buf []byte
	buf, status = runtime.ReadAll(resp.Body, nil)
	if !status.OK() {
		return nil, e.Handle(status, runtime.RequestId(h), searchLocation)
	}
	status = runtime.NewStatusOK()
	status.ContentHeader().Add(http2.ContentType, resp.Header.Get(http2.ContentType))
	if len(accept) > 0 {
		status.ContentHeader().Add(runtime.ContentEncoding, resp.Header.Get(runtime.ContentEncoding))
	}

*/
