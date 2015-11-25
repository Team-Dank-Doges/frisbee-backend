package main

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	ExampleHandler()
}

func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func ExampleHandler() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
