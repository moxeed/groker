package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/moxeed/groker/broker"
)

func main() {

	addy, err := net.ResolveTCPAddr("tcp", "localhost:5000")

	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	service := new(broker.Service)
	rpc.Register(service)
	rpc.Accept(inbound)
}
