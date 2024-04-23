package provider

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/http"
	"time"
)

var (
	agent *messaging.Agent
)

func init() {
	var err error
	agent, err = messaging.NewDefaultAgent(PkgPath, messageHandler, false)
	if err != nil {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, err)
	}
	agent.Run()
}

func messageHandler(msg *messaging.Message) {
	start := time.Now()
	switch msg.Event() {
	case messaging.StartupEvent:
		// Any processing for a Startup event would be here
		messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		// Any processing for a Ping event would be here
		messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
	}
}
