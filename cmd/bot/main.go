package main

import (
	"pixie/internal/bot"
)

func main() {
	bot.Init()
	bot.StartListen()
	bot.StartReply()
	bot.Serve()
}
