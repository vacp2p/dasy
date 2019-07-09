// Package protobuf contains protocol buffers as well as helper functions for those buffers.
package protobuf

import (
	"crypto/sha256"
	"encoding/binary"
)

// MessageID is a hash of `message_type`, `body` & `previous_message`.
type MessageID [32]byte

// ID returns the message ID.
func (m *Message) ID() MessageID {
	t := make([]byte, 8)
	binary.LittleEndian.PutUint32(t, uint32(m.MessageType))

	b := append([]byte("MESSAGE_ID"), t...)
	b = append(b, m.Body...)
	b = append(b, m.PreviousMessage...)

	return sha256.Sum256(b)
}
