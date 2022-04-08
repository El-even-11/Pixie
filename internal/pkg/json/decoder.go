package json

import (
	"encoding/json"
	"log"
)

func DecodeMessage(data []byte) (Message, error) {
	var wsResp WsResp
	if err := json.Unmarshal(data, &wsResp); err != nil {
		log.Printf("json: unmarshaling failed: %s", err)
		return Message{}, err
	}
	content, err := json.Marshal(wsResp.Data)
	if err != nil {
		log.Printf("json: marshaling failed: %s", err)
		return Message{}, err
	}

	var message Message
	if err := json.Unmarshal(content, &message); err != nil {
		log.Printf("json: unmarshaling failed: %s", err)
		return Message{}, err
	}

	return message, nil
}

func DecodeEvent(data []byte) (Event, error) {
	var wsResp WsResp
	if err := json.Unmarshal(data, &wsResp); err != nil {
		log.Printf("json: unmarshaling failed: %s", err)
		return Event{}, err
	}
	content, err := json.Marshal(wsResp.Data)
	if err != nil {
		log.Printf("json: marshaling failed: %s", err)
		return Event{}, err
	}

	var event Event
	if err := json.Unmarshal(content, &event); err != nil {
		log.Printf("json: unmarshaling failed: %s", err)
		return Event{}, err
	}

	return event, nil
}
