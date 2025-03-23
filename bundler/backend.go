package bundler

import "github.com/ethereum/go-ethereum/p2p/enode"

// Handler is a callback to invoke from an outside runner after the boilerplate
// exchanges have passed.
type Handler func(peer *Peer) error

type Backend interface {
	// userOpPool retrieves the transaction pool object to serve data.

	userOpPool() Pool

	// AcceptTxs retrieves whether transaction processing is enabled on the node
	// or if inbound transactions should simply be dropped.
	AcceptTxs() bool

	// RunPeer is invoked when a peer joins on the `erc4337` protocol. The handler
	// should do any peer maintenance work, handshakes and validations. If all
	// is passed, control should be given back to the `handler` to process the
	// inbound messages going forward.
	RunPeer(peer *Peer, handler Handler) error

	// PeerInfo retrieves all known `eth` information about a peer.
	PeerInfo(id enode.ID) any

	// Handle is a callback to be invoked when a data packet is received from
	// the remote peer. Only packets not consumed by the protocol handler will
	// be forwarded to the backend.
	Handle(peer *Peer, packet Packet) error
}
