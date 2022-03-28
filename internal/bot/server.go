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
			go serveMessage(m)
		case e := <-EventRecvCh:
			go serveEvent(e)
		}
	}
}

func serveMessage(data []byte) {
	messageChainItf, err := json.Decode(data, true)

	if err != nil {
		log.Printf("%s", err)
		return nil
	}
	debug.DPrintf("decode: %v", messageChainItf)

	messageChain, ok := messageChainItf.(json.MessageChain)
	if !ok {
		panic("message chain type error!")
	}

	wsReqs, err := messageHandler(messageChain)
	if err != nil {
		return nil
	}

	messages := make([]Message, len(wsReqs))
	for i, wsReq := range wsReqs {
		bytes, err := json.Encode(wsReq)
		if err != nil {
			messages[i] = Message{
				Empty: true,
			}
			continue
		}
		messages[i] = Message{
			Empty: false,
			Data:  bytes,
			Done:  make(chan struct{}),
		}
	}

	return messages
}

func serveEvent(data []byte) {
	eventItf, err := json.Decode(data, false)
	if err != nil {
		log.Printf("%s", err)
		return
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
		return
	}

	wsReq, err := eventHandler(event)
	if err != nil {
		return
	}

	wsReqBytes, err := json.Encode(wsReq)
	if err != nil {
		return
	}

	return Event{
		Empty: false,
		Done:  make(chan struct{}),
		Data:  wsReqBytes,
	}
}
