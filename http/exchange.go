package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/search/google"
	"github.com/advanced-go/search/module"
	"github.com/advanced-go/search/yahoo"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
	"time"
)

// DEBUG : https://search.yahoo.com/search?p=golang
// TEST  : https://www.bing.com/search?q=C+Language
// STAGE : https://www.google.com/search?q=C%2B%2B
// PROD  : https://duckduckgo.com/?q=Pascal

const (
	googleProvider = "google"
	yahooProvider  = "yahoo"
)

var (
	versionResponse = httpx.NewResponse(core.StatusOK(), core.VersionContent(module.Version))
)

// Controllers - authority controllers
func Controllers() []*controller.Controller {
	return []*controller.Controller{
		controller.NewController("google-search", controller.NewPrimaryResource("www.google.com", time.Second*2, "", nil), nil),
		controller.NewController("yahoo-search", controller.NewPrimaryResource("search.yahoo.com", time.Second*2, "", nil), nil),
	}
}

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	_, path, status0 := httpx.ValidateRequestURL(r, module.Authority)
	if !status0.OK() {
		return httpx.NewErrorResponse(status0), status0
	}
	core.AddRequestId(r)
	switch strings.ToLower(path) {
	case googleProvider:
		return google.Search[core.Log](r)
	case yahooProvider:
		return yahoo.Search[core.Log](r)
	case core.VersionPath:
		return versionResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.HealthResponseOK, core.StatusOK()
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", path)))
		return httpx.NewErrorResponse(status), status
	}
}
