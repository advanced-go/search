package service

import (
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/search/google"
	"net/http"
)

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	buf, status := google.Search(r)
	http2.WriteResponse[runtime.Log](w, buf, status, status.ContentHeader())
}
