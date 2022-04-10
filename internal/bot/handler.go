package bot

// import (
// 	"pixie/internal/pkg/json"
// 	"reflect"
// 	"strings"
// )

// func messageHandler(message json.Message) {
// 	go senderHandler(message.Sender)

// 	for _, messageItem := range message.MessageChain {
// 		switch messageItem.Type {
// 		case "Source":
// 		case "At":
// 		case "Face":
// 		case "Plain":
// 			go textHandler(messageItem, message)
// 		case "Image":
// 		default:
// 		}
// 	}
// }

// func senderHandler(sender json.Sender) {
// 	if _, ok := reflect.TypeOf(sender).FieldByName("Group"); !ok {
// 		return
// 	}
// }

// func textHandler(inMessage json.MessageItem, inMessageChain json.Message) {
// 	if strings.HasPrefix(inMessage.Text, "/") {
// 		go commandHandler(inMessage, inMessageChain)
// 		return
// 	}

// 	wsReq := json.WsReq{
// 		SyncId: "1",
// 	}

// 	outMessageChain := json.Message{
// 		MessageChain: make([]json.MessageItem, 0),
// 	}

// 	switch inMessageChain.Type {
// 	case "GroupMessage":
// 		wsReq.Command = "sendGroupMessage"
// 		outMessageChain.Target = inMessageChain.Sender.Group.ID
// 	case "FriendMessage":
// 		wsReq.Command = "sendFriendMessage"
// 		outMessageChain.Target = inMessageChain.Sender.ID
// 	default:
// 		return
// 	}

// 	switch mode {
// 	case echo:
// 		outMessageChain.MessageChain = append(outMessageChain.MessageChain, json.BuildMessage([]string{"Plain"}, []string{inMessage.Text})...)
// 	case trigger:
// 	default:
// 		panic("unknown plain handler mode")
// 	}

// 	wsReq.Content = outMessageChain
// 	go func() {
// 		SendCh <- wsReq
// 	}()
// }

// func eventHandler(event json.Event) {

// }
