package bot

import (
	"log"
	"pixie/internal/pkg/debug"
	"pixie/internal/pkg/net"
)

func StartListen() {
	go func() {
		for {
			_, message, err := net.MessageConn.ReadMessage()
			debug.DPrintf("recv: %s", message)
			if err != nil {
				log.Printf("read: %s", err)
			}
			go func() {
				MessageRecvCh <- message
			}()
		}
	}()

	go func() {
		for {
			_, event, err := net.EventConn.ReadMessage()
			debug.DPrintf("recv: %s", event)
			if err != nil {
				log.Printf("read: %s", err)
			}
			go func() {
				EventRecvCh <- event
			}()
		}
	}()
}
