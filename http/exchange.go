package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/search/bing"
	"github.com/advanced-go/search/duck"
	"github.com/advanced-go/search/google"
	"github.com/advanced-go/search/module"
	"github.com/advanced-go/search/yahoo"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
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

var authorityResponse = httpx.NewAuthorityResponse(module.Authority)

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	if r == nil || r.URL == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}, core.StatusBadRequest()
	}
	p, status := httpx.ValidateURL(r.URL, module.Authority)
	if !status.OK() {
		return httpx.NewResponse[core.Log](status.HttpCode(), nil, status.Err)
	}
	core.AddRequestId(r)
	switch p.Path {
	case googleProvider:
		return google.Search[core.Log](r)
	case yahooProvider:
		return yahoo.Search[core.Log](r)
	case bingProvider:
		return bing.Search[core.Log](r)
	case duckProvider:
		return duck.Search[core.Log](r)
	case core.VersionPath:
		return httpx.NewVersionResponse(module.Version), core.StatusOK()
	case core.AuthorityPath:
		return authorityResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", p.Resource)))
		return httpx.NewResponse[core.Log](status.HttpCode(), nil, status.Err)
	}
}
