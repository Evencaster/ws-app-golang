package hub

import "sync"

type Subscriber interface {
	Update()
}

type Observer struct {
	m sync.RWMutex
	subscribers map[string][]Subscriber
}

func NewObserver() *Observer {
	var obs Observer
	obs.subscribers = make(map[string][]Subscriber)
	return &obs
}

func (obs *Observer) Subscribe(topic string, sub Subscriber) error {
	obs.m.Lock()
	defer obs.m.Unlock()

	obs.subscribers[topic] = append(obs.subscribers[topic], sub)

	return nil
}

func (obs *Observer) NotifyAll(topic string) error {
	obs.m.RLock()
	topicSubs := obs.subscribers[topic]
	obs.m.RUnlock()

	for _, sub := range topicSubs {
		sub.Update()
	}
	
	return nil
}
