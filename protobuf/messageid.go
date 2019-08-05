package protobuf

import (
	"crypto/sha256"
	"encoding/binary"
)

func (m Message) ID() [32]byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(m.MessageType))
	b = append(b, m.Body...)
	b = append(b, m.PreviousMessage...)

	return sha256.Sum256(b)
}
