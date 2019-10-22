package client

import (
	"crypto/ecdsa"
	"crypto/rand"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"testing"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/vacp2p/dasy/client/internal"
	"github.com/vacp2p/dasy/crypto"
	"github.com/vacp2p/dasy/event"
	"github.com/vacp2p/dasy/protobuf"
	mvdsproto "github.com/vacp2p/mvds/protobuf"
	"github.com/vacp2p/mvds/state"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

// @todo think about turning feed into an interface so we can mock it and ensure its never called when sigs fail

func TestClient_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := internal.NewMockDataSyncNode(ctrl)
	identity, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)

	client := Client{
		node:         node,
		identity:     identity,
	}

	chat := Chat([32]byte{0x01, 0x2, 0x3, 0x4})
	msgid := state.MessageID([32]byte{0x01, 0x2, 0x3})

	node.EXPECT().AppendMessage(state.GroupID(chat), gomock.Any()).Return(msgid, nil)

	ret, _ := client.Post(chat, []byte("string"))
	if msgid != ret {
		t.Error("returned message ID does not match expected")
	}
}

func TestClient_Listen_MessageSentToFeed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := internal.NewMockDataSyncNode(ctrl)

	client := Client{
		node: node,
		feeds: make(map[protobuf.Message_MessageType]*event.Feed),
	}

	sub := make(chan mvdsproto.Message)
	node.EXPECT().Subscribe().Return(sub)

	go client.Listen()

	msg := createMessage()

	ok := make(chan event.Payload)
	client.Feed(msg.MessageType).Subscribe(ok)

	val, _ := proto.Marshal(msg)

	sub<-mvdsproto.Message{
		Body: val,
	}

	received := <-ok
	if !reflect.DeepEqual(received.Body, msg.Body) {
		t.Error("expected message did not equal received")
	}
}

func createMessage() *protobuf.Message {
	msg := &protobuf.Message{
		MessageType: protobuf.Message_POST,
		Body: []byte("hi"),
	}

	identity, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	_ = crypto.Sign(identity, msg)
	return msg
}
