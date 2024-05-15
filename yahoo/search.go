package yahoo

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

	uri := resolver.Build(searchPath, r.URL.Query().Encode())
	resp, status := httpx.GetExchange(r.Context(), uri, httpx.Forward(nil, r.Header, io.AcceptEncoding))
	if !status.OK() {
		if !status.Timeout() {
			e.Handle(status, core.RequestId(r))
		}
	}
	return resp, status
}
