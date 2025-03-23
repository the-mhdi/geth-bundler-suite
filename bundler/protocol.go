package bundler

import (
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

const (
	erc4337      = 1
	ProtocolName = "erc4337"
)

var ProtocolVersions = []uint{erc4337}

var protocolLengths = map[uint]uint64{erc4337: 17}

// MakeProtocols constructs the P2P protocol definitions for `eth`.
func MakeProtocols(backend Backend, network uint64, disc enode.Iterator) []p2p.Protocol {
	protocols := make([]p2p.Protocol, 0, len(ProtocolVersions))
	for _, version := range ProtocolVersions {
		protocols = append(protocols, p2p.Protocol{
			Name:    ProtocolName,
			Version: version,
			Length:  protocolLengths[version],
			Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
				peer := NewPeer(version, p, rw, backend.TxPool())
				defer peer.Close()

				return backend.RunPeer(peer, func(peer *Peer) error {
					return Handle(backend, peer)
				})
			},
			NodeInfo: func() interface{} {
				return nodeInfo(backend.Chain(), network)
			},
			PeerInfo: func(id enode.ID) interface{} {
				return backend.PeerInfo(id)
			},
			DialCandidates: disc,
			Attributes:     []enr.Entry{currentENREntry(backend.Chain())},
		})
	}
	return protocols
}
