package broker

import (
	"fmt"
	"net/rpc"
)

type AsyncSender struct {
	hook string
}

func NewAsyncSender(hook string) *AsyncSender {
	return &AsyncSender{hook: hook}
}

func (sender *AsyncSender) Ack(header string) {
	client, err := rpc.Dial("tcp", sender.hook)

	if err != nil {
		fmt.Println(err)
		return
	}

	var relpy string
	err = client.Call("Client.Ack", header, &relpy)

	if err != nil {
		fmt.Println(err)
	}
}
