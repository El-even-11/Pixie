package test_test

import (
	"log"
	"pixie/internal/bot"
	"pixie/internal/pkg/net"
	"testing"
)

func TestInit(t *testing.T) {
	bot.Init()
	_, message, err := net.MessageConn.ReadMessage()
	if err != nil {
		log.Fatalf("read: %s", err)
	}
	log.Printf("recv: %s", message)

	_, message, err = net.MessageConn.ReadMessage()
	if err != nil {
		log.Fatalf("read: %s", err)
	}
	log.Printf("recv: %s", message)

	_, message, err = net.EventConn.ReadMessage()
	if err != nil {
		log.Fatalf("read: %s", err)
	}
	log.Printf("recv: %s", message)

	_, message, err = net.EventConn.ReadMessage()
	if err != nil {
		log.Fatalf("read: %s", err)
	}
	log.Printf("recv: %s", message)
}