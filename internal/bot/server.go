package bot

func Serve() {
	for {
		select {
		case m := <-MessageRecvCh:
			go messageHandler(m)
		case e := <-EventRecvCh:
			go eventHandler(e)
		}
	}
}
