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
			data := <-BytesSendCh
			debug.DPrintf("Write: %s", data)
			err := net.MessageConn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Printf("write fail: %s", err)
			}
			time.Sleep(INTERVAL)
		}
	}()
}
