package provider

import (
	uri2 "github.com/advanced-go/core/uri"
)

// http://localhost:8080/search?q=golang
// DEBUG : https://search.yahoo.com/search?p=golang
// TEST  : https://www.bing.com/search?q=C+Language
// STAGE : https://www.google.com/search?q=C%2B%2B
// PROD  : https://duckduckgo.com/?q=Pascal

const (
	searchPath     = "/search?%v"
	searchResource = "search"
)

var (
	resolver = uri2.NewResolver()
)

func init() {
	resolver.SetTemplates([]uri2.Pair{{searchPath, "https://www.google.com/search?%v"}})
}
