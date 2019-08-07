package client

import (
	"crypto/ecdsa"
	"crypto/rand"
	"io/ioutil"
	"log"
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
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	m.Run()
}

// @todo think about turning feed into an interface so we can mock it and ensure its never called when sigs fail

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

func TestClient_Listen_RequestsMissingParent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := internal.NewMockDataSyncNode(ctrl)
	store := internal.NewMockMessageStore(ctrl)

	client := Client{
		node: node,
		store: store,
		feeds: make(map[protobuf.Message_MessageType]*event.Feed),
	}

	sub := make(chan mvdsproto.Message)
	node.EXPECT().Subscribe().Return(sub)

	store.EXPECT().Has(gomock.Any()).Return(false, nil)
	node.EXPECT().RequestMessage(gomock.Any(), gomock.Any()).Return(nil)

	go client.Listen()

	parent := []byte("parent")
	msg := createMessage()
	msg.PreviousMessage = parent

	ok := make(chan event.Payload)
	client.Feed(msg.MessageType).Subscribe(ok)

	val, _ := proto.Marshal(msg)

	sub<-mvdsproto.Message{
		Body: val,
	}

	<-ok
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
