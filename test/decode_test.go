package test_test

import (
	"log"
	"pixie/internal/pkg/json"
	"testing"
)

func TestDecode(t *testing.T) {
	data := `{"syncId":"-1","data":{"type":"GroupMessage","messageChain":[{"type":"Source","id":1207912,"time":1648286819}],"sender":{"id":2733984991,"memberName":"眺望狗","specialTitle":"Three","permission":"ADMINISTRATOR","joinTimestamp":1468756346,"lastSpeakTimestamp":1648286819,"muteTimeRemaining":0,"group":{"id":317109237,"name":"作业恋爱联盟","permission":"ADMINISTRATOR"}}}}`
	messageChain, err := json.Decode([]byte(data), true)
	if err != nil {
		panic("fail")
	}
	log.Printf("%s", messageChain)
}
