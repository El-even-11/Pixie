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
