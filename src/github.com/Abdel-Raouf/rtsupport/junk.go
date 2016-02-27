package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func main() {
	recRawMsg := []byte(`{"name":"channel add",` + `"data":{"name":"Hardware Support"}}`)
	var recMessage Message
	err := json.Unmarshal(recRawMsg, &recMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", recMessage)

	if recMessage.Name == "channel add" {
		channel, err := addChannel(recMessage.Data)
		var sendMessage Message
		sendMessage.Name = "channel add"
		sendMessage.Data = channel
		sendRawMsg, err := json.Marshal(sendMessage)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(string(sendRawMsg))
	}
}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return channel, err
	}
	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	return channel, nil
}









