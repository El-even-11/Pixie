package bot

var VERIFY_KEY = "1234567890"
var QQ_NUMBER = "2473537565"
var HOST = "localhost:8080"

var CH_MAX_BUFSIZE = 10

var MessageRecvCh chan []byte
var EventRecvCh chan []byte

var MessageSendCh chan Message
var EventSendCh chan Event

type Message struct {
	Data []byte
	Done chan struct{}
}

type Event struct {
	Data []byte
	Done chan struct{}
}
