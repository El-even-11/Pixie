package bot

var VERIFY_KEY = "1234567890"
var QQ_NUMBER = "2473537565"
var HOST = "localhost:8080"

var CH_MAX_BUFSIZE = 10

var messageRecvCh chan []byte
var eventRecvCh chan []byte

var messageSendCh chan []byte
var eventSendCh chan []byte