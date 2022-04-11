package bot

import (
	"math/rand"
	"pixie/internal/pkg/json"
	"pixie/internal/pkg/redis"
	"strings"
)

func (sess *session) messageHandler(message json.Message) {
	if message.Type == "GroupMessage" {
		go sess.senderHandler(message.Sender)
	}

	for _, messageItem := range message.MessageChain {
		if messageItem.Type == "Plain" && strings.HasPrefix(messageItem.Text, "/") {
			// searching for command
			sess.commandHandler(messageItem.Text, message)
			return
		}
	}

	// no command
	switch sess.mode {
	case Echo:
		sess.echo(message)
	case Trigger:
		sess.trigger(message)
	}
}

func (sess *session) eventHandler(event json.Event) {

}

func (sess *session) senderHandler(sender json.Sender) {

}

func (sess *session) echo(message json.Message) {
	SendCh <- json.WsReq{
		SyncId:  "0",
		Command: "send" + sessionTypeMap[int(sess.sesstype)] + "Message",
		Content: json.Message{
			MessageChain: message.MessageChain,
			Target:       sess.number,
		},
	}
}

const MAX_SPLIT_LENGTH = 10

func (sess *session) trigger(message json.Message) {

	var text string
	for _, messageItem := range message.MessageChain {
		if messageItem.Type == "Plain" {
			text = messageItem.Text
			break
		}
	}

	runetext := []rune(text)
	set := make(map[string]struct{})

	set[text] = struct{}{}

	for i := 0; i < len(runetext); i++ {
		for l := 2; l <= MAX_SPLIT_LENGTH && i+l <= len(runetext); l++ {
			set[string(runetext[i:i+l])] = struct{}{}
		}
	}

	for k := range set {
		// -t shows that the key is a text trigger
		v, err := redis.SMembers(k + "-t")
		if err == nil && len(v) > 0 {
			r := rand.Intn(len(v))
			SendCh <- json.BuildWsReq(
				sess.number,
				"send"+sessionTypeMap[int(sess.sesstype)]+"Message",
				[]string{"Plain"},
				[]string{v[r]},
			)
		}
		// -i shows that the key is a image trigger
		v, err = redis.SMembers(k + "-i")
		if err == nil && len(v) > 0 {
			r := rand.Intn(len(v))
			SendCh <- json.BuildWsReq(
				sess.number,
				"send"+sessionTypeMap[int(sess.sesstype)]+"Message",
				[]string{"Image"},
				[]string{v[r]},
			)
		}
	}
}
