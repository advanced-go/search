package provider

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/advanced-go/core/runtime"
	uri2 "github.com/advanced-go/core/uri"
	"io/fs"
	"net/url"
)

//go:embed resource/*
var f embed.FS

// DEBUG : https://search.yahoo.com/search?p=golang
// TEST  : https://www.bing.com/search?q=C+Language
// STAGE : https://www.google.com/search?q=C%2B%2B
// PROD  : https://duckduckgo.com/?q=Pascal

const (
	searchTag       = "{SEARCH}"
	defaultPath     = "/search?%v"
	duckPath        = "/?%v"
	defaultQueryArg = "q"
	yahooQueryArg   = "p"

	debugPath = "resource/authorities-debug.json"
	testPath  = "resource/authorities-test.json"
	stagePath = "resource/authorities-stage.json"
	prodPath  = "resource/authorities-prod.json"
)

var (
	resolver   = uri2.NewResolver()
	initError  error
	searchPath = defaultPath
	queryArg   = defaultQueryArg
)

func init() {
	debug, ok := initAuthorities(debugPath)
	if !ok {
		return
	}
	// Debug has a different query arg
	queryArg = yahooQueryArg
	resolver.SetAuthorities(debug)
}

func initResolver() error {
	var ok bool
	var attrs []uri2.KV

	if initError != nil || runtime.IsDebugEnvironment() {
		return initError
	}
	queryArg = defaultQueryArg
	if runtime.IsTestEnvironment() {
		attrs, ok = initAuthorities(testPath)
	} else {
		if runtime.IsStageEnvironment() {
			attrs, ok = initAuthorities(stagePath)
		} else {
			// production has no path
			searchPath = duckPath
			attrs, ok = initAuthorities(prodPath)
		}
	}
	if ok {
		resolver.SetAuthorities(attrs)
	}
	return initError
}

func readAuthorities(path string) ([]uri2.KV, error) {
	var attrs []uri2.KV

	buf, err := fs.ReadFile(f, path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &attrs)
	return attrs, err
}

func initAuthorities(path string) ([]uri2.KV, bool) {
	authorities, err := readAuthorities(path)
	if err != nil {
		initError = errors.New(fmt.Sprintf("%v : %v", err, path))
		return nil, false
	}
	return authorities, true
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
