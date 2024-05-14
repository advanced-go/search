package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/search/google"
	"github.com/advanced-go/search/module"
	"github.com/advanced-go/search/yahoo"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

// DEBUG : https://search.yahoo.com/search?p=golang
// TEST  : https://www.bing.com/search?q=C+Language
// STAGE : https://www.google.com/search?q=C%2B%2B
// PROD  : https://duckduckgo.com/?q=Pascal

const (
	googleProvider = "google"
	yahooProvider  = "yahoo"
)

func Exchange(r *http.Request) (*http.Response, *core.Status) {
	_, path, status0 := httpx.ValidateRequestURL(r, module.Path)
	if !status0.OK() {
		return httpx.NewErrorResponse(status0), status0
	}
	core.AddRequestId(r)
	switch strings.ToLower(path) {
	case googleProvider:
		return google.Search[core.Log](r)
	case yahooProvider:
		return yahoo.Search[core.Log](r)
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", path)))
		return httpx.NewErrorResponse(status), status
	}

	//return google.HttpExchange(r)
}
