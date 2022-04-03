package broker

type Service int

type BaseMessage struct {
	Header string
	Body   string
}

type AsyncMessage struct {
	BaseMessage
	Hook string
}

func (service *Service) Subscribe(endpoint string, response *string) error {
	subscriber := Subscriber{endpoint}
	Subscribe(subscriber)

	*response = "ok"

	return nil
}

func (service *Service) SyncPublish(data BaseMessage, response *string) error {
	sender := NewSyncSender()
	message := Message{
		Body:   data.Body,
		Header: data.Header,
		Sender: sender,
	}

	err := Send(message)

	if err != nil {
		return err
	}

	sender.Wait()

	*response = "ok"

	return nil
}

func (service *Service) AsyncPublish(data AsyncMessage, response *string) error {
	sender := NewAsyncSender(data.Hook)
	message := Message{
		Body:   data.Body,
		Header: data.Header,
		Sender: sender,
	}

	err := Send(message)

	if err != nil {
		return err
	}

	*response = "ok"

	return nil
}
