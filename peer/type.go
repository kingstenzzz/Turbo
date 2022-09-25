package peer

import (
	"github.com/lynn9388/p2p"
	"sync"
	"time"
)

type Tx struct {
	Id          string // tx id
	Epoch       int    // epoch of tx
	Sender      string // sender id
	Receiver    string // receiver id
	Amount      int    // tx amount
	LeaderSig   []byte // leader sig
	SenderSig   []byte // sender sig
	ReceiverSig []byte // receiver sig

	StartTime time.Time     // tx start time
	Duration  time.Duration // tx duration time
}
type Peer struct {
	Id           string    // node id
	ChannelName  string    // channel net name
	NodeName     string    // node name TODO: 有没有用，待最终确认
	NetAddr      string    // node network address
	Balance      int       // node balance
	Joined       bool      // node joined channel
	Exited       bool      // node exited channel
	stateConfirm chan bool // consensus state confirm

	state     *State  // node state
	leader    *Leader // only leader use
	txSet     []Tx    // an epoch tx set
	nodePhase Phase   // node phase

	//keyPair         *crypto.KeyPair        // node public and private key pair
	//nodeSignOpts    *cmCrypto.SignOpts     // Signing options
	otherNodePubKey map[string]*NodeEntity // exist node public key map

	//NodeLibp2pNet *libp2pnet.LibP2pNet //node network
	Node *p2p.Node

	sync.RWMutex
}

type NodeEntity struct {
	//PubKey *crypto.PubKeyEntity
	Seed string
}
type Leader struct {
	txSet            []Tx                // an epoch tx set
	enrollmentSet    map[string]int      // node enrollment set
	withdrawSet      map[string]struct{} // node exit set
	sendAmountSet    map[string]int      // all nodes send amount
	receiveAmountSet map[string]int      // all nodes receive amount
	txIdUsed         map[string]bool     // the epoch used txid
	newState         State
}

// State state
type State struct {
	Epoch       int
	LeaderId    string
	BalanceSet  map[string]int
	WithdrawSet map[string]struct{} // node exit set
	SigSet      map[string][]byte
}

// Phase phase
type Phase int

const (
	InitPhase        Phase = iota // init phase, no operation can be performed
	TradePhase                    // Trades can be made at this stage
	ConsensusPhase                // The trading phase timeout triggers entering the consensus phase
	ArbitrationPhase              // Unable to reach a consensus, all nodes enter the challenge phase
	ExitPhase                     // exit phase
)

type JoinData struct {
	//KeyType cmCrypto.KeyType
	PubKey []byte // user public key
	Seed   string // user p2p net seed
}
