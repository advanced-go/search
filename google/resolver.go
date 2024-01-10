package google

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/advanced-go/core/runtime"
	uri2 "github.com/advanced-go/core/uri"
	"os"
)

const (
	searchTag  = "search"
	searchPath = "/search"

	debugPath = "file://[cwd]/resource/authorities-debug.json"
	testPath  = "file://[cwd]/resource/authorities-test.json"
	stagePath = "file://[cwd]/resource/authorities-stage.json"
	prodPath  = "file://[cwd]/resource/authorities-prod.json"
)

var (
	resolver       = uri2.NewResolver()
	testAuthority  []uri2.Attr
	stageAuthority []uri2.Attr
	prodAuthority  []uri2.Attr
	initError      error
)

func init() {
	debug, ok := initAuthorities(debugPath)
	if !ok {
		return
	}
	resolver.SetAuthorities(debug)

	testAuthority, ok = initAuthorities(testPath)
	if !ok {
		return
	}
	stageAuthority, ok = initAuthorities(stagePath)
	if !ok {
		return
	}
	prodAuthority, ok = initAuthorities(prodPath)
}

func initResolver() error {
	var ok bool
	var attrs []uri2.Attr

	if runtime.IsTestEnvironment() {
		attrs, ok = initAuthorities(testPath)
	} else {
		if runtime.IsStageEnvironment() {
			attrs, ok = initAuthorities(stagePath)
		} else {
			attrs, ok = initAuthorities(prodPath)
		}
	}
	if ok {
		resolver.SetAuthorities(attrs)
	}
	return initError
}

func readAuthorities(path string) ([]uri2.Attr, error) {
	var attrs []uri2.Attr

	fname := uri2.FileName(path)
	buf, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &attrs)
	return attrs, err
}

func initAuthorities(path string) ([]uri2.Attr, bool) {
	authorities, err := readAuthorities(path)
	if err != nil {
		initError = errors.New(fmt.Sprintf("%v : %v", err, path))
		return nil, false
	}
	return authorities, true
}
