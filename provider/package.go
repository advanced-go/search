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
		httpx.WriteResponse[core.Log](w, status0.Err, status0, nil)
		return
	}
	httpx.AddRequestId(r)
	switch strings.ToLower(path) {
	case searchResource:
		buf, h, status := search[core.Log](r.Context(), r.Header, r.URL.Query())
		if !status.OK() {
			httpx.WriteResponse[core.Log](w, nil, status, nil)
		} else {
			httpx.WriteResponse[core.Log](w, buf, status, h)
		}
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: [%v]", path)))
		httpx.WriteResponse[core.Log](w, status.Err, status, nil)
	}
}
