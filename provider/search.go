package provider

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"net/http"
	"net/url"
)

func search[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) ([]byte, http.Header, *core.Status) {
	if values == nil {
		return nil, nil, core.NewStatus(http.StatusBadRequest)
	}
	var e E

	newHeader := make(http.Header)
	if h != nil {
		accept := h.Get(io.AcceptEncoding)
		if len(accept) > 0 {
			newHeader.Add(io.AcceptEncoding, accept)
		}
	}
	uri := resolver.Build(searchPath, values.Encode())
	resp, status := httpx.Get(ctx, uri, newHeader)
	if !status.OK() {
		if status.Code != http.StatusGatewayTimeout {
			e.Handle(status, core.RequestId(h))
		}
		return nil, nil, status
	}
	buf, status1 := io.ReadAll(resp.Body, h)
	if !status1.OK() {
		return nil, nil, e.Handle(status1, core.RequestId(h))
	}
	resp.ContentLength = int64(len(buf))
	return buf, resp.Header, status1
}
