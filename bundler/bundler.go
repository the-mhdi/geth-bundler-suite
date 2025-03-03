package bundler

import (
	"github.com/ethereum/go-ethereum/common"
)

// Bundler strcut controls the process of adding incoming UserOperations to the mempool.
type Bundler struct {
}

type BundlerService struct {
	bundler *Bundler
}

func (b *BundlerService) eth_sendUserOperation(userOp UserOperation, entryPoint common.Address) (string, error) {

	return UserOpHash, nil
}
