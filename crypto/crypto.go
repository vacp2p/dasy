package crypto

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vacp2p/mvds/state"
)

func PublicKeyToPeerID(k ecdsa.PublicKey) state.PeerID {
	var p state.PeerID
	copy(p[:], crypto.FromECDSAPub(&k))
	return p
}
