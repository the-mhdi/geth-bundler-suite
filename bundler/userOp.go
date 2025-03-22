package bundler

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type UserOperation struct {
	Sender                        common.Address `json:"sender"               mapstructure:"sender"               validate:"required"`
	Nonce                         *big.Int       `json:"nonce"                mapstructure:"nonce"                validate:"required"`
	InitCode                      []byte         `json:"initCode"             mapstructure:"initCode"             validate:"required"`
	CallData                      []byte         `json:"callData"             mapstructure:"callData"             validate:"required"`
	CallGasLimit                  *big.Int       `json:"callGasLimit"         mapstructure:"callGasLimit"         validate:"required"`
	VerificationGasLimit          *big.Int       `json:"verificationGasLimit" mapstructure:"verificationGasLimit" validate:"required"`
	PreVerificationGas            *big.Int       `json:"preVerificationGas"   mapstructure:"preVerificationGas"   validate:"required"`
	MaxFeePerGas                  *big.Int       `json:"maxFeePerGas"         mapstructure:"maxFeePerGas"         validate:"required"`
	MaxPriorityFeePerGas          *big.Int       `json:"maxPriorityFeePerGas" mapstructure:"maxPriorityFeePerGas" validate:"required"`
	Paymaster                     common.Address `json:"Paymaster"    		 mapstructure:"Paymaster"    		 validate:"required"`
	PaymasterData                 []byte         `json:"PaymasterData"     mapstructure:"PaymasterData"     validate:"required"`
	paymasterVerificationGasLimit *big.Int       `json:"paymasterVerificationGasLimit"     mapstructure:"paymasterVerificationGasLimit"     validate:"required"`
	paymasterPostOpGasLimit       *big.Int       `json:"paymasterPostOpGasLimit"     mapstructure:"paymasterPostOpGasLimit"     validate:"required"`
	Signature                     []byte         `json:"signature"            mapstructure:"signature"            validate:"required"`
	Factory                       common.Address //empty if none
	FactoryData                   []byte
	EntryPoint                    common.Address
	ChainID                       uint64
	AuthorizationTuple            authorizationTuple // authorization tuple for 7702 txns
	Aggregator                    common.Address     //also known as "authorizer contract" - a contract that enables multiple UserOperations to share a single validation
}

type authorizationTuple struct {
	ChainID  uint64
	Address  common.Address
	Nonce    *big.Int
	y_parity uint32
	r        []byte
	s        []byte
}
