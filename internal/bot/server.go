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
			respMessages := serveMessage(m)
			for _, message := range respMessages {
				if message.Empty {
					continue
				}
				go func() {
					MessageSendCh <- message
				}()
				<-message.Done
			}

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

func serveMessage(data []byte) []Message {
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
	for _, wsReq := range wsReqs {
		bytes, err := json.Encode(wsReq)
		if err != nil {
			messages = append(messages,
				Message{
					Empty: true,
				},
			)
			continue
		}
		messages = append(messages,
			Message{
				Empty: false,
				Data:  bytes,
				Done:  make(chan struct{}),
			},
		)
	}

	return messages
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
