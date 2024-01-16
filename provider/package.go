package provider

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"strings"
)

type pkg struct{}

const (
	PkgPath = "github/advanced-go/search/provider"

	searchLocation = PkgPath + ":search"
)

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	path, status0 := http2.ValidateRequest(r, PkgPath)
	if !status0.OK() {
		http2.WriteResponse[runtime.Log](w, nil, status0, nil)
		return
	}
	//fmt.Printf("request header accept-encoding : %v\n",r.)
	runtime.AddRequestId(r)
	switch strings.ToLower(path) {
	case searchResource:
		buf, status := Search[runtime.Log](r.Header, r.URL.Query(), false)
		if !status.OK() {
			http2.WriteResponse[runtime.Log](w, nil, status, nil)
		} else {
			if r.Header.Get("Accept-Encoding") == "gzip" {
				status.ContentHeader().Add("Content-Encoding", "gzip")
			}
			http2.WriteResponse[runtime.Log](w, buf, status, status.ContentHeader())
		}
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: [%v]", path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
	}
}

func Search[E runtime.ErrorHandler](h http.Header, values url.Values, override bool) ([]byte, runtime.Status) {
	if values == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E

	newUrl := resolver.Build(searchPath, values.Encode())
	resp, status := exchange.Get(newUrl, h)
	if !status.OK() {
		return nil, e.Handle(status, runtime.RequestId(h), searchLocation)
	}
	var buf []byte
	if override {
		buf = resultsCorrupt //resultsUnix
	} else {
		buf, status = runtime.ReadAll(resp.Body)
		if buf == nil || !status.OK() {
			return nil, e.Handle(status, runtime.RequestId(h), searchLocation)
		}
	}
	status = runtime.NewStatusOK()
	//status.ContentHeader().Add(http2.ContentType, "text/html")
	//status.ContentHeader().Add(http2.ContentType, "charset=ISO-8859-1")
	status.ContentHeader().Add(http2.ContentType, resp.Header.Get(http2.ContentType))
	//status.ContentHeader().Set(http2.ContentLength, fmt.Sprintf("%v", len(resultsGolang)))
	fmt.Printf("Search() results : [content-type:%v] [override:%v]\n", http.DetectContentType(buf), override)
	return buf, status
}
