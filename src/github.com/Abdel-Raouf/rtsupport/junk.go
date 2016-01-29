package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string 
	Data interface{}
}


func main() {
	recRawMsg := []byte('{"name":"channel add",' + '"data":{"name":"Hardware Support"}}')
	var recMessage Message
	err := json.Unmarshal(recRawMsg, &recMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", recMessage)
}