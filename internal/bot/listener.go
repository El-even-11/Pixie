package bot

import (
	"pixie/internal/pkg/log"
	"pixie/internal/pkg/net"
)

func Listener() {
	go func() {
		for {
			_, message, err := net.MessageConn.ReadMessage()
			log.Log("recv m: %s", message)
			if err != nil {
				log.Log("read: %s", err)
				continue
			}
			go func() {
				MessageBytesRecvCh <- message
			}()
		}
	}()

	go func() {
		for {
			_, event, err := net.EventConn.ReadMessage()
			log.Log("recv e: %s", event)
			if err != nil {
				log.Log("read: %s", err)
				continue
			}
			go func() {
				EventBytesRecvCh <- event
			}()
		}
	}()
}
