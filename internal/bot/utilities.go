package bot

import (
	"pixie/internal/pkg/json"
)

const VERIFY_KEY = "1234567890"
const QQ_NUMBER = "2473537565"
const HOST = "localhost:8080"

const CH_MAX_BUFSIZE = 10

var BytesSendCh chan []byte

var MessageRecvCh chan json.Message
var EventRecvCh chan json.Event

var SendCh chan json.WsReq

var LogCh chan string
