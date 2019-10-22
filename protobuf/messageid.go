package protobuf

import (
	"crypto/sha256"
	"encoding/binary"
)

func (m *Message) ID() []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(m.MessageType))
	b = append(b, m.Body...)

	hash := sha256.Sum256(b)
	return hash[:]
}
