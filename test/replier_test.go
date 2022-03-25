package test_test

import (
	"log"
	"pixie/internal/bot"
	"testing"
	"time"
)

func TestStartReply(t *testing.T) {
	log.Println("Starting reply test")

	bot.Init()
	bot.StartReply()
	var data = [5]string{
		`
			{
				"syncId":1,
				"command":"sendGroupMessage",
				"content":{
					"target":317109237,
					"messageChain":[
						{"type":"Plain","text":"hello1!"}
					]
				}
			}
			`,
		`
			{
				"syncId":2,
				"command":"sendGroupMessage",
				"content":{
					"target":317109237,
					"messageChain":[
						{"type":"Plain","text":"hello2!"}
					]
				}
			}
			`,
		`
			{
				"syncId":3,
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
				"syncId":4,
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
				"syncId":5,
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

	for i := 0; i < 1; i++ {
		bot.MessageSendCh <- []byte(data[i])
	}

	for {
		time.Sleep(1000)
	}
}
