package mempool

import (
	"github.com/the-mhdi/geth-bundler-suite/bundler/types"
)

type memPool struct {
	Reputation *ReputationManager
	Validation *ValidationManager
}

func (p *memPool) SubmitTopool(up *types.UserOperation) error {

}
