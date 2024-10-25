package bing

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"net/http"
)

func Search[E core.ErrorHandler](r *http.Request) (*http.Response, *core.Status) {
	if r == nil || r.URL.Query() == nil {
		status := core.NewStatus(http.StatusBadRequest)
		return httpx.NewResponse(status.HttpCode(), nil, status.Err)
	}
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, resolver.Url(SearchHost, "", SearchPath, r.URL.Query(), r.Header), nil)
	httpx.Forward(req.Header, r.Header) //, io.AcceptEncoding)
	//req.Header.Set(io.ContentEncoding, io.GzipEncoding)
	resp, status := httpx.Exchange(req)
	if !status.OK() {
		if !status.Timeout() {
			var e E
			e.Handle(status.WithRequestId(r))
		}
	}
	return resp, status
}
