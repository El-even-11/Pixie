package bot

import (
	"pixie/internal/pkg/log"
)

func Logger() {
	go func() {
		for {
			select {
			case s := <-LogCh:
				log.Log(s)
			}
		}
	}()
}
