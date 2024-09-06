package bing

import (
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
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, Url(SearchHost, SearchPath, r.URL.Query(), r.Header), nil)
	httpx.Forward(req.Header, r.Header, io.AcceptEncoding)
	req.Header.Set(io.ContentEncoding, io.GzipEncoding)
	resp, status := httpx.Exchange(req)
	if !status.OK() {
		if !status.Timeout() {
			var e E
			e.Handle(status.WithRequestId(r))
		}
	}
	return resp, status
}
