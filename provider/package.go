package provider

import (
	"errors"
	"fmt"
	"github.com/advanced-go/search/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

const (
	PkgPath = "github/advanced-go/search/provider"
)

// HttpExchange - Process an HTTP exchange
func HttpExchange(r *http.Request) (*http.Response, *core.Status) {
	_, path, status0 := httpx.ValidateRequestURL(r, module.Path)
	if !status0.OK() {
		return httpx.NewErrorResponse(status0), status0
	}
	//path := "search"
	core.AddRequestId(r)
	switch strings.ToLower(path) {
	case searchResource:
		return Search[core.Log](r)
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", path)))
		return httpx.NewErrorResponse(status), status
	}
}
