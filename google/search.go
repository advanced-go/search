package google

import (
	"github.com/advanced-go/search/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

func Search[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	if r == nil || r.URL.Query() == nil {
		status := core.NewStatus(http.StatusBadRequest)
		return httpx.NewResponse[E](status.HttpCode(), nil, status.Err)
	}
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, resolver.Url(searchHost, "", searchResource, r.URL.Query(), r.Header), nil)
	req.Header.Set(core.XFrom, module.Authority)
	httpx.Forward(req.Header, r.Header, io.AcceptEncoding)
	resp, status := httpx.DoExchange(req)
	if !status.OK() {
		if !status.Timeout() {
			var e E
			e.Handle(status, core.RequestId(r))
		}
	}
	return resp, status
}
