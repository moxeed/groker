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

func recieve(port string) {
	addy, err := net.ResolveTCPAddr("tcp", "localhost:"+port)

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

	port := "8080"
	go recieve(port)

	client, err := rpc.Dial("tcp", "localhost:5000")

	if err != nil {
		log.Fatal(err)
	}

	var relpy string
	i := 1

	for {

		message := broker.AsyncMessage{
			BaseMessage: broker.BaseMessage{
				Header: "h1",
				Body:   fmt.Sprintf("async text from server %d", i),
			},
			Hook: "localhost:" + port,
		}

		err = client.Call("Service.AsyncPublish", message, &relpy)

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(1_000_000_000)

		err = client.Call("Service.SyncPublish", message.BaseMessage, &relpy)

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(1_000_000_000)

		i++
	}
}
