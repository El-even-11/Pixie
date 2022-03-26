package bot

import (
	"log"
	"pixie/internal/pkg/debug"
	"pixie/internal/pkg/net"
	"time"

	"github.com/gorilla/websocket"
)

// pause between messages to prevent packet loss
const INTERVAL = time.Millisecond * 500

func StartReply() {
	go func() {
		for {
			select {
			case message := <-MessageSendCh:

				debug.DPrinf("Write: %s", message.Data)

				err := net.MessageConn.WriteMessage(websocket.TextMessage, message.Data)
				message.Done <- struct{}{}

				if err != nil {
					log.Printf("write fail: %s", err)
				}
			case event := <-EventSendCh:

				debug.DPrinf("Write: %s", event.Data)

				err := net.EventConn.WriteMessage(websocket.TextMessage, event.Data)
				event.Done <- struct{}{}

				if err != nil {
					log.Printf("write fail: %s", err)
				}
			}
			time.Sleep(INTERVAL)
		}
	}()
}
