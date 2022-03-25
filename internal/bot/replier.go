package bot

import (
	"log"
	"pixie/internal/pkg/debug"
	"pixie/internal/pkg/net"

	"github.com/gorilla/websocket"
)

func StartReply() {
	go func() {
		for {
			select {
			case data := <-MessageSendCh:

				debug.DPrinf("Write: %s", data)

				err := net.MessageConn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Printf("write fail: %s", err)
				}
			case data := <-EventSendCh:

				debug.DPrinf("Write: %s", data)

				err := net.EventConn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Printf("write fail: %s", err)
				}
			}
		}
	}()
}
