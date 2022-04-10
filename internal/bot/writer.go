package bot

import (
	"pixie/internal/pkg/log"
	"pixie/internal/pkg/net"

	"github.com/gorilla/websocket"
)

// pause between messages to prevent packet loss
// const INTERVAL = time.Millisecond * 500

func Writer() {
	go func() {
		for {
			data := <-BytesSendCh
			log.Log("Write: %s", data)
			err := net.MessageConn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Log("write fail: %s", err)
			}
		}
	}()
}
