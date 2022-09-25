/*
 * Copyright © 2022 Lynn <lynn9388@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/json"
	"flag"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/kingstenzzz/Turbo/network"
	"github.com/kingstenzzz/Turbo/peer"
	"github.com/lynn9388/p2p"
	"log"
	"strconv"
	"time"
)

var (
	testSendMsg  = "lynn"
	testReplyMsg = "9388"
)

var tests = []string{
	"127.0.0.1:9188",
	"127.0.0.1:9288",
	"127.0.0.1:9388",
	"127.0.0.1:9488",
}

func main() {
	port := flag.Int("port", 9188, "port for server")
	flag.String("to", tests[1], "port for to")
	nodeType := flag.String("peer", "v", "type of peer")
	flag.Parse()

	node := p2p.NewNode("127.0.0.1:" + strconv.Itoa(*port))
	if *nodeType == "l" {
		node.RegisterProcess(&wrappers.BytesValue{}, network.LeaderMessageHandler)

	} else {
		node.RegisterProcess(&wrappers.BytesValue{}, network.VeriferMessageHandler)
	}

	node.StartServer()
	defer node.StopServer()
	node.PeerManager.AddPeers(tests[0], tests[1], tests[2])
	if *nodeType == "v" {
		go func() {
			//message, err := ptypes.MarshalAny(&wrappers.StringValue{Value: "I am "+peer.Addr})
			for i := 0; i < 1000; i++ {
				signTx := peer.Tx{
					Id:       strconv.Itoa(i),
					Epoch:    i,
					Sender:   node.Addr,
					Receiver: tests[0],
					Amount:   1,
				}
				data, _ := json.Marshal(signTx)
				msg := network.Message{MsgType: network.TxRequest, Data: data, From: node.Addr, To: tests[0]}
				err := network.SendMsg(node, &msg, tests[0])
				time.Sleep(1 * time.Second)
				if err != nil {
					log.Print(err)
				}
			}
		}()
	}
	node.Wait()

}
