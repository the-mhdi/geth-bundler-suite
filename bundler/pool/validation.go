package mempool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/the-mhdi/geth-bundler-suite/bundler/types"
)

type ValidationManager interface {
	isValid() bool
	userOperation() *types.UserOperation
	atBlockHash() common.Hash
}
