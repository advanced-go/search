package module

import "github.com/advanced-go/stdlib/core"

const (
	Authority = "github/advanced-go/search"
	Name      = "search"
	Version   = "1.1.1"
)

func Info() core.ModuleInfo {
	return core.ModuleInfo{
		Authority: Authority,
		Version:   Version,
		Name:      Name,
	}
}
