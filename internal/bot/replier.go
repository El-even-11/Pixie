package bot

func Reply() {
	go func() {
		for {
			select {
			case data := <-messageSendCh:
				
			case data := <-eventSendCh:

			}
		}
	}()
}
