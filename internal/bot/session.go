package bot

import "time"

type sessType int

const (
	Friend sessType = 0
	Group  sessType = 0
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
}

func (sess *session) serve() {

}
