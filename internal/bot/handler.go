package bot

import (
	"pixie/internal/pkg/json"
	"reflect"
)

func messageHandler(messageChain json.MessageChain) {
	go senderHandler(messageChain.Sender)

	for _, message := range messageChain.Messages {
		switch message.Type {
		case "source":
		case "at":
		case "face":
		case "plain":
			go plainHandler(message, messageChain)
		case "image":
		default:
		}
	}
	return
}

type plainHandlerMode int

const (
	trigger plainHandlerMode = 1
	echo    plainHandlerMode = 2
)

var mode plainHandlerMode = echo

func senderHandler(sender json.Sender) {
	if _, ok := reflect.TypeOf(&sender).FieldByName("Group"); !ok {
		return
	}
}

func plainHandler(inMessage json.Message, inMessageChain json.MessageChain) {
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
