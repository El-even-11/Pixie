package bot

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"

	"pixie/internal/pkg/net"
)

func Init() {
	socketInit()
}

func socketInit() {

	header := http.Header{
		"verifyKey": {"1234567890"},
		"qq":        {"2473537565"},
	}

	um := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
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
		Host:   "localhost:8080",
		Path:   "/event",
	}

	net.EventConn, _, err = websocket.DefaultDialer.Dial(ue.String(), header)
	if err != nil {
		log.Fatal("dial:", err)
	}
}

func databaseInit() {
	
}
