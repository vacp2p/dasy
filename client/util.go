package client

import "github.com/status-im/mvds/state"

func bytesToMessageID(b []byte) state.MessageID {
	var id state.MessageID
	copy(id[:], b)
	return id
}

func bytesToGroupID(b []byte) state.GroupID {
	var id state.GroupID
	copy(id[:], b)
	return id
}
