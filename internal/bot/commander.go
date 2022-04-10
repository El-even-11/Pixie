package bot

import (
	"pixie/internal/pkg/json"
	"strings"
)

func (sess *session) commandHandler(command string, iMessage json.Message) {
	paras := strings.Split(command, " ")
	switch paras[0] {
	case "/sleep":
		sess.mode = Sleep
	case "/wake":
		if sess.mode == Sleep {
			sess.mode = Trigger
		}
	case "/echo":
		sess.mode = Echo
	case "/trigger":
		sess.mode = Trigger
	default:
	}
}
