package bot

import (
	"log"
	"pixie/internal/pkg/net"
)

func StartListen() {
	go func() {
		_, message, err := net.MessageConn.ReadMessage()
		if err != nil {
			log.Printf("read: %s", err)
		}
		messageRecvCh <- message
	}()

	go func() {
		_, event, err := net.EventConn.ReadMessage()
		if err != nil {
			log.Printf("read: %s", err)
		}
		eventRecvCh <- event
	}()
}
