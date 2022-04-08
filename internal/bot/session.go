package bot

import (
	"pixie/internal/pkg/json"
	"time"
)

type sessionType int

const (
	Friend sessionType = 0
	Group  sessionType = 1
)

type sessionMode int

const (
	Sleep   sessionMode = 0
	Echo    sessionMode = 1
	Trigger sessionMode = 2
)

type session struct {
	sesstype     sessionType
	mode         sessionMode
	number       int64
	expiredTimer *time.Timer
	messageCh    chan json.Message
}

const SESSION_TIME_OUT = time.Second * 120

func (sess *session) serve() {
	for {
		
	}
}
