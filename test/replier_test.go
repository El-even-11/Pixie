package test_test

import (
	"log"
	"pixie/internal/bot"
	"testing"
	"time"
)

var data = [5]string{
	`
			{
				"syncId":1,
				"command":"sendGroupMessage",
				"content":{
					"target":317109237,
					"messageChain":[
						{"type":"Plain","text":"test1131"}
					]
				}
			}
			`,
	`
			{
				"syncId":1,
				"command":"sendGroupMessage",
				"content":{
					"target":317109237,
					"messageChain":[
						{"type":"Plain","text":"hello233!"}
					]
				}
			}
			`,
	`
			{
				"syncId":1,
				"command":"sendGroupMessage",
				"content":{
					"target":317109237,
					"messageChain":[
						{"type":"Plain","text":"hello3!"}
					]
				}
			}
			`,
	`
			{
				"syncId":1,
				"command":"sendGroupMessage",
				"content":{
					"target":317109237,
					"messageChain":[
						{"type":"Plain","text":"hello4!"}
					]
				}
			}
			`,
	`
			{
				"syncId":12,
				"command":"sendGroupMessage",
				"content":{
					"target":317109237,
					"messageChain":[
						{"type":"Plain","text":"hello5!"}
					]
				}
			}
			`,
}

func TestStartReply(t *testing.T) {
	log.Println("Starting reply test")

	bot.Init()
	bot.Writer()

	for i := 0; i < 5; i++ {
		m := bot.Message{
			Data: []byte(data[i]),
			Done: make(chan struct{}),
		}
		bot.MessageSendCh <- m
		<-m.Done
	}
}

func TestListenAndReply(t *testing.T) {
	bot.Init()
	bot.Listener()
	bot.Writer()

	time.Sleep(time.Second * 1)
}
