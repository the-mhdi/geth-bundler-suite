package bundler

import (
	"github.com/ethereum/go-ethereum/p2p/enode"
)

type handler struct {
	nodeID     enode.ID //same as the geth node with "bnode" prefix
	networkID  uint64   // as the geth node
	uoPool     Pool
	maxPeers   int
	aggregator []byte
}
