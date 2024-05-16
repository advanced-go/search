package google

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

func Search[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	if r == nil || r.URL.Query() == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}, core.NewStatus(http.StatusBadRequest)
	}
	resp, status := httpx.GetExchange(r.Context(), resolver.Build(searchPath, r.URL.Query().Encode()), httpx.Forward(nil, r.Header, io.AcceptEncoding))
	if !status.OK() {
		if !status.Timeout() {
			var e E
			e.Handle(status, core.RequestId(r))
		}
	}
	return resp, status
}
