package test_test

import (
	"pixie/internal/bot"
	"testing"
)

func TestStartListen(t *testing.T) {
	bot.Init()
	bot.StartListen()
	for i := 0; i < 5; i++ {
		<-bot.MessageRecvCh
	}
}
