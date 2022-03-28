package bot

import (
	"log"
	"pixie/internal/pkg/json"
)

func Encoder() {
	go func() {
		for {
			wsReq := <-SendCh
			data, err := json.Encode(wsReq)
			if err != nil {
				log.Printf("encoder: encode fail %s", err)
				continue
			}
			go func() {
				BytesSendCh <- data
			}()
		}
	}()
}
