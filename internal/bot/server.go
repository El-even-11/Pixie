package bot

import (
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/log"
)

func Serve() {
	sessMap := make(map[int64]*session)

	for {
		select {
		case m := <-MessageRecvCh:
			switch m.Type {
			case "GroupMessage":
				num := m.Sender.Group.ID
				if _, ok := sessMap[num]; !ok {
					sessMap[num] = &session{
						sesstype:       Group,
						mode:           Sleep,
						number:         num,
						messageCh:      make(chan json.Message, 10),
					}
					go sessMap[num].serve()
				}
				go func() {
					sessMap[num].messageCh <- m
				}()

			case "FriendMessage":
				num := m.Sender.ID
				if _, ok := sessMap[num]; !ok {
					sessMap[num] = &session{
						sesstype:  Friend,
						mode:      Sleep,
						number:    num,
						messageCh: make(chan json.Message, 10),
					}
					go sessMap[num].serve()
				}
				go func() {
					sessMap[num].messageCh <- m
				}()

			default:
				log.Log("unknown message type")
			}

		case e := <-EventRecvCh:
			switch e.Type {
			default:
				log.Log("unknown event type")
			}
		}
	}
}
