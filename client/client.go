package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"time"

	"github.com/moxeed/groker/broker"
)

type Client int

func (client *Client) Recieve(message broker.BaseMessage, reply *string) error {
	fmt.Println("message recived " + message.Body)

	return nil
}

func recieve(endpoint string) {
	addy, err := net.ResolveTCPAddr("tcp", endpoint)

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
	rand.Seed(time.Now().Unix())
	endpoint := fmt.Sprintf("localhost:%d", rand.Intn(1000)+9000)
	client, err := rpc.Dial("tcp", "localhost:5000")

	if err != nil {
		log.Fatal(err)
	}

	var relpy string
	err = client.Call("Service.Subscribe", endpoint, &relpy)

	if err != nil {
		log.Fatal(err)
	}

	recieve(endpoint)
}
