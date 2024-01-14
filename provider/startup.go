package provider

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
	"github.com/advanced-go/core/runtime"
	"time"
)

var (
	agent messaging.Agent
)

func init() {
	var status runtime.Status
	agent, status = messaging.NewDefaultAgent(PkgPath, messageHandler, false)
	if !status.OK() {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, status)
	}
	agent.Run()
}

func messageHandler(msg messaging.Message) {
	start := time.Now()
	switch msg.Event {
	case messaging.StartupEvent:
		status := runtime.NewStatusOK().SetDuration(time.Since(start))
		messaging.SendReply(msg, status)
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		messaging.SendReply(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
	}
}
