package main

import "golang.org/x/net/websocket"

type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Socket *websocket.Conn
}

func (u *User) String() string {
	return u.Name + " <" + u.Email + ">"
}
