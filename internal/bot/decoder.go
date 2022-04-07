package bot

import (
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/log"
)

func Decoder() {
	go func() {
		for {
			select {
			case m := <-MessageBytesRecvCh:
				messageItf, err := json.Decode(m, true)
				if err != nil {
					log.Log("%s", err)
					break
				}

				message, ok := messageItf.(json.Message)
				if !ok {
					panic("message type error!")
				}
				go func() {
					MessageRecvCh <- message
				}()

			case e := <-EventBytesRecvCh:
				eventItf, err := json.Decode(e, false)
				if err != nil {
					log.Log("%s", err)
					break
				}

				event, ok := eventItf.(json.Event)
				if !ok {
					panic("event type error!")
				}

				go func() {
					EventRecvCh <- event
				}()
			}
		}
	}()
}
