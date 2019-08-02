package client

type Feed interface {
	Subscribe(channel interface{}) // @todo consider returning events
	Send(value Payload) int
}
