package bot

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func Init() {
	socketInit()
}

func socketInit() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/message",
	}

	header := http.Header{
		"verifyKey": {"1234567890"},
		"qq":        {"2473537565"},
	}

	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Fatal("dial:", err)
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
		}
		log.Printf("recv: %s", message)
	}
}

func databaseInit() {

}
