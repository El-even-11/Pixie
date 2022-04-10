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

var sessionTypeMap map[int]string = map[int]string{
	0: "Friend",
	1: "Group",
}

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
	oMessage := json.Message{
		MessageChain: message.MessageChain,
		Target:       sess.number,
	}

	wsReq := json.WsReq{
		SyncId:  "0", // message synchronization
		Command: "send" + sessionTypeMap[int(sess.sesstype)] + "Message",
		Content: oMessage,
	}

	go func() {
		SendCh <- wsReq
	}()
}
