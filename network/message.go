package network

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/lynn9388/p2p"
	"log"
	"math/rand"
	"time"
)

type MsgType int32

const (
	HBPing MsgType = iota
	NodeJoin
	NodeExit
	NodePubkeySync
	TxRequest
	TxCommit
	TxNotify
	UpdateStateRequest
	UpdateStateReply
	UpdateStateConfirm
	TradePhaseNotify
	ConsensusPhaseNotify
)

type Message struct {
	MsgType MsgType
	Data    []byte
	From    string
	To      string
}

func checkStringMessage(msg *any.Any) (*any.Any, error) {
	var err error
	m := &wrappers.StringValue{}
	if err = ptypes.UnmarshalAny(msg, m); err != nil {
		return nil, err
	}
	log.Println(m.Value)

	reply, err := ptypes.MarshalAny(&wrappers.StringValue{Value: "reply:OK"})
	if err != nil {
		return nil, err
	}

	return reply, err

}

func SendMsg(node *p2p.Node, msg *Message, receiveAddr string) (err error) {
	var data []byte
	data, err = json.Marshal(msg)
	if err != nil {
		log.Panic("json marshal err, %v", err)
		return
	}

	for i := 0; i < 3; i++ {
		_, err := node.SendMessage(receiveAddr, &wrappers.BytesValue{Value: data}, 1*time.Second)
		if err != nil {
			log.Panic("[%s] send msg to [%s] err, %v", node.Addr, receiveAddr, err)
			time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
			continue
		}
		break
	}
	return
}
