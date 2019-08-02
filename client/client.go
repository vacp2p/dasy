// Package client contains the implementation of a `dasy` client.
package client

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/vacp2p/dasy/protobuf"
	mvds "github.com/vacp2p/mvds/node"
	mvdsproto "github.com/vacp2p/mvds/protobuf"
	"github.com/vacp2p/mvds/state"
	"github.com/vacp2p/mvds/store"
)

// Chat is the ID for a specific chat.
type Chat state.GroupID

// Peer is the ID for a specific peer.
type Peer state.PeerID

// Client is the actual daisy client.
type Client struct {
	node  mvds.Node
	store store.MessageStore // @todo we probably need a different message store, not sure tho

	identity *ecdsa.PrivateKey

	lastMessage state.MessageID // @todo maybe make type
}

// Invite invites a peer to a chat.
func (c *Client) Invite(chat Chat, peer Peer) {

}

// Join joins a chat.
func (c *Client) Join(chat Chat) {

}

// Leave leaves a chat.
func (c *Client) Leave(chat Chat) {

}

// Kick kicks peer from a chat.
func (c *Client) Kick(chat Chat, peer Peer) {

}

// Ack acknowledges `Join`, `Leave` and `Kick` messages.
func (c *Client) Ack(chat Chat, messageID state.MessageID) {
	// @todo: we may not need this as we can rely on the acks of data sync
}

// Post sends a message to a chat.
func (c *Client) Post(chat Chat, body []byte) error {
	return c.send(chat, protobuf.Message_POST, body)
}

func (c *Client) send(chat Chat, t protobuf.Message_MessageType, body []byte) error {
	msg := &protobuf.Message{
		MessageType:     protobuf.Message_MessageType(t),
		Body:            body,
		PreviousMessage: c.lastMessage[:],
	}

	err := c.sign(msg)
	if err != nil {
		return errors.Wrap(err, "failed to sign message")
	}

	buf, err := proto.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "failed to marshall message")
	}

	id, err := c.node.AppendMessage(state.GroupID(chat), buf)
	if err != nil {
		return errors.Wrap(err, "failed to append message")
	}

	c.lastMessage = id

	return nil
}

func (c *Client) onReceive(message mvdsproto.Message) {
	var msg protobuf.Message
	err := proto.Unmarshal(message.Body, &msg)
	if err != nil {
		log.Printf("error while unmarshalling message: %s", err.Error())
		return
	}

	pubkey, err := crypto.SigToPub(msg.ID(), msg.Signature)
	if err != nil {
		log.Printf("error while recovering pubkey: %s", err.Error())
		// @todo
		return
	}

	// @todo probably store the sender somewhere?

	// @todo pump messages to subscriber channels

	if len(msg.PreviousMessage) == 0 {
		return
	}

	c.handlePreviousMessage(
		bytesToGroupID(message.GroupId),
		bytesToMessageID(msg.PreviousMessage),
	)
}

func (c *Client) handlePreviousMessage(group state.GroupID, previousMessage state.MessageID) {
	if c.store.Has(previousMessage) {
		return
	}

	err := c.node.RequestMessage(group, previousMessage)
	if err != nil {
		log.Printf("error while requesting message: %s", err.Error())
	}
}

// sign signs generates a signature of the message and adds it to the message.
func (c *Client) sign(m *protobuf.Message) error {
	hash := m.ID()

	sig, err := crypto.Sign(hash[:], c.identity)
	if err != nil {
		return err
	}

	m.Signature = sig
	return nil
}
