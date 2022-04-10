package bot

import (
	"pixie/internal/pkg/json"
	"strings"
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

type session struct {
	sesstype  sessionType
	mode      sessionMode
	number    int64
	messageCh chan json.Message
	eventCh   chan json.Event
}

func (sess *session) serve() {

	for {
		if sess.mode == Sleep {
			select {
			case m := <-sess.messageCh:
				for _, messageItem := range m.MessageChain {
					if messageItem.Type == "Plain" && strings.HasPrefix(messageItem.Text, "/") {
						// waiting for wake command
						sess.commandHandler(messageItem.Text, m)
					}
				}
			case <-sess.eventCh:
			}
			continue
		}
		select {
		case m := <-sess.messageCh:
			sess.messageHandler(m)
		case e := <-sess.eventCh:
			sess.eventHandler(e)
		}
	}
}

func (sess *session) messageHandler(message json.Message) {
	go sess.senderHandler(message.Sender)

	switch sess.mode {
	case Echo:
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

	go func() {
		SendCh <- wsReq
	}()
}
