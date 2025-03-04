package bundler

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type UserOperation struct {
	Sender               common.Address `json:"sender"               `
	Nonce                *big.Int       `json:"nonce"                `
	InitCode             []byte         `json:"initCode"             `
	CallData             []byte         `json:"callData"             `
	CallGasLimit         *big.Int       `json:"callGasLimit"         `
	VerificationGasLimit *big.Int       `json:"verificationGasLimit" `
	PreVerificationGas   *big.Int       `json:"preVerificationGas"   `
	MaxFeePerGas         *big.Int       `json:"maxFeePerGas"         `
	MaxPriorityFeePerGas *big.Int       `json:"maxPriorityFeePerGas" `
	PaymasterAndData     []byte         `json:"paymasterAndData"     `
	Signature            []byte         `json:"signature"            `
}

/*
validation rules: Full validation rules must be applied between all Bundlers

*/
func ()