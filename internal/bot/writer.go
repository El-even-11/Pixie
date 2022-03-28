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

func Writer() {
	go func() {
		for {
			select {
			case message := <-MessageBytesSendCh:
				debug.DPrintf("Write: %s", message)
				err := net.MessageConn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Printf("write fail: %s", err)
				}
			case event := <-EventBytesSendCh:
				debug.DPrintf("Write: %s", event)
				err := net.EventConn.WriteMessage(websocket.TextMessage, event)
				if err != nil {
					log.Printf("write fail: %s", err)
				}
			}
			time.Sleep(INTERVAL)
		}
	}()
}
