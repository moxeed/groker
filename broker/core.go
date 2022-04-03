package broker

import (
	"fmt"
	"sync"
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
	client      sync.Mutex
}

var instance broker

func Send(message Message) error {
	fmt.Println("***Message Recieved")

	select {
	case instance.queue <- message:
	default:
		println("overflow")
		return fmt.Errorf("buffer is full")
	}

	return nil
}

func Subscribe(subscriber Subscriber) {
	defer instance.client.Unlock()

	fmt.Println("***Subscriber Registered " + subscriber.hook)
	instance.subscribers = append(instance.subscribers, subscriber)
}

func init() {
	instance = broker{
		make(chan Message, 10),
		[]Subscriber{},
		sync.Mutex{},
	}

	instance.start(1)
}

func publish(message Message) {
	if len(instance.subscribers) == 0 {
		instance.client.Lock()
	}

	for index, subscriber := range instance.subscribers {
		err := subscriber.Push(message)
		if err != nil {
			instance.subscribers = append(instance.subscribers[:index], instance.subscribers[index+1:]...)
		}
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
