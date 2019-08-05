package event

type Subscription chan <-Payload

type Feed struct {
	subscribers []Subscription
}

func (f *Feed) Subscribe(channel Subscription) { // @todo think about returning a subscription like prysm
	f.subscribers = append(f.subscribers, channel)
}

// Send sends a payload to all the subscribers for the specific feed.
func (f *Feed) Send(value Payload) {
	// @todo is this good enough for now?
	for _, sub := range f.subscribers {
		sub <- value
	}
}
