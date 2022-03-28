package bot

import (
	"log"
	"pixie/internal/pkg/debug"
	"pixie/internal/pkg/json"
)

func Decoder() {
	go func() {
		for {
			select {
			case m := <-MessageBytesRecvCh:
				messageChainItf, err := json.Decode(m, true)
				if err != nil {
					log.Printf("%s", err)
					return
				}
				debug.DPrintf("decode: %v", messageChainItf)

				messageChain, ok := messageChainItf.(json.MessageChain)
				if !ok {
					panic("message chain type error!")
				}
				go func() {
					MessageRecvCh <- messageChain
				}()

			case e := <-EventBytesRecvCh:
				eventItf, err := json.Decode(e, false)
				if err != nil {
					log.Printf("%s", err)
					return
				}

				debug.DPrintf("decode: %v", eventItf)

				event, ok := eventItf.(json.Event)
				if !ok {
					panic("message chain type error!")
				}

				go func() {
					EventRecvCh <- event
				}()
			}
		}
	}()
}
