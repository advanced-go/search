package google

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

func Search[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	if r == nil || r.URL.Query() == nil {
		status := core.NewStatus(http.StatusBadRequest)
		return httpx.NewResponse(status, status.Err), status
	}
	resp, status := httpx.GetExchange(r.Context(), buildURL(r.URL), httpx.Forward(nil, r.Header, io.AcceptEncoding))
	if !status.OK() {
		if !status.Timeout() {
			var e E
			e.Handle(status, core.RequestId(r))
		}
	}
	return resp, status
}
