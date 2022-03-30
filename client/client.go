package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"

	"github.com/moxeed/groker/broker"
)

type Client int

func (client *Client) Ack(data string, reply *string) error {
	fmt.Println("message acknoledged " + data)

	return nil
}

func (client *Client) Recieve(message broker.BaseMessage, reply *string) error {
	fmt.Println("message recived " + message.Body)

	return nil
}

var wg sync.WaitGroup

func recieve(endpoint string) {

	wg.Add(1)

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

	wg.Done()
}

func main() {
	rand.Seed(time.Now().Unix())
	endpoint := fmt.Sprintf("0.0.0.0:%d", rand.Intn(10000))
	client, err := rpc.Dial("tcp", "0.0.0.0:5000")

	if err != nil {
		log.Fatal(err)
	}

	go recieve(endpoint)

	var relpy string
	message := broker.AsyncMessage{
		BaseMessage: broker.BaseMessage{
			Header: "h1",
			Body:   "h2",
		},
		Hook: endpoint,
	}

	err = client.Call("Service.Subscribe", endpoint, &relpy)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Call("Service.AsyncPublish", message, &relpy)

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
