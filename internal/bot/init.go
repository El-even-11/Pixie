package bot

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"

	"pixie/internal/pkg/net"
)

func Init() {
	parasInit()
	socketInit()
	databaseInit()
}

func parasInit() {
	MessageRecvCh = make(chan []byte, CH_MAX_BUFSIZE)
	EventRecvCh = make(chan []byte, CH_MAX_BUFSIZE)
	MessageSendCh = make(chan Message, CH_MAX_BUFSIZE)
	EventSendCh = make(chan Event, CH_MAX_BUFSIZE)
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
}

func databaseInit() {

}
