package bot

import (
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"

	"pixie/internal/pkg/json"
	"pixie/internal/pkg/net"
	"pixie/internal/pkg/redis"
)

func Init() {
	rand.Seed(time.Now().Unix())
	parasInit()
	socketInit()
	databaseInit()
}

func parasInit() {
	MessageRecvCh = make(chan json.Message, CH_MAX_BUFSIZE)
	EventRecvCh = make(chan json.Event, CH_MAX_BUFSIZE)
	BytesSendCh = make(chan []byte, CH_MAX_BUFSIZE)
	SendCh = make(chan json.WsReq, CH_MAX_BUFSIZE)
	LogCh = make(chan string, CH_MAX_BUFSIZE)
}

func socketInit() {

	header := http.Header{
		"verifyKey": {VERIFY_KEY},
		"qq":        {QQ_NUMBER},
	}

	um := url.URL{
		Scheme: "ws",
		Host:   HOST,
		Path:   "/message",
	}

	log.Printf("connecting to %s", um.String())

	var err error
	net.MessageConn, _, err = websocket.DefaultDialer.Dial(um.String(), header)
	if err != nil {
		log.Fatal("dial:", err)
	}

	ue := url.URL{
		Scheme: "ws",
		Host:   HOST,
		Path:   "/event",
	}

	net.EventConn, _, err = websocket.DefaultDialer.Dial(ue.String(), header)
	if err != nil {
		log.Fatal("dial:", err)
	}

	ua := url.URL{
		Scheme: "ws",
		Host:   HOST,
		Path:   "/all",
	}

	net.AllConn, _, err = websocket.DefaultDialer.Dial(ua.String(), header)
	if err != nil {
		log.Fatal("dial:", err)
	}
}

func databaseInit() {
	redis.RedisInit()
}
