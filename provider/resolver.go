package provider

import (
	uri2 "github.com/advanced-go/core/uri"
)

/*
//go:embed resource/*
var f embed.FS

*/

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
	resolver.SetOverrides([]uri2.Pair{{searchPath, "https://www.google.com/search?%v"}})
}

/*
resultsPath    = "resource/results.html"
	sloPath        = "resource/query-slo.html"
	winPath        = "resource/q-golang-2.html"
	unixPath       = "resource/q-golang-1.html"
	corruptPath    = "resource/q-golang-corrupt.html"
resultsGolang  []byte
	resultsSLO     []byte
	resultsWin     []byte
	resultsUnix    []byte
	resultsCorrupt []byte
	resultsErr     error
resultsGolang, resultsErr = fs.ReadFile(f, resultsPath)
	resultsSLO, resultsErr = fs.ReadFile(f, sloPath)
	resultsWin, resultsErr = fs.ReadFile(f, winPath)
	resultsUnix, resultsErr = fs.ReadFile(f, unixPath)
	resultsCorrupt, resultsErr = fs.ReadFile(f, corruptPath)
	if resultsErr != nil {
		fmt.Printf("results error: %v\n", resultsErr)
	}

r := bytes.NewReader(resultsCorrupt)
	zr, err := gzip.NewReader(r)
	if err != nil {
		fmt.Printf("gzip error: %v\n", err)
		return
	}
	buf, status := runtime.ReadAll(io.NopCloser(zr))
	if !status.OK() {
		fmt.Printf("gzip read error: %v\n", status)
		return
	}
	s := string(buf)
	if len(s) > 0 {

	}
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
