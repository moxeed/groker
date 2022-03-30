package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/moxeed/groker/broker"
)

type Client int

func (client *Client) Ack(data string, reply *string) error {
	fmt.Println("message acknoledged " + data)

	return nil
}

func recieve() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:5001")

	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	service := new(Client)
	rpc.Register(service)
	rpc.Accept(inbound)
}

func main() {

	go recieve()

	client, err := rpc.Dial("tcp", "0.0.0.0:5000")

	if err != nil {
		log.Fatal(err)
	}

	var relpy string
	message := broker.AsyncMessage{
		BaseMessage: broker.BaseMessage{
			Header: "h1",
			Body:   "h2",
		},
		Hook: "0.0.0.0:5001",
	}

	err = client.Call("Service.AsyncPublish", message, &relpy)

	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1000)
}
