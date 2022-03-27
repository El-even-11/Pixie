package bot

import (
	"log"
	"pixie/internal/pkg/debug"
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
	messageChain, err := json.Decode(data, true)
	if err != nil {
		log.Printf("%s", err)
	}
	debug.DPrintf("decode: %v", messageChain)
	return Message{
		Empty: true,
	}
}

func serveEvent(data []byte) Event {
	event, err := json.Decode(data, false)
	if err != nil {
		log.Printf("%s", err)
	}

	debug.DPrintf("decode: %v", event)
	return Event{
		Empty: true,
	}
}
