package broker

import (
	"fmt"
	"net/rpc"
)

type Subscriber struct {
	hook string
}

func (subscriber Subscriber) Push(message Message) {
	client, err := rpc.Dial("tcp", subscriber.hook)

	if err != nil {
		fmt.Println(err)
	}

	var relpy string
	baseMessage := BaseMessage{Header: message.Header, Body: message.Body}
	err = client.Call("Client.Recieve", baseMessage, &relpy)

	if err != nil {
		fmt.Println(err)
	}
}
