package bot

import (
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/log"
)

func MessageDecoder(data []byte) {
	message, err := json.DecodeMessage(data)
	if err != nil {
		log.Log("decoder: decode fail, %s", err)
		return
	}
	MessageRecvCh <- message
}

func EventDecoder(data []byte) {
	event, err := json.DecodeEvent(data)
	if err != nil {
		log.Log("decoder: decode fail, %s", err)
		return
	}
	EventRecvCh <- event
}
