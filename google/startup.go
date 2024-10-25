package google

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/host"
	"github.com/advanced-go/common/messaging"
	"net/http"
	"time"
)

func init() {
	a, err := host.RegisterControlAgent(PkgPath, messageHandler)
	if err != nil {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, err)
	}
	a.Run()
}

func messageHandler(msg *messaging.Message) {
	start := time.Now()
	switch msg.Event() {
	case messaging.StartupEvent:
		// Any processing for a Startup event would be here
		messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		// Any processing for a Shutdown/Ping event would be here
		messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
	}
}
