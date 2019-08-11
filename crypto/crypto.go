package crypto

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vacp2p/dasy/protobuf"
	"github.com/vacp2p/mvds/state"
)

// PublicKeyToPeerID turns an ECDSA Public Key to an mvds PeerID.
func PublicKeyToPeerID(k ecdsa.PublicKey) state.PeerID {
	var p state.PeerID
	copy(p[:], crypto.FromECDSAPub(&k))
	return p
}

// Sign signs generates a signature of the message and adds it to the message.
func Sign(identity *ecdsa.PrivateKey, m *protobuf.Message) error {
	hash := m.ID()

	sig, err := crypto.Sign(hash[:], identity)
	if err != nil {
		return err
	}

	m.Signature = sig
	return nil
}
