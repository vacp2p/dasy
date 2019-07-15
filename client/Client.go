// Package client contains the implementation of a `dasy` client.
package client

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/status-im/dasy/protobuf"
	mvds "github.com/status-im/mvds/node"
	mvdsproto "github.com/status-im/mvds/protobuf"
	"github.com/status-im/mvds/state"
	"github.com/status-im/mvds/store"
)

// Chat is the ID for a specific chat.
type Chat state.GroupID

// Peer is the ID for a specific peer.
type Peer state.PeerID

// Client is the actual daisy client.
type Client struct {
	node  mvds.Node
	store store.MessageStore // @todo we probably need a different message store, not sure tho

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

// We may not need this as we can rely on the acks of data sync
// Ack acknowledges `Join`, `Leave` and `Kick` messages.
func (c *Client) Ack(chat Chat, messageID state.MessageID) {

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

	// @todo sign

	buf, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	id, err := c.node.AppendMessage(state.GroupID(chat), buf)
	if err != nil {
		return err
	}

	c.lastMessage = id

	return nil
}

func (c *Client) onReceive(message mvdsproto.Message) {
	// @todo other shit

	// this is psuedo code
	var msg protobuf.Message
	err := proto.Unmarshal(message.Body, &msg)
	if err != nil {
		// @todo
	}

	if len(msg.PreviousMessage) == 0 {
		return
	}

	id := toMessageID(msg.PreviousMessage)
	if c.store.Has(id) {
		return
	}

	err = c.node.RequestMessage(toGroupID(message.GroupId), id)
	if err != nil {
		log.Printf("error while requesting message: %s", err.Error())
	}
}

// @todo make helper function in MVDS
func toMessageID(b []byte) state.MessageID {
	var id state.MessageID
	copy(id[:], b)
	return id
}

func toGroupID(b []byte) state.GroupID {
	var id state.GroupID
	copy(id[:], b)
	return id
}
