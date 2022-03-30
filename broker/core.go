package broker

import (
	"fmt"
)

type Sender interface {
	Ack(string)
}

type Message struct {
	Header string
	Body   string
	Sender Sender
}

type broker struct {
	queue       chan Message
	subscribers []Subscriber
}

var instance broker

func Send(message Message) {
	fmt.Println("***Message Recieved")
	instance.queue <- message
}

func Subscribe(subscriber Subscriber) {
	fmt.Println("***Subscriber Registered " + subscriber.hook)
	instance.subscribers = append(instance.subscribers, subscriber)
}

func init() {
	instance = broker{
		make(chan Message, 10),
		[]Subscriber{},
	}

	instance.start(1)
}

func publish(message Message) {
	for _, subscriber := range instance.subscribers {
		subscriber.Push(message)
	}
	fmt.Println("***Message Published")
}

func queue_agent(queue chan Message) {
	for message := range queue {
		publish(message)
		message.Sender.Ack(message.Header)
	}
}

func (broker *broker) start(agent_count int) {
	for i := 0; i < agent_count; i++ {
		go queue_agent(broker.queue)
	}
}
