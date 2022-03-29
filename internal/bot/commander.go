package bot

import (
	"pixie/internal/pkg/json"
	"strings"
)

func commandHandler(inMessage json.Message, inMessageChain json.MessageChain) {
	paras := strings.Split(inMessage.Text, " ")
	switch paras[0] {
	case "/sleep":
		sleep()
	case "/wake":
		wake()
	default:
	}
}

func sleep() {
	SleepCh <- struct{}{}
}

func wake() {
	WakeCh <- struct{}{}
}
