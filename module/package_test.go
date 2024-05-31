package module

import (
	"fmt"
	"runtime/debug"
)

func ExampleDebugInfo() {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range bi.Deps {
			fmt.Printf("Dep: %+v\n", dep)
		}
		//fmt.Printf("test: DebugInfo() -> [%v]\n", info)
	}

	//Output:
	//fail

}
