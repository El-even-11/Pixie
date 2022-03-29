package bot

import (
	"log"
	"pixie/internal/pkg/debug"
	"pixie/internal/pkg/net"

	"github.com/gorilla/websocket"
)

// pause between messages to prevent packet loss
// const INTERVAL = time.Millisecond * 500

func Writer() {
	go func() {
		for {
			select {
			case data := <-BytesSendCh:
				debug.DPrintf("Write: %s", data)
				err := net.MessageConn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Printf("write fail: %s", err)
				}
				// time.Sleep(INTERVAL)
			case <-SleepCh:
				for {
					select {
					case <-BytesSendCh:
						// while sleeping, ignore the sending messages
					case <-WakeCh:
						break
					}
				}
			}
		}
	}()
}
