package bot

func Serve() {
	sessMap := make(map[int64]session)

	for {
		select {
		case m := <-MessageRecvCh:
			go messageHandler(m)
		case e := <-EventRecvCh:
			go eventHandler(e)
		}
	}
}
