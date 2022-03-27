package bot

func Serve() {
	for {
		select {
		case m := <-MessageRecvCh:
			respMessage := serveMessage(m)
			if respMessage.Empty {
				continue
			}
			go func() {
				MessageSendCh <- respMessage
			}()
			<-respMessage.Done
		case e := <-EventRecvCh:
			respEvent := serveEvent(e)
			if respEvent.Empty {
				continue
			}
			go func() {
				EventSendCh <- respEvent
			}()
			<-respEvent.Done
		}
	}
}

func serveMessage(data []byte) Message {
	return Message{}
}

func serveEvent(data []byte) Event {
	return Event{}
}

func decodeEvent() {

}
