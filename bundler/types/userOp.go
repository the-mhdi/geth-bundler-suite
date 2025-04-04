package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type UserOperation struct {
	Sender                        common.Address `json:"sender"`
	Nonce                         *big.Int       `json:"nonce"`
	InitCode                      []byte         `json:"initCode"`
	CallData                      []byte         `json:"callData"`
	CallGasLimit                  *big.Int       `json:"callGasLimit"`
	VerificationGasLimit          *big.Int       `json:"verificationGasLimit"`
	PreVerificationGas            *big.Int       `json:"preVerificationGas"`
	MaxFeePerGas                  *big.Int       `json:"maxFeePerGas"`
	MaxPriorityFeePerGas          *big.Int       `json:"maxPriorityFeePerGas"`
	Paymaster                     common.Address `json:"Paymaster"`
	PaymasterData                 []byte         `json:"PaymasterData"`
	PaymasterVerificationGasLimit *big.Int       `json:"paymasterVerificationGasLimit"`
	PaymasterPostOpGasLimit       *big.Int       `json:"paymasterPostOpGasLimit"`
	paymasterAndData              []byte         //or 0.6 ep only
	Signature                     []byte         `json:"signature"`
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

func (op *UserOperation) GetUserOpHash(entryPoint common.Address, chainID *big.Int) common.Hash {

}

func (op *UserOperation) Pack() PackedUserOperation {

}

func (op *UserOperation) getPaymasterAndData() []byte {

	//paymaster_verification_gas_limit + paymaster_post_op_gas_limit + paymaster_data +

	//todo:check if getpaymasteranddata can be formed

	//If PaymasterVerificationGasLimit or PaymasterPostOpGasLimit can be nil, you should add checks to handle those cases.

	// Convert PaymasterVerificationGasLimit and PaymasterPostOpGasLimit to byte slices
	paymasterVerificationGasLimitBytes := op.PaymasterVerificationGasLimit.Bytes()
	paymasterPostOpGasLimitBytes := op.PaymasterPostOpGasLimit.Bytes()

	// Concatenate the byte slices
	result := append(paymasterVerificationGasLimitBytes, paymasterPostOpGasLimitBytes...)
	result = append(result, op.PaymasterData...)

	return result

}
