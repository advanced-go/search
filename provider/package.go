package provider

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

const (
	PkgPath = "github/advanced-go/search/provider"
)

// HttpExchange - Process an HTTP exchange
func HttpExchange(w http.ResponseWriter, r *http.Request) {
	path, status0 := httpx.ValidateRequest(r, PkgPath)
	if !status0.OK() {
		httpx.WriteResponse[core.Log](w, nil, status0.HttpCode(), status0.Err)
		return
	}
	httpx.AddRequestId(r)
	switch strings.ToLower(path) {
	case searchResource:
		buf, h, status := search[core.Log](r.Context(), r.Header, r.URL.Query())
		httpx.WriteResponse[core.Log](w, h, status.HttpCode(), buf)
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", path)))
		httpx.WriteResponse[core.Log](w, nil, status.HttpCode(), status.Err)
	}
}
