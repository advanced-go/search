package provider

import (
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	searchLocation = PkgPath + ":search"
	httpHandlerLoc = PkgPath + ":HttpHandler"
)

func search[E runtime.ErrorHandler](h http.Header, values url.Values) (*http.Response, runtime.Status) {
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
