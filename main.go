package main

import (
	"log"
	"net/http"

	"github.com/golang-samples/websocket/websocket-chat/src/chat"
)

func main() {
	log.SetFlags(log.Lshortfile)

	server := chat.NewServer("/entry")
	go server.Listen()

	http.Handle("/", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServe(":7123", nil))
}
