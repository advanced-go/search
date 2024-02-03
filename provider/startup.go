package provider

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
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

func messageHandler(msg messaging.Message) {
	start := time.Now()
	switch msg.Event {
	case messaging.StartupEvent:
		//status := runtime.NewStatusOK().SetDuration(time.Since(start))
		messaging.SendReply(msg, messaging.NewStatusDuration(http.StatusOK, time.Since(start)))
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		messaging.SendReply(msg, messaging.NewStatusDuration(http.StatusOK, time.Since(start)))
	}
}
