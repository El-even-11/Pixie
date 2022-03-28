package bot

import (
	"errors"
	"pixie/internal/pkg/json"
)

func messageHandler(messageChain json.MessageChain) ([]json.WsReq, error) {

	wsReqArr := make([]json.WsReq, 0)
	var commandType string

	switch messageChain.Type {
	case "FriendMessage":
		commandType = "sendFriendMessage"
	case "GroupMessage":
		commandType = "sendGroupMessage"
	default:
		return nil, nil
	}

	for _, message := range messageChain.Messages {
		switch message.Type {
		case "source":
		case "at":
		case "face":
		case "plain":
			outMessageChain := plainHandler(message, messageChain)
			wsReq := json.WsReq{
				SyncId: "1",
			}

		case "image":
		default:
			return nil, errors.New("unknown message type")
		}
	}
	return nil, nil
}

type plainHandlerMode int

const (
	trigger plainHandlerMode = 1
	echo    plainHandlerMode = 2
)

var mode plainHandlerMode = echo

func plainHandler(inMessage json.Message, inMessageChain json.MessageChain) json.MessageChain {
	outMessageChain := json.MessageChain{}
	switch mode {
	case echo:
	case trigger:
	default:
		panic("unknown plain handler mode")
	}
}

func eventHandler(event json.Event) {

}
