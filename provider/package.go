package provider

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"strings"
)

const (
	PkgPath = "github/advanced-go/search/provider"
)

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	path, status0 := http2.ValidateRequest(r, PkgPath)
	if !status0.OK() {
		http2.WriteResponse[runtime.Log](w, status0.Error(), status0, nil)
		return
	}
	runtime.AddRequestId(r)
	switch strings.ToLower(path) {
	case searchResource:
		buf, h, status := search[runtime.Log](nil, r.Header, r.URL.Query())
		if !status.OK() {
			http2.WriteResponse[runtime.Log](w, nil, status, nil)
		} else {
			http2.WriteResponse[runtime.Log](w, buf, status, h)
		}
	default:
		status := runtime.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: [%v]", path)), nil)
		http2.WriteResponse[runtime.Log](w, status.Error(), status, nil)
	}
}
