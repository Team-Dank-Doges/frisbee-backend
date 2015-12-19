package main

type Message struct {
	Sender User `json:"sender"`
	// Thinking `receiver` might not be necessary with websockets.
	//Receiver []User `json:"receiver"`
	Body string `json:"body"`
}

func (msg *Message) String() string {
	return msg.Sender.String() + " says " + msg.Body
}
