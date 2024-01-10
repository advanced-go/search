package service

import (
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/search/provider"
	"net/http"
)

const (
	PkgPath = "github.com/advanced-go/search/service"
)

type pkg struct{}

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	buf, status := provider.Search(r.Header, r.URL.Query())
	http2.WriteResponse[runtime.Log](w, buf, status, status.ContentHeader())
}
