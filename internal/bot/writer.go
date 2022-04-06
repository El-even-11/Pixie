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
			select {
			case data := <-BytesSendCh:
				log.Log("Write: %s", data)
				err := net.MessageConn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Log("write fail: %s", err)
				}
				// time.Sleep(INTERVAL)
			case <-SleepCh:
			FORLOOP:
				for {
					select {
					case <-BytesSendCh:
						// while sleeping, ignore the sending messages
					case <-SleepCh:
						// ignore multi sleep commands
					case <-WakeCh:
						log.Log("wake!")
						break FORLOOP
					}
				}
			case <-WakeCh:
				// while awake, ignore wake commands
			}
		}
	}()
}
