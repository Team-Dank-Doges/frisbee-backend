package main

type Message struct {
	Sender   User   `json:"sender"`
	Receiver []User `json:"receiver"`
	Body     string `json:"body"`
}
