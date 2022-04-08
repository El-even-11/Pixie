package bot

import (
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/log"
)

func MessageDecoder(data []byte) {
	message, err := json.DecodeMessage(data)
	if err != nil {
		log.Log("%s", err)

	}
	MessageRecvCh <- message
}

func EventDecoder(data []byte) {
	event, err := json.DecodeEvent(data)
	if err != nil {
		log.Log("%s", err)

	}
	EventRecvCh <- event
}
