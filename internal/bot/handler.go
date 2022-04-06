package bot

import (
	"pixie/internal/pkg/debug"
	"pixie/internal/pkg/json"
	"reflect"
	"strings"
)

func messageHandler(messageChain json.MessageChain) {
	go senderHandler(messageChain.Sender)

	debug.DPrintf("%d", len(messageChain.Messages))
	for _, message := range messageChain.Messages {
		switch message.Type {
		case "Source":
		case "At":
		case "Face":
		case "Plain":
			go textHandler(message, messageChain)
		case "Image":
		default:
		}
	}
}

type plainHandlerMode int

const (
	trigger plainHandlerMode = 1
	echo    plainHandlerMode = 2
)

var mode plainHandlerMode = echo

func senderHandler(sender json.Sender) {
	if _, ok := reflect.TypeOf(sender).FieldByName("Group"); !ok {
		return
	}
}

func textHandler(inMessage json.Message, inMessageChain json.MessageChain) {
	if strings.HasPrefix(inMessage.Text, "/") {
		go commandHandler(inMessage, inMessageChain)
		return
	}

	wsReq := json.WsReq{
		SyncId: "1",
	}

	outMessageChain := json.MessageChain{
		Messages: make([]json.Message, 0),
	}

	switch inMessageChain.Type {
	case "GroupMessage":
		wsReq.Command = "sendGroupMessage"
		outMessageChain.Target = inMessageChain.Sender.Group.ID
	case "FriendMessage":
		wsReq.Command = "sendFriendMessage"
		outMessageChain.Target = inMessageChain.Sender.ID
	default:
		return
	}

	switch mode {
	case echo:
		outMessageChain.Messages = append(outMessageChain.Messages, json.BuildMessage([]string{"Plain"}, []string{inMessage.Text})...)
	case trigger:
	default:
		panic("unknown plain handler mode")
	}

	wsReq.Content = outMessageChain
	go func() {
		SendCh <- wsReq
	}()
}

func eventHandler(event json.Event) {

}
