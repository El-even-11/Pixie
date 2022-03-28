package main

import (
	"pixie/internal/bot"
)

func main() {
	bot.Init()
	bot.Listener()
	bot.Writer()
	bot.Serve()
}
