package bot

import (
	"pixie/internal/pkg/json"
	"time"
)

type sessType int

const (
	Friend sessType = 0
	Group  sessType = 1
)

type sessMode int

const (
	Sleep   sessMode = 0
	Echo    sessMode = 1
	Trigger sessMode = 2
)

type session struct {
	Type         sessType
	Mode         sessMode
	Number       int64
	ExpiredTimer *time.Timer
	MessageCh    chan json.Message
}

func (sess *session) serve() {

}
