package broker

import (
	"fmt"
	"sync"
)

type syncSender struct {
	wg sync.WaitGroup
}

func NewSyncSender() *syncSender {
	sender := syncSender{}
	sender.wg.Add(1)
	return &sender
}

func (sender *syncSender) Ack(header string) {
	fmt.Println("Ack")
	sender.wg.Done()
}

func (sender *syncSender) Wait() {
	sender.wg.Wait()
}
