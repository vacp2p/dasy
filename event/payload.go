package event

import "github.com/vacp2p/mvds/state"

// Payload represents a `dasy` packet.
type Payload struct {
	Body      interface{}
	Signature []byte
	Sender    state.PeerID
	Timestamp int64
}
