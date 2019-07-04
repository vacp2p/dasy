// Package client contains the implementation of a `dasy` client.
package client

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/status-im/dasy/protobuf"
	mvds "github.com/status-im/mvds/node"
	"github.com/status-im/mvds/state"
)

// Chat is the ID for a specific chat.
type Chat state.GroupID

// Client is the actual daisy client.
type Client struct {
	node mvds.Node
}

// Invite invites a peer to a chat.
func (c *Client) Invite(chat Chat, peer []byte) {

}

// Join joins a chat.
func (c *Client) Join(chat Chat) {

}

// Leave leaves a chat.
func (c *Client) Leave(chat Chat) {

}

// Kick kicks peer from a chat.
func (c *Client) Kick(chat Chat, peer []byte) {

}

// Ack acknowledges `Join`, `Leave` and `Kick` messages.
func (c *Client) Ack(chat Chat, messageID []byte) {

}

// Post sends a message to a chat.
func (c *Client) Post(chat Chat, body []byte) {
	c.send(chat, protobuf.Message_POST, body) // @todo
}

func (c *Client) send(chat Chat, t protobuf.Message_MessageType, body []byte) error {
	msg := &protobuf.Message{
		Clock: 0,
		Timestamp: uint64(time.Now().Unix()), // @todo we may be able to take this from mvds
		MessageType: protobuf.Message_MessageType(t),
		Body: body,
	}

	// @todo sign

	buf, err := proto.Marshal(msg)
	if err != nil {
		return  err
	}

	_, err = c.node.AppendMessage(state.GroupID(chat), buf)
	if err != nil {
		return err
	}

	return nil
}
