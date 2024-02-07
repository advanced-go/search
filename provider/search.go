package provider

import (
	"context"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	searchLocation = PkgPath + ":search"
	httpHandlerLoc = PkgPath + ":HttpHandler"
)

func search[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) ([]byte, http.Header, *runtime.Status) {
	if values == nil {
		return nil, nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E
	var newCtx context.Context
	var status *runtime.Status
	var resp *http.Response

	newHeader := make(http.Header)
	if h != nil {
		accept := h.Get(io2.AcceptEncoding)
		if len(accept) > 0 {
			newHeader.Add(io2.AcceptEncoding, accept)
		}
	}
	uri := resolver.Build(searchPath, values.Encode())
	defer apply(ctx, &newCtx, access.NewRequest(h, http.MethodGet, uri), &resp, googleControllerName, access.StatusCode(&status))()
	resp, status = exchange.Get(newCtx, uri, newHeader)
	if !status.OK() {
		return nil, nil, e.Handle(status, runtime.RequestId(h), searchLocation)
	}
	buf, status1 := io2.ReadAll(resp.Body, h)
	if !status1.OK() {
		return nil, nil, e.Handle(status1, runtime.RequestId(h), searchLocation)
	}
	resp.ContentLength = int64(len(buf))
	return buf, resp.Header, status1
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
