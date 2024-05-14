package http

import (
	"github.com/advanced-go/search/provider"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

const (
	ModulePath = "github/advanced-go/search"
)

func Exchange(r *http.Request) (*http.Response, *core.Status) {
	return provider.HttpExchange(r)
}
