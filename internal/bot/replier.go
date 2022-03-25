package bot

import (
	"log"
	"pixie/internal/pkg/net"

	"github.com/gorilla/websocket"
)

func StartReply() {
	go func() {
		for {
			select {
			case data := <-messageSendCh:
				err := net.MessageConn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Printf("write: %s", err)
				}
			case data := <-eventSendCh:
				err := net.EventConn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Printf("write: %s", err)
				}
			}
		}
	}()
}
