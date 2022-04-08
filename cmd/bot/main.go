package main

import (
	"pixie/internal/bot"
)

func main() {
	bot.Init()
	bot.Logger()
	bot.Listener()
	bot.Writer()
	bot.Encoder()
	bot.Serve()
}
