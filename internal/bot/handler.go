package bot

import (
	"errors"
	"pixie/internal/pkg/json"
)

func messageHandler(messageChain json.MessageChain) ([]json.WsReqData, error) {

	for _, message := range messageChain.Messages {
		switch message.Type {
		case "source":
		case "at":
		case "face":
		case "plain":

		case "image":
		default:
			return nil, errors.New("unknown message type")
		}
	}
	return nil, nil
}

func plainHandler(message json.Message, messageChain json.MessageChain) (json.MessageChain, error) {

}
