package bundler

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//userops go through these phases : 1. local simulation of validateOp()

type pool struct {
	uop sync.Map
}

type NewPooledUserOperationHashesPacket struct {
	Types  []byte
	Sizes  []uint32
	Hashes []common.Hash
}

// GetPooledTransactionsRequest represents a transaction query.
type GetPooledUserOperationRequest []common.Hash

// GetPooledUserOperationRequest represents a uop query with request ID wrapping.
type GetPooledUserOperationPacket struct {
	RequestId uint64
	GetPooledUserOperationRequest
}

// PooledUserOperationResponse is the network packet for uop distribution.
type PooledUserOperationResponse []*types.Transaction

// PooledUserOperationPacket is the network packet for uop distribution
// with request ID wrapping.
type PooledUserOperationPacket struct {
	RequestId uint64
	PooledUserOperationResponse
}
