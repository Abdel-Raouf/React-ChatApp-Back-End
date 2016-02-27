package main

import(
	"fmt"
	"time"
	"math/rand"
)

type Message struct {
	Name string 		`json:"name"` 
	Data interface{}	`json:"data"`
}

type Client struct{
	send chan Message
}

func (client *Client) write(){
	for msg := range client.send {
		//TODO: socket.sendJSON(msg)
		fmt.Printf("%#v\n", msg)
	}
}

func (client *Client) subscribeChannels(){
	//TODO: changefeed Query RethinkDB
	for {
		time.Sleep(r())
		client.send <- Message{"channel add", ""}
	}
}

func (client *Client) subscribeMessages(){
	//TODO: changefeed Query RethinkDB
	for {
		time.Sleep(r())
		client.send <- Message{"message add", ""}
	}
}


func r() time.Duration {
	// generates random values between 0 and 1000 * milli-sec to give us seconde or less than fpr the DB simulation in time.Sleep().
	return time.Millisecond * time.Duration(rand.Intn(1000))
}

func NewClient() *Client{
	//return the pointer (*Client) for the new instantiated client
	return &Client{
		// one of the little gotchas u will need comma even if it is the last field.
		send: make(chan Message),
	}
}

func main() {
	//to instantiate new client
	client := NewClient()
	go client.subscribeChannels()
	go client.subscribeMessages()
	client.write()
}












