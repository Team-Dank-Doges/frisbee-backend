package main

import "golang.org/x/net/websocket"

type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Socket *websocket.Conn
}
