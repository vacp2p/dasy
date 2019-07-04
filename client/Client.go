// Package client contains the implementation of a `dasy` client.
package client

import (
	"github.com/golang/protobuf/proto"
	mvds "github.com/status-im/mvds/node"
	"github.com/status-im/mvds/protobuf"
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

}

func (c *Client) send(chat Chat, message *protobuf.Message) error {
	buf, err := proto.Marshal(message)
	if err != nil {
		return  err
	}

	_, err = c.node.AppendMessage(state.GroupID(chat), buf)
	if err != nil {
		return err
	}

	return nil
}
