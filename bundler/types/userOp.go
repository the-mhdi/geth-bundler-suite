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

type PackedUserOperation struct {
	Sender             common.Address `json:"sender"`
	Nonce              *big.Int       `json:"nonce"`
	InitCode           []byte         `json:"initCode"`
	CallData           []byte         `json:"callData"`
	AccountGasLimits   []byte         `json:"accountGasLimits"` //concatenation of verificationGasLimit (16 bytes) and callGasLimit (16 bytes)
	PreVerificationGas *big.Int       `json:"preVerificationGas"`
	GasFees            []byte         `json:"gasFees"` //concatenation of maxPriorityFeePerGas (16 bytes) and maxFeePerGas (16 bytes)
	Signature          []byte         `json:"signature"`
	paymasterAndData   []byte         `json:"paymasterAndData"`
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

	//pack the UserOperation struct into PackedUserOperation struct

	packed := PackedUserOperation{}

	packed.Sender = op.Sender

	packed.Nonce = op.Nonce

	packed.InitCode = op.InitCode

	packed.CallData = op.CallData

	packed.PreVerificationGas = op.PreVerificationGas

	packed.AccountGasLimits = op.getAccountGasLimits()

	packed.GasFees = op.getGasFees()

	packed.paymasterAndData = op.getPaymasterAndData()

	return packed
}

func (op *UserOperation) Unpack(packed PackedUserOperation) UserOperation {

	//unpack the PackedUserOperation struct into UserOperation struct

	op.Sender = packed.Sender

	op.Nonce = packed.Nonce

	op.InitCode = packed.InitCode

	op.CallData = packed.CallData

	op.PreVerificationGas = packed.PreVerificationGas

	op.CallGasLimit = new(big.Int).SetBytes(packed.AccountGasLimits[16:])

	op.VerificationGasLimit = new(big.Int).SetBytes(packed.AccountGasLimits[:16])

	op.MaxPriorityFeePerGas = new(big.Int).SetBytes(packed.GasFees[16:])

	op.MaxFeePerGas = new(big.Int).SetBytes(packed.GasFees[:16])

	op.paymasterAndData = packed.paymasterAndData

	return *op
}

func (op *UserOperation) getAccountGasLimits() []byte {

	//verificationGasLimit + callGasLimit

	verificationGasLimitBytes := op.VerificationGasLimit.Bytes()

	callGasLimitBytes := op.CallGasLimit.Bytes()

	result := append(verificationGasLimitBytes, callGasLimitBytes...)

	return result

}

func (op *UserOperation) getGasFees() []byte {

	//maxPriorityFeePerGas + maxFeePerGas

	maxPriorityFeePerGasBytes := op.MaxPriorityFeePerGas.Bytes()

	maxFeePerGasBytes := op.MaxFeePerGas.Bytes()

	result := append(maxPriorityFeePerGasBytes, maxFeePerGasBytes...)

	return result

}

func (op *UserOperation) getPaymasterAndData() []byte {

	//paymaster_verification_gas_limit + paymaster_post_op_gas_limit + paymaster_data +

	//todo:check if getpaymasteranddata can be formed

	//If PaymasterVerificationGasLimit or PaymasterPostOpGasLimit can be nil, you should add checks to handle those cases.

	// Convert PaymasterVerificationGasLimit and PaymasterPostOpGasLimit to byte slices

	var result []byte

	paymasterAddressBytes := op.Paymaster.Bytes()

	paymasterVerificationGasLimitBytes := op.PaymasterVerificationGasLimit.Bytes()

	paymasterPostOpGasLimitBytes := op.PaymasterPostOpGasLimit.Bytes()

	result = append(result, paymasterAddressBytes...)
	result = append(result, paymasterVerificationGasLimitBytes...)
	result = append(result, paymasterPostOpGasLimitBytes...)
	result = append(result, op.PaymasterData...)

	// Concatenate the byte slices
	//result := append(paymasterVerificationGasLimitBytes, paymasterPostOpGasLimitBytes...)

	return result

}
