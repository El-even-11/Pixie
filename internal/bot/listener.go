package bot

import (
	"log"
	"pixie/internal/pkg/debug"
	"pixie/internal/pkg/net"
)

func StartListen() {
	go func() {
		_, message, err := net.MessageConn.ReadMessage()
		debug.DPrintf("recv: %s", message)
		if err != nil {
			log.Printf("read: %s", err)
		}
		MessageRecvCh <- message
	}()

	go func() {
		_, event, err := net.EventConn.ReadMessage()
		debug.DPrintf("recv: %s", event)
		if err != nil {
			log.Printf("read: %s", err)
		}
		EventRecvCh <- event
	}()
}
