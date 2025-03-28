package mempool

import (
	"github.com/the-mhdi/geth-bundler-suite/bundler/types"
)

type Pool interface {
	eth.txPool
}

// memPool implements SubPool interface of geth client as specified in its source code

type memPool struct {
	Reputation *ReputationManager
	Validation *ValidationManager
}

func (p *memPool) SubmitTopool(up *types.UserOperation) error {

}
