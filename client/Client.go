// Package client contains the implementation of a `dasy` client.
package client

import (
	"github.com/golang/protobuf/proto"
	"github.com/status-im/dasy/protobuf"
	mvds "github.com/status-im/mvds/node"
	"github.com/status-im/mvds/state"
)

// Chat is the ID for a specific chat.
type Chat state.GroupID

// Peer is the ID for a specific peer.
type Peer state.PeerID

// Client is the actual daisy client.
type Client struct {
	node mvds.Node

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
