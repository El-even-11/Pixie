package bot

import (
	"pixie/internal/pkg/json"
)

func Serve() {
	for {
		select {
		case m := <-MessageRecvCh:
			respMessage := serveMessage(m)
			if respMessage.Empty {
				continue
			}
			go func() {
				MessageSendCh <- respMessage
			}()
			<-respMessage.Done
		case e := <-EventRecvCh:
			respEvent := serveEvent(e)
			if respEvent.Empty {
				continue
			}
			go func() {
				EventSendCh <- respEvent
			}()
			<-respEvent.Done
		}
	}
}

func serveMessage(data []byte) Message {
	json.Decode(data, true)
	return Message{
		Empty: true,
	}
}

func serveEvent(data []byte) Event {
	json.Decode(data, false)
	return Event{
		Empty: true,
	}
}
