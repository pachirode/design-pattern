package pub_sub

type Publisher struct {
	subscribers map[string]chan string
}

func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[string]chan string),
	}
}

func (p *Publisher) Subscribe(name string, ch chan string) {
	p.subscribers[name] = ch
}

func (p *Publisher) Unsubscribe(name string) {
	delete(p.subscribers, name)
}

func (p *Publisher) Publish(topic, message string) {
	if ch, ok := p.subscribers[topic]; ok {
		ch <- message
	}
}
