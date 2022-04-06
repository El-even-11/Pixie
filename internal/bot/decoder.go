package bot

import (
	"pixie/internal/pkg/debug"
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/log"
)

func Decoder() {
	go func() {
		for {
			select {
			case m := <-MessageBytesRecvCh:
				messageChainItf, err := json.Decode(m, true)
				if err != nil {
					log.Log("%s", err)
					break
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
					log.Log("%s", err)
					break
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
