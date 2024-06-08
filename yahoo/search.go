package yahoo

import (
	"github.com/advanced-go/search/module"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"time"
)

func Search[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	if r == nil || r.URL.Query() == nil {
		status := core.NewStatus(http.StatusBadRequest)
		return httpx.NewResponse[E](status.HttpCode(), nil, status.Err)
	}
	start := time.Now().UTC()
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, uri.Resolve(searchHost, "", searchResource, r.URL.Query(), r.Header), nil)
	req.Header.Set(core.XFrom, module.Authority)
	httpx.Forward(nil, r.Header, io.AcceptEncoding)
	resp, status := httpx.DoExchange(req)
	if !status.OK() {
		if !status.Timeout() {
			var e E
			e.Handle(status, core.RequestId(r))
		}
	}
	access.LogEgress(start, time.Since(start), req, resp, module.Authority, module.YahooRouteName, "", 0, 0, 0, "")
	return resp, status
}
