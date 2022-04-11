package bot

import (
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/log"
	"pixie/internal/pkg/redis"
	"strings"
)

func (sess *session) commandHandler(command string, iMessage json.Message) {
	command = strings.TrimSpace(command)
	paras := strings.Fields(command)
	switch paras[0] {
	case "/sleep":
		sess.doSleep()
	case "/wake":
		sess.doWake()
	case "/echo":
		sess.doEcho()
	case "/trigger":
		sess.doTrigger()
	case "/addtext":
		sess.doAddTextTrigger(paras[1:])
	case "/addimg":
		sess.doAddImageTrigger(paras[1:])
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

func (sess *session) doAddTextTrigger(paras []string) {
	if len(paras) < 2 {
		wsReq := json.BuildWsReq(
			sess.number,
			"send"+sessionTypeMap[int(sess.sesstype)]+"Message",
			[]string{"Plain"},
			[]string{"too few parameters!"},
		)
		SendCh <- wsReq
		return
	}

	err := redis.SAdd(paras[0], paras[1:])
	if err != nil {
		log.Log("redis: sadd failed, %s", err)
		wsReq := json.BuildWsReq(
			sess.number,
			"send"+sessionTypeMap[int(sess.sesstype)]+"Message",
			[]string{"Plain"},
			[]string{err.Error()},
		)
		SendCh <- wsReq
		return
	}

	wsReq := json.BuildWsReq(
		sess.number,
		"send"+sessionTypeMap[int(sess.sesstype)]+"Message",
		[]string{"Plain"},
		[]string{"add text trigger success"},
	)
	SendCh <- wsReq
}

func (sess *session) doAddImageTrigger(paras []string) {

}
