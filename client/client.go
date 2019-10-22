// Package client contains the implementation of a `dasy` client.
package client

// @todo add ordering by default

import (
	"crypto/ecdsa"
	"log"
	"sync"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/vacp2p/dasy/client/internal"
	"github.com/vacp2p/dasy/crypto"
	"github.com/vacp2p/dasy/event"
	"github.com/vacp2p/dasy/protobuf"
	mvdsproto "github.com/vacp2p/mvds/protobuf"
	"github.com/vacp2p/mvds/state"
	"github.com/vacp2p/mvds/store"
)

// Chat is the ID for a specific chat.
type Chat state.GroupID

// Peer is the ID for a specific peer.
type Peer state.PeerID

// Client is the actual dasy client.
type Client struct {
	sync.Mutex

	id Peer // @todo think of turning dataSyncNode ID into a func

	node  internal.DataSyncNode
	store store.MessageStore

	identity *ecdsa.PrivateKey

	feeds        map[protobuf.Message_MessageType]*event.Feed
}

// Invite invites a peer to a chat.
func (c *Client) Invite(chat Chat, peer Peer) {

}

// Join joins a chat.
func (c *Client) Join(chat Chat) (state.MessageID, error) {
	return c.send(chat, protobuf.Message_JOIN, c.id[:])
}

// Leave leaves a chat.
func (c *Client) Leave(chat Chat) (state.MessageID, error) {
	return c.send(chat, protobuf.Message_LEAVE, c.id[:])
}

// Kick kicks peer from a chat.
func (c *Client) Kick(chat Chat, peer Peer) (state.MessageID, error) {
	return c.send(chat, protobuf.Message_KICK, peer[:])
}

// Ack acknowledges `Join`, `Leave` and `Kick` messages.
func (c *Client) Ack(chat Chat, messageID state.MessageID) (state.MessageID, error) {
	// @todo We may not need this as we can rely on the acks of data sync
	return c.send(chat, protobuf.Message_ACK, messageID[:])
}

// Post sends a message to a chat.
func (c *Client) Post(chat Chat, body []byte) (state.MessageID, error) {
	return c.send(chat, protobuf.Message_POST, body)
}

// Feed is a subscription feed for the specified message type.
func (c *Client) Feed(msg protobuf.Message_MessageType) *event.Feed {
	c.Lock()
	defer c.Unlock()

	if c.feeds[msg] == nil {
		c.feeds[msg] = new(event.Feed)
	}

	return c.feeds[msg]
}

// Listen listens for newly received messages and handles them appropriately.
func (c *Client) Listen() {
	sub := c.node.Subscribe()

	for {
		go c.onReceive(<-sub)
	}
}

func (c *Client) send(chat Chat, t protobuf.Message_MessageType, body []byte) (state.MessageID, error) {
	c.Lock()
	defer c.Unlock()

	msg := &protobuf.Message{
		MessageType:     protobuf.Message_MessageType(t),
		Body:            body,
	}

	err := crypto.Sign(c.identity, msg)
	if err != nil {
		return state.MessageID{}, errors.Wrap(err, "failed to sign message")
	}

	buf, err := proto.Marshal(msg)
	if err != nil {
		return state.MessageID{}, errors.Wrap(err, "failed to marshall message")
	}

	id, err := c.node.AppendMessage(state.GroupID(chat), buf)
	if err != nil {
		return state.MessageID{}, errors.Wrap(err, "failed to append message")
	}

	return id, nil
}

// onReceive handles lower level message receiving logic, such as requesting all previous message dependencies that we
// may not have, as well as unmarshalling and storing the message.
func (c *Client) onReceive(message mvdsproto.Message) {
	var msg protobuf.Message
	err := proto.Unmarshal(message.Body, &msg)
	if err != nil {
		log.Printf("error while unmarshalling message: %s", err.Error())
		return
	}

	pubkey, err := ethcrypto.SigToPub(msg.ID(), msg.Signature)
	if err != nil {
		log.Printf("error while recovering pubkey: %s", err.Error())
		// @todo
		return
	}

	payload := event.Payload{
		Body:      msg.Body,      // @todo this might need to be unmarshalled depending on the message type like invite?
		Signature: msg.Signature, // @todo recover from signature
		Sender:    crypto.PublicKeyToPeerID(*pubkey),
		Timestamp: message.Timestamp,
	}

	go c.Feed(msg.MessageType).Send(payload)
}
