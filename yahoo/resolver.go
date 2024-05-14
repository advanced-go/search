package yahoo

import (
	"github.com/advanced-go/stdlib/uri"
)

// http://localhost:8080/search?q=golang
// DEBUG : https://search.yahoo.com/search?p=golang
// TEST  : https://www.bing.com/search?q=C+Language
// STAGE : https://www.google.com/search?q=C%2B%2B
// PROD  : https://duckduckgo.com/?q=Pascal

const (
	searchPath = "/search?%v"
)

var (
	resolver = uri.NewResolver()
)

func init() {
	resolver.SetTemplates([]uri.Attr{{searchPath, "https://search.yahoo.com/search?%v"}})
}
