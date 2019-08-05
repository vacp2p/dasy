package internal

import (
	"github.com/vacp2p/mvds/protobuf"
	"github.com/vacp2p/mvds/state"
)

type DataSyncNode interface {
	AppendMessage(groupID state.GroupID, data []byte) (state.MessageID, error)
	Subscribe(sub chan protobuf.Message)
	RequestMessage(group state.GroupID, id state.MessageID) error
}
