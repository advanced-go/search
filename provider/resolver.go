package provider

import (
	"embed"
	"github.com/advanced-go/core/runtime"
	uri2 "github.com/advanced-go/core/uri"
	"io/fs"
)

//go:embed resource/*
var f embed.FS

// http://localhost:8080/search?q=golang
// DEBUG : https://search.yahoo.com/search?p=golang
// TEST  : https://www.bing.com/search?q=C+Language
// STAGE : https://www.google.com/search?q=C%2B%2B
// PROD  : https://duckduckgo.com/?q=Pascal

const (
	searchPath     = "/search?%v"
	searchResource = "search"
	resultsPath    = "resource/results.html"
)

var (
	resolver = uri2.NewResolver()
	results  []byte
	err      error
)

func init() {
	resolver.SetOverrides([]runtime.Pair{{searchPath, "https://www.google.com/search?%v"}})
	results, err = fs.ReadFile(f, resultsPath)
}

/*
func readAuthorities(path string) ([]runtime.Pair, error) {
	var pairs []runtime.Pair

	buf, err := fs.ReadFile(f, path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &pairs)
	return pairs, err
}

func newValues(values url.Values) url.Values {
	if values == nil {
		values = make(url.Values)
		values.Add(queryArg, "")
		return values
	}
	if queryArg == defaultQueryArg {
		return values
	}
	q := values.Get(defaultQueryArg)
	values.Del(defaultQueryArg)
	values.Set(queryArg, q)
	return values
}


*/
