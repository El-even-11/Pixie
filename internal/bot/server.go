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
	messageChainItf, err := json.Decode(data, true)

	if err != nil {
		log.Printf("%s", err)
		return Message{
			Empty: true,
		}
	}
	debug.DPrintf("decode: %v", messageChainItf)

	messageChain, ok := messageChainItf.(json.MessageChain)
	if !ok {
		panic("message chain type error!")
	}

	var messageHandler func(json.MessageChain) (json.WsReqData, error)

	switch messageChain.Type {
	case "FriendMessage":
		messageHandler = friendMessageHandler
	case "GroupMessage":
		messageHandler = groupMessageHandler
	default:
		log.Printf("unknown message chain type")
		return Message{
			Empty: true,
		}
	}

	wsReq, err := messageHandler(messageChain)
	if err != nil {
		return Message{
			Empty: true,
		}
	}

	wsReqBytes, err := json.Encode(wsReq)
	if err != nil {
		return Message{
			Empty: true,
		}
	}

	return Message{
		Empty: false,
		Done:  make(chan struct{}),
		Data:  wsReqBytes,
	}
}

func serveEvent(data []byte) Event {
	eventItf, err := json.Decode(data, false)
	if err != nil {
		log.Printf("%s", err)
		return Event{
			Empty: true,
		}
	}

	debug.DPrintf("decode: %v", eventItf)

	event, ok := eventItf.(json.Event)
	if !ok {
		panic("message chain type error!")
	}

	var eventHandler func(json.Event) (json.WsReqData, error)

	switch event.Type {
	default:
		log.Printf("unknown event type")
		return Event{
			Empty: true,
		}
	}

	wsReq, err := eventHandler(event)
	if err != nil {
		return Event{
			Empty: true,
		}
	}

	wsReqBytes, err := json.Encode(wsReq)
	if err != nil {
		return Event{
			Empty: true,
		}
	}

	return Event{
		Empty: false,
		Done:  make(chan struct{}),
		Data:  wsReqBytes,
	}
}
