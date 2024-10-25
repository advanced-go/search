package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"github.com/advanced-go/common/uri"
	"github.com/advanced-go/search/bing"
	"github.com/advanced-go/search/duck"
	"github.com/advanced-go/search/google"
	"github.com/advanced-go/search/module"
	"github.com/advanced-go/search/yahoo"
	"net/http"
)

const (
	healthLivenessPath  = "health/liveness"
	healthReadinessPath = "health/readiness"
	versionPath         = "version"
	authorityPath       = "authority"
)

// DEBUG : https://search.yahoo.com/search?p=golang
// TEST  : https://www.bing.com/search?q=C+Language
// STAGE : https://www.google.com/search?q=C%2B%2B
// PROD  : https://duckduckgo.com/?q=Pascal

const (
	googleProvider = "google"
	yahooProvider  = "yahoo"
	bingProvider   = "bing"
	duckProvider   = "duck"
)

var authorityResponse = NewAuthorityResponse(module.Authority)

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (resp *http.Response, status *core.Status) {
	if r == nil || r.URL == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}, core.StatusBadRequest()
	}
	p, err := uri.ValidateURL(r.URL, module.Authority)
	if err != nil {
		status1 := core.NewStatusError(http.StatusBadRequest, err)
		resp, _ = httpx.NewResponse(status1.HttpCode(), nil, status1.Err)
		return resp, status1
	}
	core.AddRequestId(r)
	switch p.Path {
	case googleProvider:
		resp, status = google.Search[core.Log](r)
		resp.Header.Add(core.XRoute, google.SearchRoute)
		return
	case yahooProvider:
		resp, status = yahoo.Search[core.Log](r)
		resp.Header.Add(core.XRoute, yahoo.SearchRoute)
		return
	case bingProvider:
		resp, status = bing.Search[core.Log](r)
		resp.Header.Add(core.XRoute, bing.SearchRoute)
		return
	case duckProvider:
		resp, status = duck.Search[core.Log](r)
		resp.Header.Add(core.XRoute, duck.SearchRoute)
		return
	case versionPath:
		resp, status = NewVersionResponse(module.Version), core.StatusOK()
		resp.Header.Add(core.XRoute, module.VersionRoute)
		return
	case authorityPath:
		return authorityResponse, core.StatusOK()
	case healthReadinessPath, healthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", p.Resource)))
		return httpx.NewResponse(status.HttpCode(), nil, status.Err)
	}
}
