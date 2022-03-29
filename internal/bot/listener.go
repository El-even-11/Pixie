package bot

import (
	"log"
	"pixie/internal/pkg/debug"
	"pixie/internal/pkg/net"
)

func Listener() {
	go func() {
		for {
			_, message, err := net.MessageConn.ReadMessage()
			debug.DPrintf("recv m: %s", message)
			if err != nil {
				log.Printf("read: %s", err)
			}
			go func() {
				MessageBytesRecvCh <- message
			}()
		}
	}()

	go func() {
		for {
			_, event, err := net.EventConn.ReadMessage()
			debug.DPrintf("recv e: %s", event)
			if err != nil {
				log.Printf("read: %s", err)
			}
			go func() {
				EventBytesRecvCh <- event
			}()
		}
	}()
}
