package client

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vacp2p/dasy/client/internal"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	m.Run()
}

func TestClient_Listen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := internal.NewMockDataSyncNode(ctrl)

	_ := Client{
		node: node,
	}


	// @todo
}
