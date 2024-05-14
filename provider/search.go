package provider

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

func Search[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	if r == nil || r.URL.Query() == nil {
		return nil, core.NewStatus(http.StatusBadRequest)
	}
	var e E

	newHeader := make(http.Header)
	if r.Header != nil {
		accept := r.Header.Get(io.AcceptEncoding)
		if len(accept) > 0 {
			newHeader.Add(io.AcceptEncoding, accept)
		}
	}
	uri := resolver.Build(searchPath, r.URL.Query().Encode())
	resp, status := httpx.Get(r.Context(), uri, newHeader)
	if !status.OK() {
		if status.Code != http.StatusGatewayTimeout {
			e.Handle(status, core.RequestId(r))
		}
	}
	return resp, status
}
