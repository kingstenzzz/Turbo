package peer

import (
	"encoding/json"
	"github.com/kingstenzzz/Turbo/network"
	"log"
	"os"
)

func (peer *Peer) TxRequestHandler(msg network.Message, from string) {
	var receiveTx Tx
	err := json.Unmarshal(msg.Data, &receiveTx)
	if err != nil {
		log.Panic("json unmarshal err, %v", err)
		os.Exit(1)

	}
	log.Printf("--- tx: %v ", receiveTx)
	signTx := Tx{
		Id:       "0",
		Epoch:    receiveTx.Epoch,
		Sender:   receiveTx.Sender,
		Receiver: receiveTx.Receiver,
		Amount:   receiveTx.Amount,
	}
	encodingSignedTx, _ := json.Marshal(signTx)
	if err != nil {
		log.Panic("json marshal error, %v", err)
		return
	}
	newMsg := network.Message{MsgType: network.TxCommit, Data: encodingSignedTx, From: peer.Id, To: msg.From}
	go network.SendMsg(peer.Node, &newMsg, msg.From)
}
