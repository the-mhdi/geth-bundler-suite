package bundler

import (
	"github.com/ethereum/go-ethereum/p2p/enode"
)

type node struct { //b node handler
	nodeID     enode.ID //same as the geth node with "bnode" prefix
	networkID  uint64   // as the geth node
	uoPool     Pool
	peerSet    map[string]*Peer //id - peer map
	maxPeers   int
	aggregator []byte
}
