package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
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
	EntryPoint           common.Address `json:"entryPoint"` // Optional, if you want to explicitly track entrypoint

}

func genUserOp() *UserOperation {
	return &UserOperation{}
}

func HandleUserOp() {}

/*
validation rules: Full validation rules must be applied between all Bundlers
*/

// MarshalJSON custom marshaller to handle uint256.Int as hex string
func (uop *UserOperation) MarshalJSON() ([]byte, error) {
	type Alias UserOperation
	return json.Marshal(&struct {
		Nonce string `json:"nonce"`
		Alias
	}{
		Nonce: uop.Nonce.Hex(),
		Alias: Alias(*uop),
	})
}

// UnmarshalJSON custom unmarshaller to handle hex string for uint256.Int
func (uop *UserOperation) UnmarshalJSON(data []byte) error {
	type Alias UserOperation
	aux := &struct {
		Nonce string `json:"nonce"`
		*Alias
	}{
		Alias: (*Alias)(uop),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if _, ok := uop.Nonce.SetFromHexString(aux.Nonce); !ok {
		// Should not happen as hex string is validated by uint256.Int.SetFromHexString
		// but return error just in case
		return &json.UnmarshalTypeError{
			Value:  "hex string",
			Type:   &uint256.Int{},
			Offset: 0, // You might want to adjust this based on your needs
			Struct: "UserOperation",
			Field:  "Nonce",
		}
	}
	return nil
}
