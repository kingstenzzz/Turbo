package network

import (
	"Turbo/peer"
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"log"
)

func LeaderMessageHandler(msg *any.Any) (*any.Any, error) {
	var err error
	var replyMessage string
	m := &wrappers.BytesValue{}
	if err = ptypes.UnmarshalAny(msg, m); err != nil {
		return nil, err
	}
	var message Message
	errHandler := json.Unmarshal(m.Value, &message)
	if errHandler != nil {
		log.Panic("json unmarshal err = ", errHandler)
		return nil, nil
	}
	switch message.MsgType {
	case TxRequest:
		go peer.TxRequestHandler(message, message.From)
	case UpdateStateRequest:
		replyMessage = "vote receipt"
	}
	reply, err := ptypes.MarshalAny(&wrappers.StringValue{Value: replyMessage})
	if err != nil {
		return nil, err
	}

	return reply, err
}

func VeriferMessageHandler(msg *any.Any) (*any.Any, error) {
	var err error
	var replyMessage string
	m := &wrappers.BytesValue{}
	if err = ptypes.UnmarshalAny(msg, m); err != nil {
		return nil, err
	}
	var message Message
	errHandler := json.Unmarshal(m.Value, &message)
	if errHandler != nil {
		log.Panic("json unmarshal err = ", errHandler)
		return nil, nil
	}
	switch message.MsgType {
	case TxNotify:
		log.Println("receipt from leader")
	case UpdateStateConfirm:
		log.Println("update")
	}
	reply, err := ptypes.MarshalAny(&wrappers.StringValue{Value: replyMessage})
	if err != nil {
		return nil, err
	}

	return reply, err
}
