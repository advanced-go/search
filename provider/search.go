package provider

import (
	"context"
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

func search[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) ([]byte, http.Header, *runtime.Status) {
	if values == nil {
		return nil, nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E

	newHeader := make(http.Header)
	if h != nil {
		accept := h.Get(io2.AcceptEncoding)
		if len(accept) > 0 {
			newHeader.Add(io2.AcceptEncoding, accept)
		}
	}
	uri := resolver.Build(searchPath, values.Encode())
	resp, status := exchange.Get(ctx, uri, newHeader)
	if !status.OK() {
		if status.Code != http.StatusGatewayTimeout {
			e.Handle(status, runtime.RequestId(h))
		}
		return nil, nil, status
	}
	buf, status1 := io2.ReadAll(resp.Body, h)
	if !status1.OK() {
		return nil, nil, e.Handle(status1, runtime.RequestId(h))
	}
	resp.ContentLength = int64(len(buf))
	return buf, resp.Header, status1
}
