package crypto

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vacp2p/dasy/protobuf"
	"github.com/vacp2p/mvds/state"
)

func PublicKeyToPeerID(k ecdsa.PublicKey) state.PeerID {
	var p state.PeerID
	copy(p[:], crypto.FromECDSAPub(&k))
	return p
}

// Sign signs generates a signature of the message and adds it to the message.
func Sign(identity *ecdsa.PrivateKey, m *protobuf.Message) error {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(m.MessageType))
	b = append(b, m.Body...)
	b = append(b, m.PreviousMessage...)

	hash := sha256.Sum256(b)

	sig, err := crypto.Sign(hash[:], identity)
	if err != nil {
		return err
	}

	m.Signature = sig
	return nil
}
