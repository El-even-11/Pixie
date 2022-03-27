package json

import (
	"encoding/json"
	"log"
)

func Decode(data []byte, isMessage bool) (any, error) {
	var wsRespData WsRespData
	if err := json.Unmarshal(data, &wsRespData); err != nil {
		log.Printf("json: unmarshaling failed: %s", err)
		return nil, err
	}
	content, err := json.Marshal(wsRespData.Data)
	if err != nil {
		log.Printf("json: marshaling failed: %s", err)
		return nil, err
	}

	if isMessage {
		return decodeMessage(content)
	}
	return decodeEvent(content)
}

func decodeMessage(data []byte) (MessageChain, error) {
	var messageChain MessageChain
	if err := json.Unmarshal(data, &messageChain); err != nil {
		log.Printf("json: unmarshaling failed: %s", err)
		return messageChain, err
	}

	return messageChain, nil
}

func decodeEvent(data []byte) (any, error) {
	return nil, nil
}
