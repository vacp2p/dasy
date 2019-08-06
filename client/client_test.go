package client

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vacp2p/dasy/client/internal"
	mvdsproto "github.com/vacp2p/mvds/protobuf"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	m.Run()
}

func TestClient_Listen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := internal.NewMockDataSyncNode(ctrl)

	client := Client{
		node: node,
	}

	sub := make(chan mvdsproto.Message)
	node.EXPECT().Subscribe().Return(sub)

	go client.Listen()

	// @todo actual
	sub<-mvdsproto.Message{}

}
