package json

import (
	"encoding/json"
	"log"
)

func Decode(data []byte, isMessage bool) (any, error) {
	var wsResp WsResp
	if err := json.Unmarshal(data, &wsResp); err != nil {
		log.Printf("json: unmarshaling failed: %s", err)
		return nil, err
	}
	content, err := json.Marshal(wsResp.Data)
	if err != nil {
		log.Printf("json: marshaling failed: %s", err)
		return nil, err
	}

	if isMessage {
		return decodeMessage(content)
	}
	return decodeEvent(content)
}

func decodeMessage(data []byte) (Message, error) {
	var messageChain Message
	if err := json.Unmarshal(data, &messageChain); err != nil {
		log.Printf("json: unmarshaling failed: %s", err)
		return messageChain, err
	}
	
	return messageChain, nil
}

func decodeEvent(data []byte) (Event, error) {
	var event Event
	if err := json.Unmarshal(data, &event); err != nil {
		log.Printf("json: unmarshaling failed: %s", err)
		return event, err
	}

	return event, nil
}
