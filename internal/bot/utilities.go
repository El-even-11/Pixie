package bot

const VERIFY_KEY = "1234567890"
const QQ_NUMBER = "2473537565"
const HOST = "localhost:8080"

const CH_MAX_BUFSIZE = 10

var MessageRecvCh chan []byte
var EventRecvCh chan []byte

var MessageSendCh chan Message
var EventSendCh chan Event

type BotStatus int

const (
	Running  BotStatus = 1
	Sleeping BotStatus = 2
	Dead     BotStatus = 3
)

type Message struct {
	Empty bool
	Data  []byte
	Done  chan struct{}
}

type Event struct {
	Empty bool
	Data  []byte
	Done  chan struct{}
}
