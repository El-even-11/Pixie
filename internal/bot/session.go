package bot

import (
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/log"
	"strings"
	"sync"
)

type sessionType int

const (
	Friend sessionType = 0
	Group  sessionType = 1
)

type sessionMode int

const (
	Sleep   sessionMode = 0
	Trigger sessionMode = 1
	Echo    sessionMode = 2
)

var sessionModeMap map[int]string = map[int]string{
	0: "Sleep",
	1: "Trigger",
	2: "Echo",
}

type session struct {
	sesstype  sessionType
	mode      sessionMode
	number    int64
	messageCh chan json.Message
	eventCh   chan json.Event
	modeLock  sync.Mutex
}

func (sess *session) serve() {
	for {
		if sess.mode == Sleep {
			log.Log("sleeping")
			select {
			case m := <-sess.messageCh:
				for _, messageItem := range m.MessageChain {
					if messageItem.Type == "Plain" && strings.HasPrefix(messageItem.Text, "/") {
						// searching for command
						sess.commandHandler(messageItem.Text, m)
					}
				}
			case <-sess.eventCh:
			}
			continue
		}
		select {
		case m := <-sess.messageCh:
			go sess.messageHandler(m)
		case e := <-sess.eventCh:
			go sess.eventHandler(e)
		}
	}
}

func (sess *session) messageHandler(message json.Message) {
	if message.Type == "GroupMessage" {
		go sess.senderHandler(message.Sender)
	}

	for _, messageItem := range message.MessageChain {
		if messageItem.Type == "Plain" && strings.HasPrefix(messageItem.Text, "/") {
			// searching for command
			sess.commandHandler(messageItem.Text, message)
			return
		}
	}

	// no command
	switch sess.mode {
	case Echo:
		sess.echo(message)
	case Trigger:
	}
}

func (sess *session) eventHandler(event json.Event) {

}

func (sess *session) senderHandler(sender json.Sender) {

}

func (sess *session) echo(message json.Message) {
	wsReq := json.WsReq{
		SyncId: "0", // message synchronization
	}

	oMessage := json.Message{
		MessageChain: message.MessageChain,
	}

	switch sess.sesstype {
	case Friend:
		wsReq.Command = "sendFriendMessage"
		oMessage.Target = message.Sender.ID
	case Group:
		wsReq.Command = "sendGroupMessage"
		oMessage.Target = message.Sender.Group.ID
	default:
		return
	}

	wsReq.Content = oMessage

	go func() {
		SendCh <- wsReq
	}()
}
