package bot

import "pixie/internal/pkg/json"

const VERIFY_KEY = "1234567890"
const QQ_NUMBER = "2473537565"
const HOST = "localhost:8080"

const CH_MAX_BUFSIZE = 10

var MessageBytesRecvCh chan []byte
var EventBytesRecvCh chan []byte

var BytesSendCh chan []byte

var MessageRecvCh chan json.MessageChain
var EventRecvCh chan json.Event

var SendCh chan json.WsReq

type BotStatus int

const (
	Running  BotStatus = 1
	Sleeping BotStatus = 2
	Dead     BotStatus = 3
)
