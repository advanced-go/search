package yahoo

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

func Search[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	if r == nil || r.URL.Query() == nil {
		status := core.NewStatus(http.StatusBadRequest)
		return httpx.NewResponse[E](status.HttpCode(), nil, status.Err)
	}
	resp, status := httpx.GetExchange(r.Context(), uri.Resolve(searchHost, "", searchResource, r.URL.Query(), r.Header), httpx.Forward(nil, r.Header, io.AcceptEncoding, core.XAuthority))
	if !status.OK() {
		if !status.Timeout() {
			var e E
			e.Handle(status, core.RequestId(r))
		}
	}
	return resp, status
}
