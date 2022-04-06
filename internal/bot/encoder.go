package bot

import (
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/log"
)

func Encoder() {
	go func() {
		for {
			wsReq := <-SendCh
			data, err := json.Encode(wsReq)
			if err != nil {
				log.Log("encoder: encode fail %s", err)
				continue
			}
			go func() {
				BytesSendCh <- data
			}()
		}
	}()
}
