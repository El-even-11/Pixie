package main

import (
	"pixie/internal/bot"
)

func main() {
	bot.Init()
	bot.Listen()
	bot.Reply()
	
}
