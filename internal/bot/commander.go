package bot

import (
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/log"
	"strings"
)

func (sess *session) commandHandler(command string, iMessage json.Message) {
	paras := strings.Split(command, " ")
	switch paras[0] {
	case "/sleep":
		sess.doSleep()
	case "/wake":
		sess.doWake()
	case "/echo":
		sess.doEcho()
	case "/trigger":
		sess.doTrigger()
	default:
	}
}

func (sess *session) doSleep() {
	sess.modeLock.Lock()
	log.Log("sess %d mode changed: %s -> %s", sess.number, sessionModeMap[int(sess.mode)], "Sleep")
	sess.mode = Sleep
	sess.modeLock.Unlock()
}

func (sess *session) doWake() {
	sess.modeLock.Lock()
	if sess.mode == Sleep {
		log.Log("sess %d mode changed: %s -> %s", sess.number, sessionModeMap[int(sess.mode)], "Trigger")
		sess.mode = Trigger
	}
	sess.modeLock.Unlock()
}

func (sess *session) doEcho() {
	sess.modeLock.Lock()
	log.Log("sess %d mode changed: %s -> %s", sess.number, sessionModeMap[int(sess.mode)], "Echo")
	sess.mode = Echo
	sess.modeLock.Unlock()
}

func (sess *session) doTrigger() {
	sess.modeLock.Lock()
	log.Log("sess %d mode changed: %s -> %s", sess.number, sessionModeMap[int(sess.mode)], "Trigger")
	sess.mode = Trigger
	sess.modeLock.Unlock()
}
