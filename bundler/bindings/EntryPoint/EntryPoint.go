// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EntryPoint

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// EntryPointMemoryUserOp is an auto generated low-level Go binding around an user-defined struct.
type EntryPointMemoryUserOp struct {
	Sender                        common.Address
	Nonce                         *big.Int
	VerificationGasLimit          *big.Int
	CallGasLimit                  *big.Int
	PaymasterVerificationGasLimit *big.Int
	PaymasterPostOpGasLimit       *big.Int
	PreVerificationGas            *big.Int
	Paymaster                     common.Address
	MaxFeePerGas                  *big.Int
	MaxPriorityFeePerGas          *big.Int
}

// EntryPointUserOpInfo is an auto generated low-level Go binding around an user-defined struct.
type EntryPointUserOpInfo struct {
	MUserOp       EntryPointMemoryUserOp
	UserOpHash    [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
}

// IEntryPointUserOpsPerAggregator is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointUserOpsPerAggregator struct {
	UserOps    []PackedUserOperation
	Aggregator common.Address
	Signature  []byte
}

// IStakeManagerDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerDepositInfo struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// EIP712MetaData contains all meta data concerning the EIP712 contract.
var EIP712MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "ca3ae4476606bcb6cf58a94d4bf9e3d113",
}

// EIP712 is an auto generated Go binding around an Ethereum contract.
type EIP712 struct {
	abi abi.ABI
}

// NewEIP712 creates a new instance of EIP712.
func NewEIP712() *EIP712 {
	parsed, err := EIP712MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &EIP712{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *EIP712) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eIP712 *EIP712) PackEip712Domain() []byte {
	enc, err := eIP712.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eIP712 *EIP712) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eIP712.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, err

}

// EIP712EIP712DomainChanged represents a EIP712DomainChanged event raised by the EIP712 contract.
type EIP712EIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const EIP712EIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (EIP712EIP712DomainChanged) ContractEventName() string {
	return EIP712EIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eIP712 *EIP712) UnpackEIP712DomainChangedEvent(log *types.Log) (*EIP712EIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eIP712.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EIP712EIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eIP712.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eIP712.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (eIP712 *EIP712) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eIP712.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return eIP712.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], eIP712.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return eIP712.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// EIP712InvalidShortString represents a InvalidShortString error raised by the EIP712 contract.
type EIP712InvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func EIP712InvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (eIP712 *EIP712) UnpackInvalidShortStringError(raw []byte) (*EIP712InvalidShortString, error) {
	out := new(EIP712InvalidShortString)
	if err := eIP712.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// EIP712StringTooLong represents a StringTooLong error raised by the EIP712 contract.
type EIP712StringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func EIP712StringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (eIP712 *EIP712) UnpackStringTooLongError(raw []byte) (*EIP712StringTooLong, error) {
	out := new(EIP712StringTooLong)
	if err := eIP712.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "5eb3ce5f671914ab9adbddb5d74ba39063",
}

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	abi abi.ABI
}

// NewERC165 creates a new instance of ERC165.
func NewERC165() *ERC165 {
	parsed, err := ERC165MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC165{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC165) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC165 *ERC165) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC165.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC165 *ERC165) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC165.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// Eip7702SupportMetaData contains all meta data concerning the Eip7702Support contract.
var Eip7702SupportMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "cf9a9576c70e8ec78c1f014522c9f74bbe",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220b398a5cf359a8010e24e9ee8bd5cd27f17e25dace1192861c8c99e109f94ccc464736f6c634300081d0033",
}

// Eip7702Support is an auto generated Go binding around an Ethereum contract.
type Eip7702Support struct {
	abi abi.ABI
}

// NewEip7702Support creates a new instance of Eip7702Support.
func NewEip7702Support() *Eip7702Support {
	parsed, err := Eip7702SupportMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Eip7702Support{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Eip7702Support) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// EntryPointMetaData contains all meta data concerning the EntryPoint contract.
var EntryPointMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"name\":\"DelegateAndRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"PostOpReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureValidationFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureAggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparatorV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPackedUserOpTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"getSenderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAggregator\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.UserOpsPerAggregator[]\",\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleAggregatedOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymasterVerificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymasterPostOpGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.MemoryUserOp\",\"name\":\"mUserOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contextOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.UserOpInfo\",\"name\":\"opInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"innerHandleOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"nonceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"senderCreator\",\"outputs\":[{\"internalType\":\"contractISenderCreator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "6e55cf931561549947fc59af494e9417ac",
	Bin: "0x6101806040526040516100119061029a565b604051809103905ff08015801561002a573d5f5f3e3d5ffd5b5073ffffffffffffffffffffffffffffffffffffffff166101609073ffffffffffffffffffffffffffffffffffffffff16815250348015610069575f5ffd5b506040518060400160405280600781526020017f45524334333337000000000000000000000000000000000000000000000000008152506040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506100ea60028361018560201b90919060201c565b610120818152505061010660038261018560201b90919060201c565b6101408181525050818051906020012060e08181525050808051906020012061010081815250504660a081815250506101436101d260201b60201c565b608081815250503073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250505050610783565b5f6020835110156101a65761019f8361022c60201b60201c565b90506101cc565b826101b68361029160201b60201c565b5f0190816101c491906104e4565b5060ff5f1b90505b92915050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60e051610100514630604051602001610211959493929190610619565b60405160208183030381529060405280519060200120905090565b5f5f829050601f8151111561027857826040517f305a27a900000000000000000000000000000000000000000000000000000000815260040161026f91906106d0565b60405180910390fd5b8051816102849061071d565b5f1c175f1b915050919050565b5f819050919050565b61092a8061792083390190565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061032257607f821691505b602082108103610335576103346102de565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261035c565b6103a1868361035c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6103e56103e06103db846103b9565b6103c2565b6103b9565b9050919050565b5f819050919050565b6103fe836103cb565b61041261040a826103ec565b848454610368565b825550505050565b5f5f905090565b61042961041a565b6104348184846103f5565b505050565b5b818110156104575761044c5f82610421565b60018101905061043a565b5050565b601f82111561049c5761046d8161033b565b6104768461034d565b81016020851015610485578190505b6104996104918561034d565b830182610439565b50505b505050565b5f82821c905092915050565b5f6104bc5f19846008026104a1565b1980831691505092915050565b5f6104d483836104ad565b9150826002028217905092915050565b6104ed826102a7565b67ffffffffffffffff811115610506576105056102b1565b5b610510825461030b565b61051b82828561045b565b5f60209050601f83116001811461054c575f841561053a578287015190505b61054485826104c9565b8655506105ab565b601f19841661055a8661033b565b5f5b828110156105815784890151825560018201915060208501945060208101905061055c565b8683101561059e578489015161059a601f8916826104ad565b8355505b6001600288020188555050505b505050505050565b5f819050919050565b6105c5816105b3565b82525050565b6105d4816103b9565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610603826105da565b9050919050565b610613816105f9565b82525050565b5f60a08201905061062c5f8301886105bc565b61063960208301876105bc565b61064660408301866105bc565b61065360608301856105cb565b610660608083018461060a565b9695505050505050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6106a2826102a7565b6106ac818561066a565b93506106bc81856020860161067a565b6106c581610688565b840191505092915050565b5f6020820190508181035f8301526106e88184610698565b905092915050565b5f81519050919050565b5f819050602082019050919050565b5f61071482516105b3565b80915050919050565b5f610727826106f0565b82610731846106fa565b905061073c81610709565b9250602082101561077c576107777fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8360200360080261035c565b831692505b5050919050565b60805160a05160c05160e051610100516101205161014051610160516171426107de5f395f610c8d01525f61296501525f61292a01525f612b2f01525f612b0e01525f61212a01525f61218001525f6121a901526171425ff3fe60806040526004361061012c575f3560e01c806335567e1a116100aa578063850aaf621161006e578063850aaf62146104265780639b249f691461044e578063b760faf914610476578063bb9fe6bf14610492578063c23a5cea146104a8578063dbed18e0146104d05761013c565b806335567e1a1461031a5780635287ce121461035657806370a0823114610392578063765e827f146103ce57806384b0196e146103f65761013c565b806313c65a6e116100f157806313c65a6e14610226578063154e58dc146102505780631b2e01b81461027a578063205c2878146102b657806322cdde4c146102de5761013c565b806242dc531461014057806301ffc9a71461017c5780630396cb60146101b857806309ccb880146101d45780630bd28e3b146101fe5761013c565b3661013c5761013a336104f8565b005b5f5ffd5b34801561014b575f5ffd5b5061016660048036038101906101619190614616565b610557565b60405161017391906146b4565b60405180910390f35b348015610187575f5ffd5b506101a2600480360381019061019d9190614722565b610746565b6040516101af9190614767565b60405180910390f35b6101d260048036038101906101cd91906147b9565b61093b565b005b3480156101df575f5ffd5b506101e8610c8a565b6040516101f5919061483f565b60405180910390f35b348015610209575f5ffd5b50610224600480360381019061021f91906148a5565b610cb1565b005b348015610231575f5ffd5b5061023a610d49565b60405161024791906148df565b60405180910390f35b34801561025b575f5ffd5b50610264610d57565b60405161027191906148df565b60405180910390f35b348015610285575f5ffd5b506102a0600480360381019061029b91906148f8565b610d7e565b6040516102ad91906146b4565b60405180910390f35b3480156102c1575f5ffd5b506102dc60048036038101906102d79190614971565b610d9e565b005b3480156102e9575f5ffd5b5061030460048036038101906102ff91906149d2565b610f39565b60405161031191906148df565b60405180910390f35b348015610325575f5ffd5b50610340600480360381019061033b91906148f8565b610f71565b60405161034d91906146b4565b60405180910390f35b348015610361575f5ffd5b5061037c60048036038101906103779190614a19565b61101b565b6040516103899190614b1f565b60405180910390f35b34801561039d575f5ffd5b506103b860048036038101906103b39190614a19565b61112a565b6040516103c591906146b4565b60405180910390f35b3480156103d9575f5ffd5b506103f460048036038101906103ef9190614b8d565b611171565b005b348015610401575f5ffd5b5061040a611297565b60405161041d9796959493929190614d3b565b60405180910390f35b348015610431575f5ffd5b5061044c60048036038101906104479190614dbd565b61133c565b005b348015610459575f5ffd5b50610474600480360381019061046f9190614e1a565b6113e6565b005b610490600480360381019061048b9190614a19565b6104f8565b005b34801561049d575f5ffd5b506104a66114a8565b005b3480156104b3575f5ffd5b506104ce60048036038101906104c99190614e65565b611650565b005b3480156104db575f5ffd5b506104f660048036038101906104f19190614ee5565b611943565b005b5f6105038234611d83565b90508173ffffffffffffffffffffffffffffffffffffffff167f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c48260405161054b91906146b4565b60405180910390a25050565b5f5f5a90503073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c190614f8c565b60405180910390fd5b5f855f015190505f816060015190506127108260a001518201016040603f5a02816105f8576105f7614faa565b5b041015610627577fdeaddead000000000000000000000000000000000000000000000000000000005f5260205ffd5b5f5f90505f895111156106dc575f610644845f01515f8c86611ddf565b9050806106da575f610654611df6565b90505f610662610800611dff565b90505f815111156106ca57855f015173ffffffffffffffffffffffffffffffffffffffff168b602001517f1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a2018860200151846040516106c1929190615029565b60405180910390a35b6106d382611e37565b6001935050505b505b5f88608001515a8603019050610737828a8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505084611e3e565b95505050505050949350505050565b5f7f3e84f021000000000000000000000000000000000000000000000000000000007fcf28ef97000000000000000000000000000000000000000000000000000000007f989ccc580000000000000000000000000000000000000000000000000000000018187bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061085457507f989ccc58000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806108bc57507fcf28ef97000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061092457507f3e84f021000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806109345750610933826120be565b5b9050919050565b5f5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f8263ffffffff16116109c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b9906150a1565b60405180910390fd5b80600101600f9054906101000a900463ffffffff1663ffffffff168263ffffffff161015610a25576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1c90615109565b60405180910390fd5b5f348260010160019054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16610a609190615154565b90505f8111610aa4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9b906151d1565b60405180910390fd5b6dffffffffffffffffffffffffffff8016811115610af7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aee90615239565b60405180910390fd5b6040518060a00160405280835f01548152602001600115158152602001826dffffffffffffffffffffffffffff1681526020018463ffffffff1681526020015f65ffffffffffff168152505f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f01556020820151816001015f6101000a81548160ff02191690831515021790555060408201518160010160016101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff160217905550606082015181600101600f6101000a81548163ffffffff021916908363ffffffff16021790555060808201518160010160136101000a81548165ffffffffffff021916908365ffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff167fa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c018285604051610c7d929190615287565b60405180910390a2505050565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8277ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190610d41906152ae565b919050555050565b5f610d52612127565b905090565b5f7f29a0bca4af4be3421398da00295e58e6d7de38cb492214754cb6a47507dd6f8e905090565b6001602052815f5260405f20602052805f5260405f205f91509150505481565b5f5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f815f0154905080831115610e27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1e9061533f565b60405180910390fd5b8281610e33919061535d565b825f01819055503373ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb8585604051610e829291906153b0565b60405180910390a25f8473ffffffffffffffffffffffffffffffffffffffff1684604051610eaf90615404565b5f6040518083038185875af1925050503d805f8114610ee9576040519150601f19603f3d011682016040523d82523d5f602084013e610eee565b606091505b5050905080610f32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2990615462565b60405180910390fd5b5050505050565b5f5f610f44836121dd565b9050610f69610f51610d49565b610f6483866122ae90919063ffffffff16565b6122c8565b915050919050565b5f60408277ffffffffffffffffffffffffffffffffffffffffffffffff16901b60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8477ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205417905092915050565b6110236140f3565b5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060a00160405290815f8201548152602001600182015f9054906101000a900460ff161515151581526020016001820160019054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16815260200160018201600f9054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160139054906101000a900465ffffffffffff1665ffffffffffff1665ffffffffffff16815250509050919050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f01549050919050565b611179612308565b5f8383905090505f8167ffffffffffffffff81111561119b5761119a614250565b5b6040519080825280602002602001820160405280156111d457816020015b6111c161413d565b8152602001906001900390816111b95790505b5090506111e48585835f5f612387565b505f5f90507fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f97260405160405180910390a15f5f90505b8381101561127c5761126b8188888481811061123957611238615480565b5b905060200281019061124b91906154b9565b85848151811061125e5761125d615480565b5b602002602001015161241a565b82019150808060010191505061121a565b5061128784826127c7565b5050506112926128e2565b505050565b5f6060805f5f5f60606112a8612921565b6112b061295c565b46305f5f1b5f67ffffffffffffffff8111156112cf576112ce614250565b5b6040519080825280602002602001820160405280156112fd5781602001602082028036833780820191505090505b507f0f00000000000000000000000000000000000000000000000000000000000000959493929190965096509650965096509650965090919293949596565b5f5f8473ffffffffffffffffffffffffffffffffffffffff168484604051611365929190615505565b5f60405180830381855af49150503d805f811461139d576040519150601f19603f3d011682016040523d82523d5f602084013e6113a2565b606091505b509150915081816040517f994105540000000000000000000000000000000000000000000000000000000081526004016113dd92919061551d565b60405180910390fd5b5f6113ef610c8a565b73ffffffffffffffffffffffffffffffffffffffff1663570e1a3684846040518363ffffffff1660e01b8152600401611429929190615577565b6020604051808303815f875af1158015611445573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061146991906155ad565b9050806040517f6ca7b80600000000000000000000000000000000000000000000000000000000815260040161149f91906155d8565b60405180910390fd5b5f5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f81600101600f9054906101000a900463ffffffff1663ffffffff1603611543576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161153a9061563b565b60405180910390fd5b806001015f9054906101000a900460ff16611593576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161158a906156a3565b60405180910390fd5b5f81600101600f9054906101000a900463ffffffff1663ffffffff16426115ba91906156c1565b9050808260010160136101000a81548165ffffffffffff021916908365ffffffffffff1602179055505f826001015f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167ffa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a82604051611644919061572a565b60405180910390a25050565b5f5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f8160010160019054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff1690505f8111611703576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116fa9061578d565b60405180910390fd5b5f8260010160139054906101000a900465ffffffffffff1665ffffffffffff1611611763576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161175a906157f5565b60405180910390fd5b428260010160139054906101000a900465ffffffffffff1665ffffffffffff1611156117c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117bb9061585d565b60405180910390fd5b5f82600101600f6101000a81548163ffffffff021916908363ffffffff1602179055505f8260010160136101000a81548165ffffffffffff021916908365ffffffffffff1602179055505f8260010160016101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3848360405161188d9291906153b0565b60405180910390a25f8373ffffffffffffffffffffffffffffffffffffffff16826040516118ba90615404565b5f6040518083038185875af1925050503d805f81146118f4576040519150601f19603f3d011682016040523d82523d5f602084013e6118f9565b606091505b505090508061193d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611934906158c5565b60405180910390fd5b50505050565b61194b612308565b5f8383905090505f5f90505f5f90505b82811015611b2d573686868381811061197757611976615480565b5b905060200281019061198991906158e3565b9050365f82805f019061199c919061590a565b915091505f8360200160208101906119b491906159a7565b9050600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158190611a29576040517f86a9f750000000000000000000000000000000000000000000000000000000008152600401611a2091906155d8565b60405180910390fd5b505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611b14578073ffffffffffffffffffffffffffffffffffffffff16632dd811338484878060400190611a8b91906159d2565b6040518563ffffffff1660e01b8152600401611aaa9493929190615d37565b5f604051808303815f87803b158015611ac1575f5ffd5b505af1925050508015611ad2575060015b611b1357806040517f86a9f750000000000000000000000000000000000000000000000000000000008152600401611b0a91906155d8565b60405180910390fd5b5b828290508601955050505050808060010191505061195b565b505f8167ffffffffffffffff811115611b4957611b48614250565b5b604051908082528060200260200182016040528015611b8257816020015b611b6f61413d565b815260200190600190039081611b675790505b5090505f5f90505f5f90505b84811015611c0b5736888883818110611baa57611ba9615480565b5b9050602002810190611bbc91906158e3565b9050365f82805f0190611bcf919061590a565b915091505f836020016020810190611be791906159a7565b9050611bf6838389848a612387565b86019550505050508080600101915050611b8e565b507fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f97260405160405180910390a15f5f90505f91505f5f90505b85811015611d665736898983818110611c6057611c5f615480565b5b9050602002810190611c7291906158e3565b9050806020016020810190611c8791906159a7565b73ffffffffffffffffffffffffffffffffffffffff167f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d60405160405180910390a2365f82805f0190611cda919061590a565b915091505f8282905090505f5f90505b81811015611d5457611d3b88858584818110611d0957611d08615480565b5b9050602002810190611d1b91906154b9565b8b8b81518110611d2e57611d2d615480565b5b602002602001015161241a565b8701965087806001019850508080600101915050611cea565b50505050508080600101915050611c44565b50611d7186826127c7565b5050505050611d7e6128e2565b505050565b5f5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f83825f015401905080825f0181905550809250505092915050565b5f5f5f845160208601878987f19050949350505050565b5f604051905090565b60603d5f831115611e175782811115611e16578290505b5b604051602082018101604052818152815f602083013e8092505050919050565b8060405250565b5f5f5a90505f5f865f015190505f611e5582612997565b90505f8260e0015190505f896080015188039050611e778185606001516129bf565b88019750505f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611ebb57835f01519450611fd2565b8194505f89511115611fd15782880296505f5a9050600280811115611ee357611ee2615d70565b5b8c6002811115611ef657611ef5615d70565b5b14611fb8578273ffffffffffffffffffffffffffffffffffffffff16637c627b218660a001518e8d8c896040518663ffffffff1660e01b8152600401611f3f9493929190615de3565b5f604051808303815f88803b158015611f56575f5ffd5b5087f193505050508015611f68575060015b611fb7575f611f78610800611dff565b9050806040517fad7954bc000000000000000000000000000000000000000000000000000000008152600401611fae9190615e2d565b60405180910390fd5b5b5f5a82039050611fcc818760a001516129bf565b925050505b5b805a8703018801975082880296505f8a604001519050878110156120655760028081111561200357612002615d70565b5b8c600281111561201657612015615d70565b5b03612038578097506120278b6129fd565b6120338b5f8a8c612a60565b612060565b7fdeadaa51000000000000000000000000000000000000000000000000000000005f5260205ffd5b6120af565b5f88820390506120758782611d83565b505f5f600281111561208a57612089615d70565b5b8e600281111561209d5761209c615d70565b5b1490506120ac8d828c8e612a60565b50505b50505050505050949350505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480156121a257507f000000000000000000000000000000000000000000000000000000000000000046145b156121cf577f000000000000000000000000000000000000000000000000000000000000000090506121da565b6121d7612aea565b90505b90565b5f365f8380604001906121f091906159d2565b915091506121fe8282612b7f565b61220e575f5f1b925050506122a9565b5f612229855f0160208101906122249190614a19565b612bea565b90506014838390501161226757806040516020016122479190615e92565b6040516020818303038152906040528051906020012093505050506122a9565b808383601490809261227b93929190615eb4565b60405160200161228d93929190615eee565b6040516020818303038152906040528051906020012093505050505b919050565b5f6122b98383612ce7565b80519060200120905092915050565b5f6040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b612310612ddf565b15612347576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61238560016123777f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b612e18565b612e2190919063ffffffff16565b565b5f8585905090505f5f90505b81811015612410575f85828501815181106123b1576123b0615480565b5b602002602001015190505f5f6123ee8487018b8b878181106123d6576123d5615480565b5b90506020028101906123e891906154b9565b85612e28565b9150915061240084870183838a61302e565b5050508080600101915050612393565b5095945050505050565b5f5f5a90505f61242d84606001516131be565b90505f5f612439611df6565b9050365f88806060019061244d91906159d2565b9150915060605f82600381111561246357843591505b50638dd7712f60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603612589575f8b8b602001516040516024016124c892919061603e565b604051602081830303815290604052638dd7712f60e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090503073ffffffffffffffffffffffffffffffffffffffff166242dc53828d8b60405160240161253f939291906161a1565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050925050612600565b3073ffffffffffffffffffffffffffffffffffffffff166242dc5385858d8b6040516024016125bb94939291906161e6565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505091505b60205f8351602085015f305af195505f51985061261c85611e37565b5050505050806127bd575f3d8060200361263a5760205f5f3e5f5191505b507fdeaddead00000000000000000000000000000000000000000000000000000000810361269f57876040517f220266b60000000000000000000000000000000000000000000000000000000081526004016126969190616277565b60405180910390fd5b7fdeadaa5100000000000000000000000000000000000000000000000000000000810361270b575f86608001515a866126d8919061535d565b6126e29190615154565b90505f876040015190506126f5886129fd565b612701885f8385612a60565b80965050506127bb565b5f612714611df6565b9050865f01515f015173ffffffffffffffffffffffffffffffffffffffff1687602001517ff62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792895f01516020015161276c610800611dff565b60405161277a929190615029565b60405180910390a361278b81611e37565b5f87608001515a8761279d919061535d565b6127a79190615154565b90506127b66002898784611e3e565b965050505b505b5050509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612835576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161282c906162ed565b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff168260405161285a90615404565b5f6040518083038185875af1925050503d805f8114612894576040519150601f19603f3d011682016040523d82523d5f602084013e612899565b606091505b50509050806128dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128d490616355565b60405180910390fd5b505050565b61291f5f6129117f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b612e18565b612e2190919063ffffffff16565b565b606061295760027f00000000000000000000000000000000000000000000000000000000000000006131c890919063ffffffff16565b905090565b606061299260037f00000000000000000000000000000000000000000000000000000000000000006131c890919063ffffffff16565b905090565b5f5f82610100015190505f83610120015190506129b682488301613275565b92505050919050565b5f619c40830182116129d3575f90506129f7565b5f83830390505f6064600a8302816129ee576129ed614faa565b5b04905080925050505b92915050565b805f01515f015173ffffffffffffffffffffffffffffffffffffffff1681602001517f67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e835f015160200151604051612a5591906146b4565b60405180910390a350565b835f015160e0015173ffffffffffffffffffffffffffffffffffffffff16845f01515f015173ffffffffffffffffffffffffffffffffffffffff1685602001517f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f875f015160200151878787604051612adc9493929190616373565b60405180910390a450505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000004630604051602001612b649594939291906163b6565b60405160208183030381529060405280519060200120905090565b5f6002838390501015612b94575f9050612be4565b5f8335905061770260f01b7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166bffffffffffffffffffffffff1916816bffffffffffffffffffffffff1916149150505b92915050565b5f5f60175f5f853c5f51905062ef010060e81b7cffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817cffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614612cd7575f8373ffffffffffffffffffffffffffffffffffffffff163b11612c9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c9390616451565b60405180910390fd5b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cce906164b9565b60405180910390fd5b601881901b60601c915050919050565b60605f835f016020810190612cfc9190614a19565b90505f846020013590505f5f5f1b8503612d2d57612d28868060400190612d2391906159d2565b61328d565b612d2f565b845b90505f612d4a878060600190612d4591906159d2565b61328d565b90505f876080013590505f8860a0013590505f8960c0013590505f612d7d8b8060e00190612d7891906159d2565b61328d565b90507f29a0bca4af4be3421398da00295e58e6d7de38cb492214754cb6a47507dd6f8e8888888888888888604051602001612dc0999897969594939291906164d7565b6040516020818303038152906040529850505050505050505092915050565b5f612e13612e0e7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b612e18565b6132a3565b905090565b5f819050919050565b80825d5050565b5f5f5f5a90505f845f01519050612e3f86826132ad565b5f612e48611df6565b9050612e5387610f39565b866020018181525050612e6581611e37565b5f826040015190505f8361012001518461010001518560a0015186608001518760600151868960c0015117171717171790506effffffffffffffffffffffffffffff80168111158a90612eee576040517f220266b6000000000000000000000000000000000000000000000000000000008152600401612ee591906165ac565b60405180910390fd5b505f612ef985613488565b905080896040018181525050612f118b8b8b846134b9565b9750612f24855f015186602001516135db565b8b90612f66576040517f220266b6000000000000000000000000000000000000000000000000000000008152600401612f5d9190616622565b60405180910390fd5b50825a87031115612fae578a6040517f220266b6000000000000000000000000000000000000000000000000000000008152600401612fa59190616698565b60405180910390fd5b60605f73ffffffffffffffffffffffffffffffffffffffff168660e0015173ffffffffffffffffffffffffffffffffffffffff1614612ffb57612ff28c8c8c613692565b80995081925050505b6130048161376a565b8a60600181815250508a60a001355a8803018a608001818152505050505050505050935093915050565b5f5f61303985613773565b915091508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146130ad57856040517f220266b60000000000000000000000000000000000000000000000000000000081526004016130a4919061670e565b60405180910390fd5b80156130f057856040517f220266b60000000000000000000000000000000000000000000000000000000081526004016130e79190616784565b60405180910390fd5b5f6130fa85613773565b80935081925050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461317257866040517f220266b600000000000000000000000000000000000000000000000000000000815260040161316991906167fa565b60405180910390fd5b81156131b557866040517f220266b60000000000000000000000000000000000000000000000000000000081526004016131ac9190616896565b60405180910390fd5b50505050505050565b6060819050919050565b606060ff5f1b83146131e4576131dd836137c8565b905061326f565b8180546131f0906168ef565b80601f016020809104026020016040519081016040528092919081815260200182805461321c906168ef565b80156132675780601f1061323e57610100808354040283529160200191613267565b820191905f5260205f20905b81548152906001019060200180831161324a57829003601f168201915b505050505090505b92915050565b5f8183106132835781613285565b825b905092915050565b5f60405182808583378082209250505092915050565b5f815c9050919050565b815f0160208101906132bf9190614a19565b815f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508160200135816020018181525050613310826080013561383a565b8260400183606001828152508281525050508160a001358160c001818152505061333d8260c0013561383a565b8261012001836101000182815250828152505050365f838060e0019061336391906159d2565b915091505f8282905011156134825760348282905010156133b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133b090616969565b60405180910390fd5b5f6133c48383613857565b866080018760a00182815250828152508293505050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613448576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161343f906169d1565b60405180910390fd5b808460e0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050505b50505050565b5f5f8260c001518360a001518460800151856060015186604001510101010190508261010001518102915050919050565b5f5f835f015190505f815f015190506134e287868880604001906134dd91906159d2565b6138f6565b5f8260e0015190505f5f90505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603613542575f61352b8461112a565b905086811161353c5780870361353e565b5f5b9150505b61354e89898984613caf565b94505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036135cf5761358d8387613e10565b6135ce57886040517f220266b60000000000000000000000000000000000000000000000000000000081526004016135c59190616a39565b60405180910390fd5b5b50505050949350505050565b5f5f604083901c90505f8390508067ffffffffffffffff1660015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8477ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190613683906152ae565b91905055149250505092915050565b60605f5f5a90505f845f015190505f8160e0015190505f866040015190506136ba8282613e10565b6136fb57886040517f220266b60000000000000000000000000000000000000000000000000000000081526004016136f29190616aaf565b60405180910390fd5b5f8360800151905061370e8a8a8a613e80565b8097508198505050805a8603111561375d57896040517f220266b60000000000000000000000000000000000000000000000000000000081526004016137549190616b4b565b60405180910390fd5b5050505050935093915050565b5f819050919050565b5f5f5f8303613787575f5f915091506137c3565b5f61379184613fc2565b9050806040015165ffffffffffff164211806137b95750806020015165ffffffffffff164211155b9150805f01519250505b915091565b60605f6137d483614042565b90505f602067ffffffffffffffff8111156137f2576137f1614250565b5b6040519080825280601f01601f1916602001820160405280156138245781602001600182028036833780820191505090505b5090508181528360208201528092505050919050565b5f5f61384583614090565b61384e8461409f565b91509150915091565b5f5f5f84845f9060149261386d93929190615eb4565b906138789190616bb8565b60601c858560149060249261388f93929190615eb4565b9061389a9190616c41565b60801c86866024906034926138b193929190615eb4565b906138bc9190616c41565b60801c816fffffffffffffffffffffffffffffffff169150806fffffffffffffffffffffffffffffffff1690509250925092509250925092565b5f8282905014613ca8575f835f01515f015190506139148383612b7f565b156139b75760148383905011156139b15761392d610c8a565b73ffffffffffffffffffffffffffffffffffffffff1663c09ad0d9855f015160400151838686601490809261396493929190615eb4565b6040518563ffffffff1660e01b815260040161398293929190616c9f565b5f604051808303815f88803b158015613999575f5ffd5b5087f11580156139ab573d5f5f3e3d5ffd5b50505050505b50613ca9565b5f8173ffffffffffffffffffffffffffffffffffffffff163b14613a1257846040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613a099190616d19565b60405180910390fd5b6014838390501015613a5b57846040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613a529190616d8f565b60405180910390fd5b5f613a64610c8a565b73ffffffffffffffffffffffffffffffffffffffff1663570e1a36865f01516040015186866040518463ffffffff1660e01b8152600401613aa6929190615577565b6020604051808303815f8887f1158015613ac2573d5f5f3e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190613ae791906155ad565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613b5957856040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613b509190616e05565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613bc957856040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613bc09190616e7b565b60405180910390fd5b5f8173ffffffffffffffffffffffffffffffffffffffff163b03613c2457856040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613c1b9190616ef1565b60405180910390fd5b5f84845f90601492613c3893929190615eb4565b90613c439190616bb8565b60601c90508273ffffffffffffffffffffffffffffffffffffffff1686602001517fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d83895f015160e00151604051613c9c929190616f1d565b60405180910390a35050505b5b50505050565b5f5f835f01516040015190505f845f01515f015190505f5f613ccf611df6565b90505f88886020015188604051602401613ceb93929190616f44565b6040516020818303038152906040526319822f7c60e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060205f8251602084015f888af192505f51955060203d14613d51575f92505b613d5a82611e37565b505080613e05575f8273ffffffffffffffffffffffffffffffffffffffff163b03613dbc57876040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613db39190616fca565b60405180910390fd5b87613dc8610800611dff565b6040517f65c8fd4d000000000000000000000000000000000000000000000000000000008152600401613dfc929190617040565b60405180910390fd5b505050949350505050565b5f5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f815f0154905083811015613e69575f92505050613e7a565b838103825f01819055506001925050505b92915050565b60605f5f613e8c611df6565b90505f8585602001518660400151604051602401613eac93929190616f44565b6040516020818303038152906040526352b7512c60e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f855f015160e0015190505f865f01516080015190505f5f5f5f5f5f5f895160208b015f8b8bf194503d9050805f8a3e6020890151995088519250606081039150604089019a508a519350841580613f4c575060408314155b80613f59575081601f8501105b15613fa7578d613f6a610800611dff565b6040517f65c8fd4d000000000000000000000000000000000000000000000000000000008152600401613f9e9291906170cb565b60405180910390fd5b613fb189826140bc565b505050505050505050935093915050565b613fca61416e565b5f8290505f60a084901c90505f8165ffffffffffff1603613fef5765ffffffffffff90505b5f60d085901c905060405180606001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020018265ffffffffffff1681526020018365ffffffffffff168152509350505050919050565b5f5f60ff835f1c169050601f811115614087576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80915050919050565b5f6080825f1c901c9050919050565b5f815f1c6fffffffffffffffffffffffffffffffff169050919050565b6140c681836140cb565b6140ef565b6140d4826140df565b810180604052505050565b5f601f19601f8301169050919050565b5050565b6040518060a001604052805f81526020015f151581526020015f6dffffffffffffffffffffffffffff1681526020015f63ffffffff1681526020015f65ffffffffffff1681525090565b6040518060a001604052806141506141b2565b81526020015f81526020015f81526020015f81526020015f81525090565b60405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f65ffffffffffff1681526020015f65ffffffffffff1681525090565b6040518061014001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61428682614240565b810181811067ffffffffffffffff821117156142a5576142a4614250565b5b80604052505050565b5f6142b7614227565b90506142c3828261427d565b919050565b5f67ffffffffffffffff8211156142e2576142e1614250565b5b6142eb82614240565b9050602081019050919050565b828183375f83830152505050565b5f614318614313846142c8565b6142ae565b9050828152602081018484840111156143345761433361423c565b5b61433f8482856142f8565b509392505050565b5f82601f83011261435b5761435a614238565b5b813561436b848260208601614306565b91505092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6143a182614378565b9050919050565b6143b181614397565b81146143bb575f5ffd5b50565b5f813590506143cc816143a8565b92915050565b5f819050919050565b6143e4816143d2565b81146143ee575f5ffd5b50565b5f813590506143ff816143db565b92915050565b5f610140828403121561441b5761441a614374565b5b6144266101406142ae565b90505f614435848285016143be565b5f830152506020614448848285016143f1565b602083015250604061445c848285016143f1565b6040830152506060614470848285016143f1565b6060830152506080614484848285016143f1565b60808301525060a0614498848285016143f1565b60a08301525060c06144ac848285016143f1565b60c08301525060e06144c0848285016143be565b60e0830152506101006144d5848285016143f1565b610100830152506101206144eb848285016143f1565b6101208301525092915050565b5f819050919050565b61450a816144f8565b8114614514575f5ffd5b50565b5f8135905061452581614501565b92915050565b5f6101c0828403121561454157614540614374565b5b61454b60a06142ae565b90505f61455a84828501614405565b5f8301525061014061456e84828501614517565b602083015250610160614583848285016143f1565b604083015250610180614598848285016143f1565b6060830152506101a06145ad848285016143f1565b60808301525092915050565b5f5ffd5b5f5ffd5b5f5f83601f8401126145d6576145d5614238565b5b8235905067ffffffffffffffff8111156145f3576145f26145b9565b5b60208301915083600182028301111561460f5761460e6145bd565b5b9250929050565b5f5f5f5f610200858703121561462f5761462e614230565b5b5f85013567ffffffffffffffff81111561464c5761464b614234565b5b61465887828801614347565b94505060206146698782880161452b565b9350506101e085013567ffffffffffffffff81111561468b5761468a614234565b5b614697878288016145c1565b925092505092959194509250565b6146ae816143d2565b82525050565b5f6020820190506146c75f8301846146a5565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b614701816146cd565b811461470b575f5ffd5b50565b5f8135905061471c816146f8565b92915050565b5f6020828403121561473757614736614230565b5b5f6147448482850161470e565b91505092915050565b5f8115159050919050565b6147618161474d565b82525050565b5f60208201905061477a5f830184614758565b92915050565b5f63ffffffff82169050919050565b61479881614780565b81146147a2575f5ffd5b50565b5f813590506147b38161478f565b92915050565b5f602082840312156147ce576147cd614230565b5b5f6147db848285016147a5565b91505092915050565b5f819050919050565b5f6148076148026147fd84614378565b6147e4565b614378565b9050919050565b5f614818826147ed565b9050919050565b5f6148298261480e565b9050919050565b6148398161481f565b82525050565b5f6020820190506148525f830184614830565b92915050565b5f77ffffffffffffffffffffffffffffffffffffffffffffffff82169050919050565b61488481614858565b811461488e575f5ffd5b50565b5f8135905061489f8161487b565b92915050565b5f602082840312156148ba576148b9614230565b5b5f6148c784828501614891565b91505092915050565b6148d9816144f8565b82525050565b5f6020820190506148f25f8301846148d0565b92915050565b5f5f6040838503121561490e5761490d614230565b5b5f61491b858286016143be565b925050602061492c85828601614891565b9150509250929050565b5f61494082614378565b9050919050565b61495081614936565b811461495a575f5ffd5b50565b5f8135905061496b81614947565b92915050565b5f5f6040838503121561498757614986614230565b5b5f6149948582860161495d565b92505060206149a5858286016143f1565b9150509250929050565b5f5ffd5b5f61012082840312156149c9576149c86149af565b5b81905092915050565b5f602082840312156149e7576149e6614230565b5b5f82013567ffffffffffffffff811115614a0457614a03614234565b5b614a10848285016149b3565b91505092915050565b5f60208284031215614a2e57614a2d614230565b5b5f614a3b848285016143be565b91505092915050565b614a4d816143d2565b82525050565b614a5c8161474d565b82525050565b5f6dffffffffffffffffffffffffffff82169050919050565b614a8481614a62565b82525050565b614a9381614780565b82525050565b5f65ffffffffffff82169050919050565b614ab381614a99565b82525050565b60a082015f820151614acd5f850182614a44565b506020820151614ae06020850182614a53565b506040820151614af36040850182614a7b565b506060820151614b066060850182614a8a565b506080820151614b196080850182614aaa565b50505050565b5f60a082019050614b325f830184614ab9565b92915050565b5f5f83601f840112614b4d57614b4c614238565b5b8235905067ffffffffffffffff811115614b6a57614b696145b9565b5b602083019150836020820283011115614b8657614b856145bd565b5b9250929050565b5f5f5f60408486031215614ba457614ba3614230565b5b5f84013567ffffffffffffffff811115614bc157614bc0614234565b5b614bcd86828701614b38565b93509350506020614be08682870161495d565b9150509250925092565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b614c1e81614bea565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f614c5682614c24565b614c608185614c2e565b9350614c70818560208601614c3e565b614c7981614240565b840191505092915050565b614c8d81614397565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f614cc78383614a44565b60208301905092915050565b5f602082019050919050565b5f614ce982614c93565b614cf38185614c9d565b9350614cfe83614cad565b805f5b83811015614d2e578151614d158882614cbc565b9750614d2083614cd3565b925050600181019050614d01565b5085935050505092915050565b5f60e082019050614d4e5f83018a614c15565b8181036020830152614d608189614c4c565b90508181036040830152614d748188614c4c565b9050614d8360608301876146a5565b614d906080830186614c84565b614d9d60a08301856148d0565b81810360c0830152614daf8184614cdf565b905098975050505050505050565b5f5f5f60408486031215614dd457614dd3614230565b5b5f614de1868287016143be565b935050602084013567ffffffffffffffff811115614e0257614e01614234565b5b614e0e868287016145c1565b92509250509250925092565b5f5f60208385031215614e3057614e2f614230565b5b5f83013567ffffffffffffffff811115614e4d57614e4c614234565b5b614e59858286016145c1565b92509250509250929050565b5f60208284031215614e7a57614e79614230565b5b5f614e878482850161495d565b91505092915050565b5f5f83601f840112614ea557614ea4614238565b5b8235905067ffffffffffffffff811115614ec257614ec16145b9565b5b602083019150836020820283011115614ede57614edd6145bd565b5b9250929050565b5f5f5f60408486031215614efc57614efb614230565b5b5f84013567ffffffffffffffff811115614f1957614f18614234565b5b614f2586828701614e90565b93509350506020614f388682870161495d565b9150509250925092565b7f4141393220696e7465726e616c2063616c6c206f6e6c790000000000000000005f82015250565b5f614f76601783614c2e565b9150614f8182614f42565b602082019050919050565b5f6020820190508181035f830152614fa381614f6a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f614ffb82614fd7565b6150058185614fe1565b9350615015818560208601614c3e565b61501e81614240565b840191505092915050565b5f60408201905061503c5f8301856146a5565b818103602083015261504e8184614ff1565b90509392505050565b7f6d757374207370656369667920756e7374616b652064656c61790000000000005f82015250565b5f61508b601a83614c2e565b915061509682615057565b602082019050919050565b5f6020820190508181035f8301526150b88161507f565b9050919050565b7f63616e6e6f7420646563726561736520756e7374616b652074696d65000000005f82015250565b5f6150f3601c83614c2e565b91506150fe826150bf565b602082019050919050565b5f6020820190508181035f830152615120816150e7565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61515e826143d2565b9150615169836143d2565b925082820190508082111561518157615180615127565b5b92915050565b7f6e6f207374616b652073706563696669656400000000000000000000000000005f82015250565b5f6151bb601283614c2e565b91506151c682615187565b602082019050919050565b5f6020820190508181035f8301526151e8816151af565b9050919050565b7f7374616b65206f766572666c6f770000000000000000000000000000000000005f82015250565b5f615223600e83614c2e565b915061522e826151ef565b602082019050919050565b5f6020820190508181035f83015261525081615217565b9050919050565b5f61527161526c61526784614780565b6147e4565b6143d2565b9050919050565b61528181615257565b82525050565b5f60408201905061529a5f8301856146a5565b6152a76020830184615278565b9392505050565b5f6152b8826143d2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036152ea576152e9615127565b5b600182019050919050565b7f576974686472617720616d6f756e7420746f6f206c61726765000000000000005f82015250565b5f615329601983614c2e565b9150615334826152f5565b602082019050919050565b5f6020820190508181035f8301526153568161531d565b9050919050565b5f615367826143d2565b9150615372836143d2565b925082820390508181111561538a57615389615127565b5b92915050565b5f61539a8261480e565b9050919050565b6153aa81615390565b82525050565b5f6040820190506153c35f8301856153a1565b6153d060208301846146a5565b9392505050565b5f81905092915050565b50565b5f6153ef5f836153d7565b91506153fa826153e1565b5f82019050919050565b5f61540e826153e4565b9150819050919050565b7f6661696c656420746f20776974686472617700000000000000000000000000005f82015250565b5f61544c601283614c2e565b915061545782615418565b602082019050919050565b5f6020820190508181035f83015261547981615440565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f82356001610120038336030381126154d5576154d46154ad565b5b80830191505092915050565b5f6154ec83856153d7565b93506154f98385846142f8565b82840190509392505050565b5f6155118284866154e1565b91508190509392505050565b5f6040820190506155305f830185614758565b81810360208301526155428184614ff1565b90509392505050565b5f6155568385614fe1565b93506155638385846142f8565b61556c83614240565b840190509392505050565b5f6020820190508181035f83015261559081848661554b565b90509392505050565b5f815190506155a7816143a8565b92915050565b5f602082840312156155c2576155c1614230565b5b5f6155cf84828501615599565b91505092915050565b5f6020820190506155eb5f830184614c84565b92915050565b7f6e6f74207374616b6564000000000000000000000000000000000000000000005f82015250565b5f615625600a83614c2e565b9150615630826155f1565b602082019050919050565b5f6020820190508181035f83015261565281615619565b9050919050565b7f616c726561647920756e7374616b696e670000000000000000000000000000005f82015250565b5f61568d601183614c2e565b915061569882615659565b602082019050919050565b5f6020820190508181035f8301526156ba81615681565b9050919050565b5f6156cb82614a99565b91506156d683614a99565b9250828201905065ffffffffffff8111156156f4576156f3615127565b5b92915050565b5f61571461570f61570a84614a99565b6147e4565b6143d2565b9050919050565b615724816156fa565b82525050565b5f60208201905061573d5f83018461571b565b92915050565b7f4e6f207374616b6520746f2077697468647261770000000000000000000000005f82015250565b5f615777601483614c2e565b915061578282615743565b602082019050919050565b5f6020820190508181035f8301526157a48161576b565b9050919050565b7f6d7573742063616c6c20756e6c6f636b5374616b6528292066697273740000005f82015250565b5f6157df601d83614c2e565b91506157ea826157ab565b602082019050919050565b5f6020820190508181035f83015261580c816157d3565b9050919050565b7f5374616b65207769746864726177616c206973206e6f742064756500000000005f82015250565b5f615847601b83614c2e565b915061585282615813565b602082019050919050565b5f6020820190508181035f8301526158748161583b565b9050919050565b7f6661696c656420746f207769746864726177207374616b6500000000000000005f82015250565b5f6158af601883614c2e565b91506158ba8261587b565b602082019050919050565b5f6020820190508181035f8301526158dc816158a3565b9050919050565b5f823560016060038336030381126158fe576158fd6154ad565b5b80830191505092915050565b5f5f83356001602003843603038112615926576159256154ad565b5b80840192508235915067ffffffffffffffff821115615948576159476154b1565b5b602083019250602082023603831315615964576159636154b5565b5b509250929050565b5f61597682614397565b9050919050565b6159868161596c565b8114615990575f5ffd5b50565b5f813590506159a18161597d565b92915050565b5f602082840312156159bc576159bb614230565b5b5f6159c984828501615993565b91505092915050565b5f5f833560016020038436030381126159ee576159ed6154ad565b5b80840192508235915067ffffffffffffffff821115615a1057615a0f6154b1565b5b602083019250600182023603831315615a2c57615a2b6154b5565b5b509250929050565b5f82825260208201905092915050565b5f819050919050565b5f615a5b60208401846143be565b905092915050565b615a6c81614397565b82525050565b5f615a8060208401846143f1565b905092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83356001602003843603038112615ab057615aaf615a90565b5b83810192508235915060208301925067ffffffffffffffff821115615ad857615ad7615a88565b5b600182023603831315615aee57615aed615a8c565b5b509250929050565b5f82825260208201905092915050565b5f615b118385615af6565b9350615b1e8385846142f8565b615b2783614240565b840190509392505050565b5f615b406020840184614517565b905092915050565b615b51816144f8565b82525050565b5f6101208301615b695f840184615a4d565b615b755f860182615a63565b50615b836020840184615a72565b615b906020860182614a44565b50615b9e6040840184615a94565b8583036040870152615bb1838284615b06565b92505050615bc26060840184615a94565b8583036060870152615bd5838284615b06565b92505050615be66080840184615b32565b615bf36080860182615b48565b50615c0160a0840184615a72565b615c0e60a0860182614a44565b50615c1c60c0840184615b32565b615c2960c0860182615b48565b50615c3760e0840184615a94565b85830360e0870152615c4a838284615b06565b92505050615c5c610100840184615a94565b858303610100870152615c70838284615b06565b925050508091505092915050565b5f615c898383615b57565b905092915050565b5f8235600161012003833603038112615cad57615cac615a90565b5b82810191505092915050565b5f602082019050919050565b5f615cd08385615a34565b935083602084028501615ce284615a44565b805f5b87811015615d25578484038952615cfc8284615c91565b615d068582615c7e565b9450615d1183615cb9565b925060208a01995050600181019050615ce5565b50829750879450505050509392505050565b5f6040820190508181035f830152615d50818688615cc5565b90508181036020830152615d6581848661554b565b905095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60038110615dae57615dad615d70565b5b50565b5f819050615dbe82615d9d565b919050565b5f615dcd82615db1565b9050919050565b615ddd81615dc3565b82525050565b5f608082019050615df65f830187615dd4565b8181036020830152615e088186614ff1565b9050615e1760408301856146a5565b615e2460608301846146a5565b95945050505050565b5f6020820190508181035f830152615e458184614ff1565b905092915050565b5f8160601b9050919050565b5f615e6382615e4d565b9050919050565b5f615e7482615e59565b9050919050565b615e8c615e8782614397565b615e6a565b82525050565b5f615e9d8284615e7b565b60148201915081905092915050565b5f5ffd5b5f5ffd5b5f5f85851115615ec757615ec6615eac565b5b83861115615ed857615ed7615eb0565b5b6001850283019150848603905094509492505050565b5f615ef98286615e7b565b601482019150615f0a8284866154e1565b9150819050949350505050565b5f6101208301615f295f840184615a4d565b615f355f860182615a63565b50615f436020840184615a72565b615f506020860182614a44565b50615f5e6040840184615a94565b8583036040870152615f71838284615b06565b92505050615f826060840184615a94565b8583036060870152615f95838284615b06565b92505050615fa66080840184615b32565b615fb36080860182615b48565b50615fc160a0840184615a72565b615fce60a0860182614a44565b50615fdc60c0840184615b32565b615fe960c0860182615b48565b50615ff760e0840184615a94565b85830360e087015261600a838284615b06565b9250505061601c610100840184615a94565b858303610100870152616030838284615b06565b925050508091505092915050565b5f6040820190508181035f8301526160568185615f17565b905061606560208301846148d0565b9392505050565b61014082015f8201516160815f850182615a63565b5060208201516160946020850182614a44565b5060408201516160a76040850182614a44565b5060608201516160ba6060850182614a44565b5060808201516160cd6080850182614a44565b5060a08201516160e060a0850182614a44565b5060c08201516160f360c0850182614a44565b5060e082015161610660e0850182615a63565b5061010082015161611b610100850182614a44565b50610120820151616130610120850182614a44565b50505050565b6101c082015f82015161614b5f85018261606c565b50602082015161615f610140850182615b48565b506040820151616173610160850182614a44565b506060820151616187610180850182614a44565b50608082015161619b6101a0850182614a44565b50505050565b5f610200820190508181035f8301526161ba8186614ff1565b90506161c96020830185616136565b8181036101e08301526161dc8184614ff1565b9050949350505050565b5f610200820190508181035f83015261620081868861554b565b905061620f6020830185616136565b8181036101e08301526162228184614ff1565b905095945050505050565b7f41413935206f7574206f662067617300000000000000000000000000000000005f82015250565b5f616261600f83614c2e565b915061626c8261622d565b602082019050919050565b5f60408201905061628a5f8301846146a5565b818103602083015261629b81616255565b905092915050565b7f4141393020696e76616c69642062656e656669636961727900000000000000005f82015250565b5f6162d7601883614c2e565b91506162e2826162a3565b602082019050919050565b5f6020820190508181035f830152616304816162cb565b9050919050565b7f41413931206661696c65642073656e6420746f2062656e6566696369617279005f82015250565b5f61633f601f83614c2e565b915061634a8261630b565b602082019050919050565b5f6020820190508181035f83015261636c81616333565b9050919050565b5f6080820190506163865f8301876146a5565b6163936020830186614758565b6163a060408301856146a5565b6163ad60608301846146a5565b95945050505050565b5f60a0820190506163c95f8301886148d0565b6163d660208301876148d0565b6163e360408301866148d0565b6163f060608301856146a5565b6163fd6080830184614c84565b9695505050505050565b7f73656e64657220686173206e6f20636f646500000000000000000000000000005f82015250565b5f61643b601283614c2e565b915061644682616407565b602082019050919050565b5f6020820190508181035f8301526164688161642f565b9050919050565b7f6e6f7420616e204549502d373730322064656c656761746500000000000000005f82015250565b5f6164a3601883614c2e565b91506164ae8261646f565b602082019050919050565b5f6020820190508181035f8301526164d081616497565b9050919050565b5f610120820190506164eb5f83018c6148d0565b6164f8602083018b614c84565b616505604083018a6146a5565b61651260608301896148d0565b61651f60808301886148d0565b61652c60a08301876148d0565b61653960c08301866146a5565b61654660e08301856148d0565b6165546101008301846148d0565b9a9950505050505050505050565b7f41413934206761732076616c756573206f766572666c6f7700000000000000005f82015250565b5f616596601883614c2e565b91506165a182616562565b602082019050919050565b5f6040820190506165bf5f8301846146a5565b81810360208301526165d08161658a565b905092915050565b7f4141323520696e76616c6964206163636f756e74206e6f6e63650000000000005f82015250565b5f61660c601a83614c2e565b9150616617826165d8565b602082019050919050565b5f6040820190506166355f8301846146a5565b818103602083015261664681616600565b905092915050565b7f41413236206f76657220766572696669636174696f6e4761734c696d697400005f82015250565b5f616682601e83614c2e565b915061668d8261664e565b602082019050919050565b5f6040820190506166ab5f8301846146a5565b81810360208301526166bc81616676565b905092915050565b7f41413234207369676e6174757265206572726f720000000000000000000000005f82015250565b5f6166f8601483614c2e565b9150616703826166c4565b602082019050919050565b5f6040820190506167215f8301846146a5565b8181036020830152616732816166ec565b905092915050565b7f414132322065787069726564206f72206e6f74206475650000000000000000005f82015250565b5f61676e601783614c2e565b91506167798261673a565b602082019050919050565b5f6040820190506167975f8301846146a5565b81810360208301526167a881616762565b905092915050565b7f41413334207369676e6174757265206572726f720000000000000000000000005f82015250565b5f6167e4601483614c2e565b91506167ef826167b0565b602082019050919050565b5f60408201905061680d5f8301846146a5565b818103602083015261681e816167d8565b905092915050565b7f41413332207061796d61737465722065787069726564206f72206e6f742064755f8201527f6500000000000000000000000000000000000000000000000000000000000000602082015250565b5f616880602183614c2e565b915061688b82616826565b604082019050919050565b5f6040820190506168a95f8301846146a5565b81810360208301526168ba81616874565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061690657607f821691505b602082108103616919576169186168c2565b5b50919050565b7f4141393320696e76616c6964207061796d6173746572416e64446174610000005f82015250565b5f616953601d83614c2e565b915061695e8261691f565b602082019050919050565b5f6020820190508181035f83015261698081616947565b9050919050565b7f4141393820696e76616c6964207061796d6173746572000000000000000000005f82015250565b5f6169bb601683614c2e565b91506169c682616987565b602082019050919050565b5f6020820190508181035f8301526169e8816169af565b9050919050565b7f41413231206469646e2774207061792070726566756e640000000000000000005f82015250565b5f616a23601783614c2e565b9150616a2e826169ef565b602082019050919050565b5f604082019050616a4c5f8301846146a5565b8181036020830152616a5d81616a17565b905092915050565b7f41413331207061796d6173746572206465706f73697420746f6f206c6f7700005f82015250565b5f616a99601e83614c2e565b9150616aa482616a65565b602082019050919050565b5f604082019050616ac25f8301846146a5565b8181036020830152616ad381616a8d565b905092915050565b7f41413336206f766572207061796d6173746572566572696669636174696f6e475f8201527f61734c696d697400000000000000000000000000000000000000000000000000602082015250565b5f616b35602783614c2e565b9150616b4082616adb565b604082019050919050565b5f604082019050616b5e5f8301846146a5565b8181036020830152616b6f81616b29565b905092915050565b5f82905092915050565b5f7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b5f82821b905092915050565b5f616bc38383616b77565b82616bce8135616b81565b92506014821015616c0e57616c097fffffffffffffffffffffffffffffffffffffffff00000000000000000000000083601403600802616bac565b831692505b505092915050565b5f7fffffffffffffffffffffffffffffffff0000000000000000000000000000000082169050919050565b5f616c4c8383616b77565b82616c578135616c16565b92506010821015616c9757616c927fffffffffffffffffffffffffffffffff0000000000000000000000000000000083601003600802616bac565b831692505b505092915050565b5f604082019050616cb25f830186614c84565b8181036020830152616cc581848661554b565b9050949350505050565b7f414131302073656e64657220616c726561647920636f6e7374727563746564005f82015250565b5f616d03601f83614c2e565b9150616d0e82616ccf565b602082019050919050565b5f604082019050616d2c5f8301846146a5565b8181036020830152616d3d81616cf7565b905092915050565b7f4141393920696e6974436f646520746f6f20736d616c6c0000000000000000005f82015250565b5f616d79601783614c2e565b9150616d8482616d45565b602082019050919050565b5f604082019050616da25f8301846146a5565b8181036020830152616db381616d6d565b905092915050565b7f4141313320696e6974436f6465206661696c6564206f72204f4f4700000000005f82015250565b5f616def601b83614c2e565b9150616dfa82616dbb565b602082019050919050565b5f604082019050616e185f8301846146a5565b8181036020830152616e2981616de3565b905092915050565b7f4141313420696e6974436f6465206d7573742072657475726e2073656e6465725f82015250565b5f616e65602083614c2e565b9150616e7082616e31565b602082019050919050565b5f604082019050616e8e5f8301846146a5565b8181036020830152616e9f81616e59565b905092915050565b7f4141313520696e6974436f6465206d757374206372656174652073656e6465725f82015250565b5f616edb602083614c2e565b9150616ee682616ea7565b602082019050919050565b5f604082019050616f045f8301846146a5565b8181036020830152616f1581616ecf565b905092915050565b5f604082019050616f305f830185614c84565b616f3d6020830184614c84565b9392505050565b5f6060820190508181035f830152616f5c8186615f17565b9050616f6b60208301856148d0565b616f7860408301846146a5565b949350505050565b7f41413230206163636f756e74206e6f74206465706c6f796564000000000000005f82015250565b5f616fb4601983614c2e565b9150616fbf82616f80565b602082019050919050565b5f604082019050616fdd5f8301846146a5565b8181036020830152616fee81616fa8565b905092915050565b7f41413233207265766572746564000000000000000000000000000000000000005f82015250565b5f61702a600d83614c2e565b915061703582616ff6565b602082019050919050565b5f6060820190506170535f8301856146a5565b81810360208301526170648161701e565b905081810360408301526170788184614ff1565b90509392505050565b7f41413333207265766572746564000000000000000000000000000000000000005f82015250565b5f6170b5600d83614c2e565b91506170c082617081565b602082019050919050565b5f6060820190506170de5f8301856146a5565b81810360208301526170ef816170a9565b905081810360408301526171038184614ff1565b9050939250505056fea26469706673582212208dd9f8a1804405d0d9b665d95ecffa5a98d8860f659ae34c73f2ee2c1a07ce5564736f6c634300081d003360a0604052348015600e575f5ffd5b503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250506080516108c26100685f395f818160b0015281816101de015261020201526108c25ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063570e1a3614610043578063b0d691fe14610073578063c09ad0d914610091575b5f5ffd5b61005d6004803603810190610058919061039e565b6100ad565b60405161006a9190610428565b60405180910390f35b61007b6101dc565b6040516100889190610428565b60405180910390f35b6100ab60048036038101906100a691906105a3565b610200565b005b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461013c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013390610657565b60405180910390fd5b5f83835f906014926101509392919061067d565b9061015b91906106f8565b60601c90505f848460149080926101749392919061067d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090505f60205f8351602085015f875af1905080156101d3575f5193505b50505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461028e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028590610657565b60405180910390fd5b5f5f5f8351602085015f875af19050806102ef575f6102ae6108006102f4565b90505f816040517f65c8fd4d0000000000000000000000000000000000000000000000000000000081526004016102e692919061084b565b60405180910390fd5b505050565b60603d5f83111561030c578281111561030b578290505b5b604051602082018101604052818152815f602083013e8092505050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261035e5761035d61033d565b5b8235905067ffffffffffffffff81111561037b5761037a610341565b5b60208301915083600182028301111561039757610396610345565b5b9250929050565b5f5f602083850312156103b4576103b3610335565b5b5f83013567ffffffffffffffff8111156103d1576103d0610339565b5b6103dd85828601610349565b92509250509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610412826103e9565b9050919050565b61042281610408565b82525050565b5f60208201905061043b5f830184610419565b92915050565b61044a81610408565b8114610454575f5ffd5b50565b5f8135905061046581610441565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6104b58261046f565b810181811067ffffffffffffffff821117156104d4576104d361047f565b5b80604052505050565b5f6104e661032c565b90506104f282826104ac565b919050565b5f67ffffffffffffffff8211156105115761051061047f565b5b61051a8261046f565b9050602081019050919050565b828183375f83830152505050565b5f610547610542846104f7565b6104dd565b9050828152602081018484840111156105635761056261046b565b5b61056e848285610527565b509392505050565b5f82601f83011261058a5761058961033d565b5b813561059a848260208601610535565b91505092915050565b5f5f604083850312156105b9576105b8610335565b5b5f6105c685828601610457565b925050602083013567ffffffffffffffff8111156105e7576105e6610339565b5b6105f385828601610576565b9150509250929050565b5f82825260208201905092915050565b7f414139372073686f756c642063616c6c2066726f6d20456e747279506f696e745f82015250565b5f6106416020836105fd565b915061064c8261060d565b602082019050919050565b5f6020820190508181035f83015261066e81610635565b9050919050565b5f5ffd5b5f5ffd5b5f5f858511156106905761068f610675565b5b838611156106a1576106a0610679565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b5f82821b905092915050565b5f61070383836106b7565b8261070e81356106c1565b9250601482101561074e576107497fffffffffffffffffffffffffffffffffffffffff000000000000000000000000836014036008026106ec565b831692505b505092915050565b5f819050919050565b5f819050919050565b5f819050919050565b5f61078b61078661078184610756565b610768565b61075f565b9050919050565b61079b81610771565b82525050565b7f4141313320454950373730322073656e64657220696e6974206661696c6564005f82015250565b5f6107d5601f836105fd565b91506107e0826107a1565b602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61081d826107eb565b61082781856107f5565b9350610837818560208601610805565b6108408161046f565b840191505092915050565b5f60608201905061085e5f830185610792565b818103602083015261086f816107c9565b905081810360408301526108838184610813565b9050939250505056fea26469706673582212200747fe7ca42bac0b159be2c8536919a8e04cd78351433270e46514c67063e45064736f6c634300081d0033",
}

// EntryPoint is an auto generated Go binding around an Ethereum contract.
type EntryPoint struct {
	abi abi.ABI
}

// NewEntryPoint creates a new instance of EntryPoint.
func NewEntryPoint() *EntryPoint {
	parsed, err := EntryPointMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &EntryPoint{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *EntryPoint) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (entryPoint *EntryPoint) PackAddStake(unstakeDelaySec uint32) []byte {
	enc, err := entryPoint.abi.Pack("addStake", unstakeDelaySec)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (entryPoint *EntryPoint) PackBalanceOf(account common.Address) []byte {
	enc, err := entryPoint.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (entryPoint *EntryPoint) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := entryPoint.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDelegateAndRevert is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (entryPoint *EntryPoint) PackDelegateAndRevert(target common.Address, data []byte) []byte {
	enc, err := entryPoint.abi.Pack("delegateAndRevert", target, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDepositTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (entryPoint *EntryPoint) PackDepositTo(account common.Address) []byte {
	enc, err := entryPoint.abi.Pack("depositTo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (entryPoint *EntryPoint) PackEip712Domain() []byte {
	enc, err := entryPoint.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (entryPoint *EntryPoint) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := entryPoint.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, err

}

// PackGetDepositInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (entryPoint *EntryPoint) PackGetDepositInfo(account common.Address) []byte {
	enc, err := entryPoint.abi.Pack("getDepositInfo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDepositInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (entryPoint *EntryPoint) UnpackGetDepositInfo(data []byte) (IStakeManagerDepositInfo, error) {
	out, err := entryPoint.abi.Unpack("getDepositInfo", data)
	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}
	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)
	return out0, err
}

// PackGetDomainSeparatorV4 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13c65a6e.
//
// Solidity: function getDomainSeparatorV4() view returns(bytes32)
func (entryPoint *EntryPoint) PackGetDomainSeparatorV4() []byte {
	enc, err := entryPoint.abi.Pack("getDomainSeparatorV4")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDomainSeparatorV4 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x13c65a6e.
//
// Solidity: function getDomainSeparatorV4() view returns(bytes32)
func (entryPoint *EntryPoint) UnpackGetDomainSeparatorV4(data []byte) ([32]byte, error) {
	out, err := entryPoint.abi.Unpack("getDomainSeparatorV4", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGetNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (entryPoint *EntryPoint) PackGetNonce(sender common.Address, key *big.Int) []byte {
	enc, err := entryPoint.abi.Pack("getNonce", sender, key)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (entryPoint *EntryPoint) UnpackGetNonce(data []byte) (*big.Int, error) {
	out, err := entryPoint.abi.Unpack("getNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetPackedUserOpTypeHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x154e58dc.
//
// Solidity: function getPackedUserOpTypeHash() pure returns(bytes32)
func (entryPoint *EntryPoint) PackGetPackedUserOpTypeHash() []byte {
	enc, err := entryPoint.abi.Pack("getPackedUserOpTypeHash")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPackedUserOpTypeHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x154e58dc.
//
// Solidity: function getPackedUserOpTypeHash() pure returns(bytes32)
func (entryPoint *EntryPoint) UnpackGetPackedUserOpTypeHash(data []byte) ([32]byte, error) {
	out, err := entryPoint.abi.Unpack("getPackedUserOpTypeHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGetSenderAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (entryPoint *EntryPoint) PackGetSenderAddress(initCode []byte) []byte {
	enc, err := entryPoint.abi.Pack("getSenderAddress", initCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetUserOpHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (entryPoint *EntryPoint) PackGetUserOpHash(userOp PackedUserOperation) []byte {
	enc, err := entryPoint.abi.Pack("getUserOpHash", userOp)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUserOpHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (entryPoint *EntryPoint) UnpackGetUserOpHash(data []byte) ([32]byte, error) {
	out, err := entryPoint.abi.Unpack("getUserOpHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackHandleAggregatedOps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (entryPoint *EntryPoint) PackHandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) []byte {
	enc, err := entryPoint.abi.Pack("handleAggregatedOps", opsPerAggregator, beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHandleOps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (entryPoint *EntryPoint) PackHandleOps(ops []PackedUserOperation, beneficiary common.Address) []byte {
	enc, err := entryPoint.abi.Pack("handleOps", ops, beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIncrementNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (entryPoint *EntryPoint) PackIncrementNonce(key *big.Int) []byte {
	enc, err := entryPoint.abi.Pack("incrementNonce", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackInnerHandleOp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0042dc53.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (entryPoint *EntryPoint) PackInnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) []byte {
	enc, err := entryPoint.abi.Pack("innerHandleOp", callData, opInfo, context)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackInnerHandleOp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0042dc53.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (entryPoint *EntryPoint) UnpackInnerHandleOp(data []byte) (*big.Int, error) {
	out, err := entryPoint.abi.Unpack("innerHandleOp", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackNonceSequenceNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (entryPoint *EntryPoint) PackNonceSequenceNumber(arg0 common.Address, arg1 *big.Int) []byte {
	enc, err := entryPoint.abi.Pack("nonceSequenceNumber", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonceSequenceNumber is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (entryPoint *EntryPoint) UnpackNonceSequenceNumber(data []byte) (*big.Int, error) {
	out, err := entryPoint.abi.Unpack("nonceSequenceNumber", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackSenderCreator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (entryPoint *EntryPoint) PackSenderCreator() []byte {
	enc, err := entryPoint.abi.Pack("senderCreator")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSenderCreator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (entryPoint *EntryPoint) UnpackSenderCreator(data []byte) (common.Address, error) {
	out, err := entryPoint.abi.Unpack("senderCreator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (entryPoint *EntryPoint) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := entryPoint.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (entryPoint *EntryPoint) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := entryPoint.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUnlockStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (entryPoint *EntryPoint) PackUnlockStake() []byte {
	enc, err := entryPoint.abi.Pack("unlockStake")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (entryPoint *EntryPoint) PackWithdrawStake(withdrawAddress common.Address) []byte {
	enc, err := entryPoint.abi.Pack("withdrawStake", withdrawAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (entryPoint *EntryPoint) PackWithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) []byte {
	enc, err := entryPoint.abi.Pack("withdrawTo", withdrawAddress, withdrawAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// EntryPointAccountDeployed represents a AccountDeployed event raised by the EntryPoint contract.
type EntryPointAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const EntryPointAccountDeployedEventName = "AccountDeployed"

// ContractEventName returns the user-defined event name.
func (EntryPointAccountDeployed) ContractEventName() string {
	return EntryPointAccountDeployedEventName
}

// UnpackAccountDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (entryPoint *EntryPoint) UnpackAccountDeployedEvent(log *types.Log) (*EntryPointAccountDeployed, error) {
	event := "AccountDeployed"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointAccountDeployed)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointBeforeExecution represents a BeforeExecution event raised by the EntryPoint contract.
type EntryPointBeforeExecution struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const EntryPointBeforeExecutionEventName = "BeforeExecution"

// ContractEventName returns the user-defined event name.
func (EntryPointBeforeExecution) ContractEventName() string {
	return EntryPointBeforeExecutionEventName
}

// UnpackBeforeExecutionEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BeforeExecution()
func (entryPoint *EntryPoint) UnpackBeforeExecutionEvent(log *types.Log) (*EntryPointBeforeExecution, error) {
	event := "BeforeExecution"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointBeforeExecution)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointDeposited represents a Deposited event raised by the EntryPoint contract.
type EntryPointDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const EntryPointDepositedEventName = "Deposited"

// ContractEventName returns the user-defined event name.
func (EntryPointDeposited) ContractEventName() string {
	return EntryPointDepositedEventName
}

// UnpackDepositedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (entryPoint *EntryPoint) UnpackDepositedEvent(log *types.Log) (*EntryPointDeposited, error) {
	event := "Deposited"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointDeposited)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointEIP712DomainChanged represents a EIP712DomainChanged event raised by the EntryPoint contract.
type EntryPointEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const EntryPointEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (EntryPointEIP712DomainChanged) ContractEventName() string {
	return EntryPointEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (entryPoint *EntryPoint) UnpackEIP712DomainChangedEvent(log *types.Log) (*EntryPointEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointPostOpRevertReason represents a PostOpRevertReason event raised by the EntryPoint contract.
type EntryPointPostOpRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const EntryPointPostOpRevertReasonEventName = "PostOpRevertReason"

// ContractEventName returns the user-defined event name.
func (EntryPointPostOpRevertReason) ContractEventName() string {
	return EntryPointPostOpRevertReasonEventName
}

// UnpackPostOpRevertReasonEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (entryPoint *EntryPoint) UnpackPostOpRevertReasonEvent(log *types.Log) (*EntryPointPostOpRevertReason, error) {
	event := "PostOpRevertReason"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointPostOpRevertReason)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointSignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the EntryPoint contract.
type EntryPointSignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const EntryPointSignatureAggregatorChangedEventName = "SignatureAggregatorChanged"

// ContractEventName returns the user-defined event name.
func (EntryPointSignatureAggregatorChanged) ContractEventName() string {
	return EntryPointSignatureAggregatorChangedEventName
}

// UnpackSignatureAggregatorChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (entryPoint *EntryPoint) UnpackSignatureAggregatorChangedEvent(log *types.Log) (*EntryPointSignatureAggregatorChanged, error) {
	event := "SignatureAggregatorChanged"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointSignatureAggregatorChanged)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointStakeLocked represents a StakeLocked event raised by the EntryPoint contract.
type EntryPointStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const EntryPointStakeLockedEventName = "StakeLocked"

// ContractEventName returns the user-defined event name.
func (EntryPointStakeLocked) ContractEventName() string {
	return EntryPointStakeLockedEventName
}

// UnpackStakeLockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (entryPoint *EntryPoint) UnpackStakeLockedEvent(log *types.Log) (*EntryPointStakeLocked, error) {
	event := "StakeLocked"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointStakeLocked)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointStakeUnlocked represents a StakeUnlocked event raised by the EntryPoint contract.
type EntryPointStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const EntryPointStakeUnlockedEventName = "StakeUnlocked"

// ContractEventName returns the user-defined event name.
func (EntryPointStakeUnlocked) ContractEventName() string {
	return EntryPointStakeUnlockedEventName
}

// UnpackStakeUnlockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (entryPoint *EntryPoint) UnpackStakeUnlockedEvent(log *types.Log) (*EntryPointStakeUnlocked, error) {
	event := "StakeUnlocked"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointStakeUnlocked)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointStakeWithdrawn represents a StakeWithdrawn event raised by the EntryPoint contract.
type EntryPointStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const EntryPointStakeWithdrawnEventName = "StakeWithdrawn"

// ContractEventName returns the user-defined event name.
func (EntryPointStakeWithdrawn) ContractEventName() string {
	return EntryPointStakeWithdrawnEventName
}

// UnpackStakeWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (entryPoint *EntryPoint) UnpackStakeWithdrawnEvent(log *types.Log) (*EntryPointStakeWithdrawn, error) {
	event := "StakeWithdrawn"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointStakeWithdrawn)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointUserOperationEvent represents a UserOperationEvent event raised by the EntryPoint contract.
type EntryPointUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Paymaster     common.Address
	Nonce         *big.Int
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const EntryPointUserOperationEventEventName = "UserOperationEvent"

// ContractEventName returns the user-defined event name.
func (EntryPointUserOperationEvent) ContractEventName() string {
	return EntryPointUserOperationEventEventName
}

// UnpackUserOperationEventEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (entryPoint *EntryPoint) UnpackUserOperationEventEvent(log *types.Log) (*EntryPointUserOperationEvent, error) {
	event := "UserOperationEvent"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointUserOperationEvent)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointUserOperationPrefundTooLow represents a UserOperationPrefundTooLow event raised by the EntryPoint contract.
type EntryPointUserOperationPrefundTooLow struct {
	UserOpHash [32]byte
	Sender     common.Address
	Nonce      *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const EntryPointUserOperationPrefundTooLowEventName = "UserOperationPrefundTooLow"

// ContractEventName returns the user-defined event name.
func (EntryPointUserOperationPrefundTooLow) ContractEventName() string {
	return EntryPointUserOperationPrefundTooLowEventName
}

// UnpackUserOperationPrefundTooLowEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
func (entryPoint *EntryPoint) UnpackUserOperationPrefundTooLowEvent(log *types.Log) (*EntryPointUserOperationPrefundTooLow, error) {
	event := "UserOperationPrefundTooLow"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointUserOperationPrefundTooLow)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointUserOperationRevertReason represents a UserOperationRevertReason event raised by the EntryPoint contract.
type EntryPointUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const EntryPointUserOperationRevertReasonEventName = "UserOperationRevertReason"

// ContractEventName returns the user-defined event name.
func (EntryPointUserOperationRevertReason) ContractEventName() string {
	return EntryPointUserOperationRevertReasonEventName
}

// UnpackUserOperationRevertReasonEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (entryPoint *EntryPoint) UnpackUserOperationRevertReasonEvent(log *types.Log) (*EntryPointUserOperationRevertReason, error) {
	event := "UserOperationRevertReason"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointUserOperationRevertReason)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// EntryPointWithdrawn represents a Withdrawn event raised by the EntryPoint contract.
type EntryPointWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const EntryPointWithdrawnEventName = "Withdrawn"

// ContractEventName returns the user-defined event name.
func (EntryPointWithdrawn) ContractEventName() string {
	return EntryPointWithdrawnEventName
}

// UnpackWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (entryPoint *EntryPoint) UnpackWithdrawnEvent(log *types.Log) (*EntryPointWithdrawn, error) {
	event := "Withdrawn"
	if log.Topics[0] != entryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(EntryPointWithdrawn)
	if len(log.Data) > 0 {
		if err := entryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range entryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (entryPoint *EntryPoint) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], entryPoint.abi.Errors["DelegateAndRevert"].ID.Bytes()[:4]) {
		return entryPoint.UnpackDelegateAndRevertError(raw[4:])
	}
	if bytes.Equal(raw[:4], entryPoint.abi.Errors["FailedOp"].ID.Bytes()[:4]) {
		return entryPoint.UnpackFailedOpError(raw[4:])
	}
	if bytes.Equal(raw[:4], entryPoint.abi.Errors["FailedOpWithRevert"].ID.Bytes()[:4]) {
		return entryPoint.UnpackFailedOpWithRevertError(raw[4:])
	}
	if bytes.Equal(raw[:4], entryPoint.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return entryPoint.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], entryPoint.abi.Errors["PostOpReverted"].ID.Bytes()[:4]) {
		return entryPoint.UnpackPostOpRevertedError(raw[4:])
	}
	if bytes.Equal(raw[:4], entryPoint.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return entryPoint.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], entryPoint.abi.Errors["SenderAddressResult"].ID.Bytes()[:4]) {
		return entryPoint.UnpackSenderAddressResultError(raw[4:])
	}
	if bytes.Equal(raw[:4], entryPoint.abi.Errors["SignatureValidationFailed"].ID.Bytes()[:4]) {
		return entryPoint.UnpackSignatureValidationFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], entryPoint.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return entryPoint.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// EntryPointDelegateAndRevert represents a DelegateAndRevert error raised by the EntryPoint contract.
type EntryPointDelegateAndRevert struct {
	Success bool
	Ret     []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DelegateAndRevert(bool success, bytes ret)
func EntryPointDelegateAndRevertErrorID() common.Hash {
	return common.HexToHash("0x9941055444a6b8379a4c66beac4880d5e96ca9c2fff188903cb0bc945a4dae05")
}

// UnpackDelegateAndRevertError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DelegateAndRevert(bool success, bytes ret)
func (entryPoint *EntryPoint) UnpackDelegateAndRevertError(raw []byte) (*EntryPointDelegateAndRevert, error) {
	out := new(EntryPointDelegateAndRevert)
	if err := entryPoint.abi.UnpackIntoInterface(out, "DelegateAndRevert", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// EntryPointFailedOp represents a FailedOp error raised by the EntryPoint contract.
type EntryPointFailedOp struct {
	OpIndex *big.Int
	Reason  string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedOp(uint256 opIndex, string reason)
func EntryPointFailedOpErrorID() common.Hash {
	return common.HexToHash("0x220266b6eadfd2488e398d00abc1c765e1f119da6226c1b55ec9cc0186560475")
}

// UnpackFailedOpError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedOp(uint256 opIndex, string reason)
func (entryPoint *EntryPoint) UnpackFailedOpError(raw []byte) (*EntryPointFailedOp, error) {
	out := new(EntryPointFailedOp)
	if err := entryPoint.abi.UnpackIntoInterface(out, "FailedOp", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// EntryPointFailedOpWithRevert represents a FailedOpWithRevert error raised by the EntryPoint contract.
type EntryPointFailedOpWithRevert struct {
	OpIndex *big.Int
	Reason  string
	Inner   []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedOpWithRevert(uint256 opIndex, string reason, bytes inner)
func EntryPointFailedOpWithRevertErrorID() common.Hash {
	return common.HexToHash("0x65c8fd4dd32f2bb83f7d57728e1afa231e9956aaf4bdeea76c78b967426acb8c")
}

// UnpackFailedOpWithRevertError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedOpWithRevert(uint256 opIndex, string reason, bytes inner)
func (entryPoint *EntryPoint) UnpackFailedOpWithRevertError(raw []byte) (*EntryPointFailedOpWithRevert, error) {
	out := new(EntryPointFailedOpWithRevert)
	if err := entryPoint.abi.UnpackIntoInterface(out, "FailedOpWithRevert", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// EntryPointInvalidShortString represents a InvalidShortString error raised by the EntryPoint contract.
type EntryPointInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func EntryPointInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (entryPoint *EntryPoint) UnpackInvalidShortStringError(raw []byte) (*EntryPointInvalidShortString, error) {
	out := new(EntryPointInvalidShortString)
	if err := entryPoint.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// EntryPointPostOpReverted represents a PostOpReverted error raised by the EntryPoint contract.
type EntryPointPostOpReverted struct {
	ReturnData []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PostOpReverted(bytes returnData)
func EntryPointPostOpRevertedErrorID() common.Hash {
	return common.HexToHash("0xad7954bcae0d69c56f2323097b645bc6b6033c9f3777375271d77a608fc206ae")
}

// UnpackPostOpRevertedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PostOpReverted(bytes returnData)
func (entryPoint *EntryPoint) UnpackPostOpRevertedError(raw []byte) (*EntryPointPostOpReverted, error) {
	out := new(EntryPointPostOpReverted)
	if err := entryPoint.abi.UnpackIntoInterface(out, "PostOpReverted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// EntryPointReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the EntryPoint contract.
type EntryPointReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func EntryPointReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (entryPoint *EntryPoint) UnpackReentrancyGuardReentrantCallError(raw []byte) (*EntryPointReentrancyGuardReentrantCall, error) {
	out := new(EntryPointReentrancyGuardReentrantCall)
	if err := entryPoint.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// EntryPointSenderAddressResult represents a SenderAddressResult error raised by the EntryPoint contract.
type EntryPointSenderAddressResult struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SenderAddressResult(address sender)
func EntryPointSenderAddressResultErrorID() common.Hash {
	return common.HexToHash("0x6ca7b806055bb56d141e3a30604d4c6540fe6f010574d1b40dfd951da77b956d")
}

// UnpackSenderAddressResultError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SenderAddressResult(address sender)
func (entryPoint *EntryPoint) UnpackSenderAddressResultError(raw []byte) (*EntryPointSenderAddressResult, error) {
	out := new(EntryPointSenderAddressResult)
	if err := entryPoint.abi.UnpackIntoInterface(out, "SenderAddressResult", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// EntryPointSignatureValidationFailed represents a SignatureValidationFailed error raised by the EntryPoint contract.
type EntryPointSignatureValidationFailed struct {
	Aggregator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SignatureValidationFailed(address aggregator)
func EntryPointSignatureValidationFailedErrorID() common.Hash {
	return common.HexToHash("0x86a9f750e187f100e9e8a18befbbf44e6f512d6f78cfe16a061541aaaad52523")
}

// UnpackSignatureValidationFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SignatureValidationFailed(address aggregator)
func (entryPoint *EntryPoint) UnpackSignatureValidationFailedError(raw []byte) (*EntryPointSignatureValidationFailed, error) {
	out := new(EntryPointSignatureValidationFailed)
	if err := entryPoint.abi.UnpackIntoInterface(out, "SignatureValidationFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// EntryPointStringTooLong represents a StringTooLong error raised by the EntryPoint contract.
type EntryPointStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func EntryPointStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (entryPoint *EntryPoint) UnpackStringTooLongError(raw []byte) (*EntryPointStringTooLong, error) {
	out := new(EntryPointStringTooLong)
	if err := entryPoint.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExecMetaData contains all meta data concerning the Exec contract.
var ExecMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "e3a6df279b6ae020a6f0a001d353e8e726",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea26469706673582212206e043986f7fa91ad18e042ff8859cbec29795cc4256c5a516362a38d19202b6764736f6c634300081d0033",
}

// Exec is an auto generated Go binding around an Ethereum contract.
type Exec struct {
	abi abi.ABI
}

// NewExec creates a new instance of Exec.
func NewExec() *Exec {
	parsed, err := ExecMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Exec{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Exec) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// IAccountMetaData contains all meta data concerning the IAccount contract.
var IAccountMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "4919412b74a737289a0329422d575bba9d",
}

// IAccount is an auto generated Go binding around an Ethereum contract.
type IAccount struct {
	abi abi.ABI
}

// NewIAccount creates a new instance of IAccount.
func NewIAccount() *IAccount {
	parsed, err := IAccountMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IAccount{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IAccount) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackValidateUserOp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (iAccount *IAccount) PackValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) []byte {
	enc, err := iAccount.abi.Pack("validateUserOp", userOp, userOpHash, missingAccountFunds)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidateUserOp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (iAccount *IAccount) UnpackValidateUserOp(data []byte) (*big.Int, error) {
	out, err := iAccount.abi.Unpack("validateUserOp", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// IAccountExecuteMetaData contains all meta data concerning the IAccountExecute contract.
var IAccountExecuteMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"executeUserOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "fd98bf629bbcb4b004234638daf0d0812a",
}

// IAccountExecute is an auto generated Go binding around an Ethereum contract.
type IAccountExecute struct {
	abi abi.ABI
}

// NewIAccountExecute creates a new instance of IAccountExecute.
func NewIAccountExecute() *IAccountExecute {
	parsed, err := IAccountExecuteMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IAccountExecute{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IAccountExecute) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackExecuteUserOp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8dd7712f.
//
// Solidity: function executeUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash) returns()
func (iAccountExecute *IAccountExecute) PackExecuteUserOp(userOp PackedUserOperation, userOpHash [32]byte) []byte {
	enc, err := iAccountExecute.abi.Pack("executeUserOp", userOp, userOpHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// IAggregatorMetaData contains all meta data concerning the IAggregator contract.
var IAggregatorMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"}],\"name\":\"aggregateSignatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"aggregatedSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"validateUserOpSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"sigForUserOp\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "feda3ff01de5efb7f8606d469b5ebdc3b3",
}

// IAggregator is an auto generated Go binding around an Ethereum contract.
type IAggregator struct {
	abi abi.ABI
}

// NewIAggregator creates a new instance of IAggregator.
func NewIAggregator() *IAggregator {
	parsed, err := IAggregatorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IAggregator{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IAggregator) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAggregateSignatures is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) view returns(bytes aggregatedSignature)
func (iAggregator *IAggregator) PackAggregateSignatures(userOps []PackedUserOperation) []byte {
	enc, err := iAggregator.abi.Pack("aggregateSignatures", userOps)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAggregateSignatures is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) view returns(bytes aggregatedSignature)
func (iAggregator *IAggregator) UnpackAggregateSignatures(data []byte) ([]byte, error) {
	out, err := iAggregator.abi.Unpack("aggregateSignatures", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, err
}

// PackValidateSignatures is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) returns()
func (iAggregator *IAggregator) PackValidateSignatures(userOps []PackedUserOperation, signature []byte) []byte {
	enc, err := iAggregator.abi.Pack("validateSignatures", userOps, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackValidateUserOpSignature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes sigForUserOp)
func (iAggregator *IAggregator) PackValidateUserOpSignature(userOp PackedUserOperation) []byte {
	enc, err := iAggregator.abi.Pack("validateUserOpSignature", userOp)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidateUserOpSignature is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes sigForUserOp)
func (iAggregator *IAggregator) UnpackValidateUserOpSignature(data []byte) ([]byte, error) {
	out, err := iAggregator.abi.Unpack("validateUserOpSignature", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, err
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "a33a1eaf2949eb84428aeece5515d1fc34",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	abi abi.ABI
}

// NewIERC165 creates a new instance of IERC165.
func NewIERC165() *IERC165 {
	parsed, err := IERC165MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC165{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC165) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC165 *IERC165) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := iERC165.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC165 *IERC165) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := iERC165.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// IERC5267MetaData contains all meta data concerning the IERC5267 contract.
var IERC5267MetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "1c4527d1ae1005321690e330bf30b1045c",
}

// IERC5267 is an auto generated Go binding around an Ethereum contract.
type IERC5267 struct {
	abi abi.ABI
}

// NewIERC5267 creates a new instance of IERC5267.
func NewIERC5267() *IERC5267 {
	parsed, err := IERC5267MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC5267{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC5267) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (iERC5267 *IERC5267) PackEip712Domain() []byte {
	enc, err := iERC5267.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (iERC5267 *IERC5267) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := iERC5267.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, err

}

// IERC5267EIP712DomainChanged represents a EIP712DomainChanged event raised by the IERC5267 contract.
type IERC5267EIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const IERC5267EIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (IERC5267EIP712DomainChanged) ContractEventName() string {
	return IERC5267EIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (iERC5267 *IERC5267) UnpackEIP712DomainChangedEvent(log *types.Log) (*IERC5267EIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != iERC5267.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC5267EIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := iERC5267.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC5267.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointMetaData contains all meta data concerning the IEntryPoint contract.
var IEntryPointMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"name\":\"DelegateAndRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"PostOpReverted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureValidationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureAggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"getSenderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAggregator\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.UserOpsPerAggregator[]\",\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleAggregatedOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"senderCreator\",\"outputs\":[{\"internalType\":\"contractISenderCreator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "a49a45610987d37b31917b146a38b15008",
}

// IEntryPoint is an auto generated Go binding around an Ethereum contract.
type IEntryPoint struct {
	abi abi.ABI
}

// NewIEntryPoint creates a new instance of IEntryPoint.
func NewIEntryPoint() *IEntryPoint {
	parsed, err := IEntryPointMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IEntryPoint{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IEntryPoint) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (iEntryPoint *IEntryPoint) PackAddStake(unstakeDelaySec uint32) []byte {
	enc, err := iEntryPoint.abi.Pack("addStake", unstakeDelaySec)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iEntryPoint *IEntryPoint) PackBalanceOf(account common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iEntryPoint *IEntryPoint) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iEntryPoint.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDelegateAndRevert is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (iEntryPoint *IEntryPoint) PackDelegateAndRevert(target common.Address, data []byte) []byte {
	enc, err := iEntryPoint.abi.Pack("delegateAndRevert", target, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDepositTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (iEntryPoint *IEntryPoint) PackDepositTo(account common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("depositTo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetDepositInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (iEntryPoint *IEntryPoint) PackGetDepositInfo(account common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("getDepositInfo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDepositInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (iEntryPoint *IEntryPoint) UnpackGetDepositInfo(data []byte) (IStakeManagerDepositInfo, error) {
	out, err := iEntryPoint.abi.Unpack("getDepositInfo", data)
	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}
	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)
	return out0, err
}

// PackGetNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (iEntryPoint *IEntryPoint) PackGetNonce(sender common.Address, key *big.Int) []byte {
	enc, err := iEntryPoint.abi.Pack("getNonce", sender, key)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (iEntryPoint *IEntryPoint) UnpackGetNonce(data []byte) (*big.Int, error) {
	out, err := iEntryPoint.abi.Unpack("getNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetSenderAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (iEntryPoint *IEntryPoint) PackGetSenderAddress(initCode []byte) []byte {
	enc, err := iEntryPoint.abi.Pack("getSenderAddress", initCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetUserOpHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (iEntryPoint *IEntryPoint) PackGetUserOpHash(userOp PackedUserOperation) []byte {
	enc, err := iEntryPoint.abi.Pack("getUserOpHash", userOp)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUserOpHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (iEntryPoint *IEntryPoint) UnpackGetUserOpHash(data []byte) ([32]byte, error) {
	out, err := iEntryPoint.abi.Unpack("getUserOpHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackHandleAggregatedOps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (iEntryPoint *IEntryPoint) PackHandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("handleAggregatedOps", opsPerAggregator, beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHandleOps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (iEntryPoint *IEntryPoint) PackHandleOps(ops []PackedUserOperation, beneficiary common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("handleOps", ops, beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIncrementNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (iEntryPoint *IEntryPoint) PackIncrementNonce(key *big.Int) []byte {
	enc, err := iEntryPoint.abi.Pack("incrementNonce", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSenderCreator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (iEntryPoint *IEntryPoint) PackSenderCreator() []byte {
	enc, err := iEntryPoint.abi.Pack("senderCreator")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSenderCreator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (iEntryPoint *IEntryPoint) UnpackSenderCreator(data []byte) (common.Address, error) {
	out, err := iEntryPoint.abi.Unpack("senderCreator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUnlockStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (iEntryPoint *IEntryPoint) PackUnlockStake() []byte {
	enc, err := iEntryPoint.abi.Pack("unlockStake")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (iEntryPoint *IEntryPoint) PackWithdrawStake(withdrawAddress common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("withdrawStake", withdrawAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (iEntryPoint *IEntryPoint) PackWithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) []byte {
	enc, err := iEntryPoint.abi.Pack("withdrawTo", withdrawAddress, withdrawAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// IEntryPointAccountDeployed represents a AccountDeployed event raised by the IEntryPoint contract.
type IEntryPointAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const IEntryPointAccountDeployedEventName = "AccountDeployed"

// ContractEventName returns the user-defined event name.
func (IEntryPointAccountDeployed) ContractEventName() string {
	return IEntryPointAccountDeployedEventName
}

// UnpackAccountDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (iEntryPoint *IEntryPoint) UnpackAccountDeployedEvent(log *types.Log) (*IEntryPointAccountDeployed, error) {
	event := "AccountDeployed"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointAccountDeployed)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointBeforeExecution represents a BeforeExecution event raised by the IEntryPoint contract.
type IEntryPointBeforeExecution struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const IEntryPointBeforeExecutionEventName = "BeforeExecution"

// ContractEventName returns the user-defined event name.
func (IEntryPointBeforeExecution) ContractEventName() string {
	return IEntryPointBeforeExecutionEventName
}

// UnpackBeforeExecutionEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BeforeExecution()
func (iEntryPoint *IEntryPoint) UnpackBeforeExecutionEvent(log *types.Log) (*IEntryPointBeforeExecution, error) {
	event := "BeforeExecution"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointBeforeExecution)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointDeposited represents a Deposited event raised by the IEntryPoint contract.
type IEntryPointDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IEntryPointDepositedEventName = "Deposited"

// ContractEventName returns the user-defined event name.
func (IEntryPointDeposited) ContractEventName() string {
	return IEntryPointDepositedEventName
}

// UnpackDepositedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (iEntryPoint *IEntryPoint) UnpackDepositedEvent(log *types.Log) (*IEntryPointDeposited, error) {
	event := "Deposited"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointDeposited)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointPostOpRevertReason represents a PostOpRevertReason event raised by the IEntryPoint contract.
type IEntryPointPostOpRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const IEntryPointPostOpRevertReasonEventName = "PostOpRevertReason"

// ContractEventName returns the user-defined event name.
func (IEntryPointPostOpRevertReason) ContractEventName() string {
	return IEntryPointPostOpRevertReasonEventName
}

// UnpackPostOpRevertReasonEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (iEntryPoint *IEntryPoint) UnpackPostOpRevertReasonEvent(log *types.Log) (*IEntryPointPostOpRevertReason, error) {
	event := "PostOpRevertReason"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointPostOpRevertReason)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointSignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the IEntryPoint contract.
type IEntryPointSignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const IEntryPointSignatureAggregatorChangedEventName = "SignatureAggregatorChanged"

// ContractEventName returns the user-defined event name.
func (IEntryPointSignatureAggregatorChanged) ContractEventName() string {
	return IEntryPointSignatureAggregatorChangedEventName
}

// UnpackSignatureAggregatorChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (iEntryPoint *IEntryPoint) UnpackSignatureAggregatorChangedEvent(log *types.Log) (*IEntryPointSignatureAggregatorChanged, error) {
	event := "SignatureAggregatorChanged"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointSignatureAggregatorChanged)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointStakeLocked represents a StakeLocked event raised by the IEntryPoint contract.
type IEntryPointStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const IEntryPointStakeLockedEventName = "StakeLocked"

// ContractEventName returns the user-defined event name.
func (IEntryPointStakeLocked) ContractEventName() string {
	return IEntryPointStakeLockedEventName
}

// UnpackStakeLockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (iEntryPoint *IEntryPoint) UnpackStakeLockedEvent(log *types.Log) (*IEntryPointStakeLocked, error) {
	event := "StakeLocked"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointStakeLocked)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointStakeUnlocked represents a StakeUnlocked event raised by the IEntryPoint contract.
type IEntryPointStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IEntryPointStakeUnlockedEventName = "StakeUnlocked"

// ContractEventName returns the user-defined event name.
func (IEntryPointStakeUnlocked) ContractEventName() string {
	return IEntryPointStakeUnlockedEventName
}

// UnpackStakeUnlockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (iEntryPoint *IEntryPoint) UnpackStakeUnlockedEvent(log *types.Log) (*IEntryPointStakeUnlocked, error) {
	event := "StakeUnlocked"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointStakeUnlocked)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointStakeWithdrawn represents a StakeWithdrawn event raised by the IEntryPoint contract.
type IEntryPointStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const IEntryPointStakeWithdrawnEventName = "StakeWithdrawn"

// ContractEventName returns the user-defined event name.
func (IEntryPointStakeWithdrawn) ContractEventName() string {
	return IEntryPointStakeWithdrawnEventName
}

// UnpackStakeWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (iEntryPoint *IEntryPoint) UnpackStakeWithdrawnEvent(log *types.Log) (*IEntryPointStakeWithdrawn, error) {
	event := "StakeWithdrawn"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointStakeWithdrawn)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointUserOperationEvent represents a UserOperationEvent event raised by the IEntryPoint contract.
type IEntryPointUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Paymaster     common.Address
	Nonce         *big.Int
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const IEntryPointUserOperationEventEventName = "UserOperationEvent"

// ContractEventName returns the user-defined event name.
func (IEntryPointUserOperationEvent) ContractEventName() string {
	return IEntryPointUserOperationEventEventName
}

// UnpackUserOperationEventEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (iEntryPoint *IEntryPoint) UnpackUserOperationEventEvent(log *types.Log) (*IEntryPointUserOperationEvent, error) {
	event := "UserOperationEvent"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointUserOperationEvent)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointUserOperationPrefundTooLow represents a UserOperationPrefundTooLow event raised by the IEntryPoint contract.
type IEntryPointUserOperationPrefundTooLow struct {
	UserOpHash [32]byte
	Sender     common.Address
	Nonce      *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const IEntryPointUserOperationPrefundTooLowEventName = "UserOperationPrefundTooLow"

// ContractEventName returns the user-defined event name.
func (IEntryPointUserOperationPrefundTooLow) ContractEventName() string {
	return IEntryPointUserOperationPrefundTooLowEventName
}

// UnpackUserOperationPrefundTooLowEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
func (iEntryPoint *IEntryPoint) UnpackUserOperationPrefundTooLowEvent(log *types.Log) (*IEntryPointUserOperationPrefundTooLow, error) {
	event := "UserOperationPrefundTooLow"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointUserOperationPrefundTooLow)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointUserOperationRevertReason represents a UserOperationRevertReason event raised by the IEntryPoint contract.
type IEntryPointUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const IEntryPointUserOperationRevertReasonEventName = "UserOperationRevertReason"

// ContractEventName returns the user-defined event name.
func (IEntryPointUserOperationRevertReason) ContractEventName() string {
	return IEntryPointUserOperationRevertReasonEventName
}

// UnpackUserOperationRevertReasonEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (iEntryPoint *IEntryPoint) UnpackUserOperationRevertReasonEvent(log *types.Log) (*IEntryPointUserOperationRevertReason, error) {
	event := "UserOperationRevertReason"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointUserOperationRevertReason)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IEntryPointWithdrawn represents a Withdrawn event raised by the IEntryPoint contract.
type IEntryPointWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const IEntryPointWithdrawnEventName = "Withdrawn"

// ContractEventName returns the user-defined event name.
func (IEntryPointWithdrawn) ContractEventName() string {
	return IEntryPointWithdrawnEventName
}

// UnpackWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (iEntryPoint *IEntryPoint) UnpackWithdrawnEvent(log *types.Log) (*IEntryPointWithdrawn, error) {
	event := "Withdrawn"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointWithdrawn)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (iEntryPoint *IEntryPoint) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["DelegateAndRevert"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackDelegateAndRevertError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["FailedOp"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackFailedOpError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["FailedOpWithRevert"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackFailedOpWithRevertError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["PostOpReverted"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackPostOpRevertedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["SenderAddressResult"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackSenderAddressResultError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["SignatureValidationFailed"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackSignatureValidationFailedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IEntryPointDelegateAndRevert represents a DelegateAndRevert error raised by the IEntryPoint contract.
type IEntryPointDelegateAndRevert struct {
	Success bool
	Ret     []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DelegateAndRevert(bool success, bytes ret)
func IEntryPointDelegateAndRevertErrorID() common.Hash {
	return common.HexToHash("0x9941055444a6b8379a4c66beac4880d5e96ca9c2fff188903cb0bc945a4dae05")
}

// UnpackDelegateAndRevertError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DelegateAndRevert(bool success, bytes ret)
func (iEntryPoint *IEntryPoint) UnpackDelegateAndRevertError(raw []byte) (*IEntryPointDelegateAndRevert, error) {
	out := new(IEntryPointDelegateAndRevert)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "DelegateAndRevert", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointFailedOp represents a FailedOp error raised by the IEntryPoint contract.
type IEntryPointFailedOp struct {
	OpIndex *big.Int
	Reason  string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedOp(uint256 opIndex, string reason)
func IEntryPointFailedOpErrorID() common.Hash {
	return common.HexToHash("0x220266b6eadfd2488e398d00abc1c765e1f119da6226c1b55ec9cc0186560475")
}

// UnpackFailedOpError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedOp(uint256 opIndex, string reason)
func (iEntryPoint *IEntryPoint) UnpackFailedOpError(raw []byte) (*IEntryPointFailedOp, error) {
	out := new(IEntryPointFailedOp)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "FailedOp", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointFailedOpWithRevert represents a FailedOpWithRevert error raised by the IEntryPoint contract.
type IEntryPointFailedOpWithRevert struct {
	OpIndex *big.Int
	Reason  string
	Inner   []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedOpWithRevert(uint256 opIndex, string reason, bytes inner)
func IEntryPointFailedOpWithRevertErrorID() common.Hash {
	return common.HexToHash("0x65c8fd4dd32f2bb83f7d57728e1afa231e9956aaf4bdeea76c78b967426acb8c")
}

// UnpackFailedOpWithRevertError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedOpWithRevert(uint256 opIndex, string reason, bytes inner)
func (iEntryPoint *IEntryPoint) UnpackFailedOpWithRevertError(raw []byte) (*IEntryPointFailedOpWithRevert, error) {
	out := new(IEntryPointFailedOpWithRevert)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "FailedOpWithRevert", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointPostOpReverted represents a PostOpReverted error raised by the IEntryPoint contract.
type IEntryPointPostOpReverted struct {
	ReturnData []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PostOpReverted(bytes returnData)
func IEntryPointPostOpRevertedErrorID() common.Hash {
	return common.HexToHash("0xad7954bcae0d69c56f2323097b645bc6b6033c9f3777375271d77a608fc206ae")
}

// UnpackPostOpRevertedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PostOpReverted(bytes returnData)
func (iEntryPoint *IEntryPoint) UnpackPostOpRevertedError(raw []byte) (*IEntryPointPostOpReverted, error) {
	out := new(IEntryPointPostOpReverted)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "PostOpReverted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointSenderAddressResult represents a SenderAddressResult error raised by the IEntryPoint contract.
type IEntryPointSenderAddressResult struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SenderAddressResult(address sender)
func IEntryPointSenderAddressResultErrorID() common.Hash {
	return common.HexToHash("0x6ca7b806055bb56d141e3a30604d4c6540fe6f010574d1b40dfd951da77b956d")
}

// UnpackSenderAddressResultError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SenderAddressResult(address sender)
func (iEntryPoint *IEntryPoint) UnpackSenderAddressResultError(raw []byte) (*IEntryPointSenderAddressResult, error) {
	out := new(IEntryPointSenderAddressResult)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "SenderAddressResult", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointSignatureValidationFailed represents a SignatureValidationFailed error raised by the IEntryPoint contract.
type IEntryPointSignatureValidationFailed struct {
	Aggregator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SignatureValidationFailed(address aggregator)
func IEntryPointSignatureValidationFailedErrorID() common.Hash {
	return common.HexToHash("0x86a9f750e187f100e9e8a18befbbf44e6f512d6f78cfe16a061541aaaad52523")
}

// UnpackSignatureValidationFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SignatureValidationFailed(address aggregator)
func (iEntryPoint *IEntryPoint) UnpackSignatureValidationFailedError(raw []byte) (*IEntryPointSignatureValidationFailed, error) {
	out := new(IEntryPointSignatureValidationFailed)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "SignatureValidationFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// INonceManagerMetaData contains all meta data concerning the INonceManager contract.
var INonceManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "e3008b70798da6c43fdb3277b1161e0416",
}

// INonceManager is an auto generated Go binding around an Ethereum contract.
type INonceManager struct {
	abi abi.ABI
}

// NewINonceManager creates a new instance of INonceManager.
func NewINonceManager() *INonceManager {
	parsed, err := INonceManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &INonceManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *INonceManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (iNonceManager *INonceManager) PackGetNonce(sender common.Address, key *big.Int) []byte {
	enc, err := iNonceManager.abi.Pack("getNonce", sender, key)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (iNonceManager *INonceManager) UnpackGetNonce(data []byte) (*big.Int, error) {
	out, err := iNonceManager.abi.Unpack("getNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackIncrementNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (iNonceManager *INonceManager) PackIncrementNonce(key *big.Int) []byte {
	enc, err := iNonceManager.abi.Pack("incrementNonce", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// IPaymasterMetaData contains all meta data concerning the IPaymaster contract.
var IPaymasterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "b9c7ab37ed63d9643642fc159b9af3d898",
}

// IPaymaster is an auto generated Go binding around an Ethereum contract.
type IPaymaster struct {
	abi abi.ABI
}

// NewIPaymaster creates a new instance of IPaymaster.
func NewIPaymaster() *IPaymaster {
	parsed, err := IPaymasterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IPaymaster{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IPaymaster) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackPostOp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (iPaymaster *IPaymaster) PackPostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) []byte {
	enc, err := iPaymaster.abi.Pack("postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackValidatePaymasterUserOp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (iPaymaster *IPaymaster) PackValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) []byte {
	enc, err := iPaymaster.abi.Pack("validatePaymasterUserOp", userOp, userOpHash, maxCost)
	if err != nil {
		panic(err)
	}
	return enc
}

// ValidatePaymasterUserOpOutput serves as a container for the return parameters of contract
// method ValidatePaymasterUserOp.
type ValidatePaymasterUserOpOutput struct {
	Context        []byte
	ValidationData *big.Int
}

// UnpackValidatePaymasterUserOp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (iPaymaster *IPaymaster) UnpackValidatePaymasterUserOp(data []byte) (ValidatePaymasterUserOpOutput, error) {
	out, err := iPaymaster.abi.Unpack("validatePaymasterUserOp", data)
	outstruct := new(ValidatePaymasterUserOpOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Context = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ValidationData = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// ISenderCreatorMetaData contains all meta data concerning the ISenderCreator contract.
var ISenderCreatorMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"createSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCallData\",\"type\":\"bytes\"}],\"name\":\"initEip7702Sender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "99526a7b678af7daba8b6fa87b7fa87570",
}

// ISenderCreator is an auto generated Go binding around an Ethereum contract.
type ISenderCreator struct {
	abi abi.ABI
}

// NewISenderCreator creates a new instance of ISenderCreator.
func NewISenderCreator() *ISenderCreator {
	parsed, err := ISenderCreatorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ISenderCreator{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ISenderCreator) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCreateSender is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x570e1a36.
//
// Solidity: function createSender(bytes initCode) returns(address sender)
func (iSenderCreator *ISenderCreator) PackCreateSender(initCode []byte) []byte {
	enc, err := iSenderCreator.abi.Pack("createSender", initCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateSender is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x570e1a36.
//
// Solidity: function createSender(bytes initCode) returns(address sender)
func (iSenderCreator *ISenderCreator) UnpackCreateSender(data []byte) (common.Address, error) {
	out, err := iSenderCreator.abi.Unpack("createSender", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInitEip7702Sender is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc09ad0d9.
//
// Solidity: function initEip7702Sender(address sender, bytes initCallData) returns()
func (iSenderCreator *ISenderCreator) PackInitEip7702Sender(sender common.Address, initCallData []byte) []byte {
	enc, err := iSenderCreator.abi.Pack("initEip7702Sender", sender, initCallData)
	if err != nil {
		panic(err)
	}
	return enc
}

// IStakeManagerMetaData contains all meta data concerning the IStakeManager contract.
var IStakeManagerMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "900f64e1227f6c0d757280ec50bf49d1ca",
}

// IStakeManager is an auto generated Go binding around an Ethereum contract.
type IStakeManager struct {
	abi abi.ABI
}

// NewIStakeManager creates a new instance of IStakeManager.
func NewIStakeManager() *IStakeManager {
	parsed, err := IStakeManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IStakeManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IStakeManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (iStakeManager *IStakeManager) PackAddStake(unstakeDelaySec uint32) []byte {
	enc, err := iStakeManager.abi.Pack("addStake", unstakeDelaySec)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iStakeManager *IStakeManager) PackBalanceOf(account common.Address) []byte {
	enc, err := iStakeManager.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iStakeManager *IStakeManager) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iStakeManager.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDepositTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (iStakeManager *IStakeManager) PackDepositTo(account common.Address) []byte {
	enc, err := iStakeManager.abi.Pack("depositTo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetDepositInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (iStakeManager *IStakeManager) PackGetDepositInfo(account common.Address) []byte {
	enc, err := iStakeManager.abi.Pack("getDepositInfo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDepositInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (iStakeManager *IStakeManager) UnpackGetDepositInfo(data []byte) (IStakeManagerDepositInfo, error) {
	out, err := iStakeManager.abi.Unpack("getDepositInfo", data)
	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}
	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)
	return out0, err
}

// PackUnlockStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (iStakeManager *IStakeManager) PackUnlockStake() []byte {
	enc, err := iStakeManager.abi.Pack("unlockStake")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (iStakeManager *IStakeManager) PackWithdrawStake(withdrawAddress common.Address) []byte {
	enc, err := iStakeManager.abi.Pack("withdrawStake", withdrawAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (iStakeManager *IStakeManager) PackWithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) []byte {
	enc, err := iStakeManager.abi.Pack("withdrawTo", withdrawAddress, withdrawAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// IStakeManagerDeposited represents a Deposited event raised by the IStakeManager contract.
type IStakeManagerDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IStakeManagerDepositedEventName = "Deposited"

// ContractEventName returns the user-defined event name.
func (IStakeManagerDeposited) ContractEventName() string {
	return IStakeManagerDepositedEventName
}

// UnpackDepositedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (iStakeManager *IStakeManager) UnpackDepositedEvent(log *types.Log) (*IStakeManagerDeposited, error) {
	event := "Deposited"
	if log.Topics[0] != iStakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IStakeManagerDeposited)
	if len(log.Data) > 0 {
		if err := iStakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iStakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IStakeManagerStakeLocked represents a StakeLocked event raised by the IStakeManager contract.
type IStakeManagerStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const IStakeManagerStakeLockedEventName = "StakeLocked"

// ContractEventName returns the user-defined event name.
func (IStakeManagerStakeLocked) ContractEventName() string {
	return IStakeManagerStakeLockedEventName
}

// UnpackStakeLockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (iStakeManager *IStakeManager) UnpackStakeLockedEvent(log *types.Log) (*IStakeManagerStakeLocked, error) {
	event := "StakeLocked"
	if log.Topics[0] != iStakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IStakeManagerStakeLocked)
	if len(log.Data) > 0 {
		if err := iStakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iStakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IStakeManagerStakeUnlocked represents a StakeUnlocked event raised by the IStakeManager contract.
type IStakeManagerStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IStakeManagerStakeUnlockedEventName = "StakeUnlocked"

// ContractEventName returns the user-defined event name.
func (IStakeManagerStakeUnlocked) ContractEventName() string {
	return IStakeManagerStakeUnlockedEventName
}

// UnpackStakeUnlockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (iStakeManager *IStakeManager) UnpackStakeUnlockedEvent(log *types.Log) (*IStakeManagerStakeUnlocked, error) {
	event := "StakeUnlocked"
	if log.Topics[0] != iStakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IStakeManagerStakeUnlocked)
	if len(log.Data) > 0 {
		if err := iStakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iStakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IStakeManagerStakeWithdrawn represents a StakeWithdrawn event raised by the IStakeManager contract.
type IStakeManagerStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const IStakeManagerStakeWithdrawnEventName = "StakeWithdrawn"

// ContractEventName returns the user-defined event name.
func (IStakeManagerStakeWithdrawn) ContractEventName() string {
	return IStakeManagerStakeWithdrawnEventName
}

// UnpackStakeWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (iStakeManager *IStakeManager) UnpackStakeWithdrawnEvent(log *types.Log) (*IStakeManagerStakeWithdrawn, error) {
	event := "StakeWithdrawn"
	if log.Topics[0] != iStakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IStakeManagerStakeWithdrawn)
	if len(log.Data) > 0 {
		if err := iStakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iStakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IStakeManagerWithdrawn represents a Withdrawn event raised by the IStakeManager contract.
type IStakeManagerWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const IStakeManagerWithdrawnEventName = "Withdrawn"

// ContractEventName returns the user-defined event name.
func (IStakeManagerWithdrawn) ContractEventName() string {
	return IStakeManagerWithdrawnEventName
}

// UnpackWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (iStakeManager *IStakeManager) UnpackWithdrawnEvent(log *types.Log) (*IStakeManagerWithdrawn, error) {
	event := "Withdrawn"
	if log.Topics[0] != iStakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IStakeManagerWithdrawn)
	if len(log.Data) > 0 {
		if err := iStakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iStakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "362817aca504bb11799e06f16c73252206",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220683c51e8ea97c4e4170220c794a0b9807eea6fedc6561993ea3ae8887338376e64736f6c634300081d0033",
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	abi abi.ABI
}

// NewMath creates a new instance of Math.
func NewMath() *Math {
	parsed, err := MathMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Math{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Math) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// MessageHashUtilsMetaData contains all meta data concerning the MessageHashUtils contract.
var MessageHashUtilsMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "2a46deb1e6ec75a21a12c9c84a8059f39d",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220690172a239af15dfbf2701213c92c03d95c473ee4551281a350b071600b4c77b64736f6c634300081d0033",
}

// MessageHashUtils is an auto generated Go binding around an Ethereum contract.
type MessageHashUtils struct {
	abi abi.ABI
}

// NewMessageHashUtils creates a new instance of MessageHashUtils.
func NewMessageHashUtils() *MessageHashUtils {
	parsed, err := MessageHashUtilsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MessageHashUtils{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MessageHashUtils) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// NonceManagerMetaData contains all meta data concerning the NonceManager contract.
var NonceManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"nonceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "9b166fcfd3720f8a1ce57382a45eae5f7b",
}

// NonceManager is an auto generated Go binding around an Ethereum contract.
type NonceManager struct {
	abi abi.ABI
}

// NewNonceManager creates a new instance of NonceManager.
func NewNonceManager() *NonceManager {
	parsed, err := NonceManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &NonceManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *NonceManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (nonceManager *NonceManager) PackGetNonce(sender common.Address, key *big.Int) []byte {
	enc, err := nonceManager.abi.Pack("getNonce", sender, key)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (nonceManager *NonceManager) UnpackGetNonce(data []byte) (*big.Int, error) {
	out, err := nonceManager.abi.Unpack("getNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackIncrementNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (nonceManager *NonceManager) PackIncrementNonce(key *big.Int) []byte {
	enc, err := nonceManager.abi.Pack("incrementNonce", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNonceSequenceNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (nonceManager *NonceManager) PackNonceSequenceNumber(arg0 common.Address, arg1 *big.Int) []byte {
	enc, err := nonceManager.abi.Pack("nonceSequenceNumber", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonceSequenceNumber is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (nonceManager *NonceManager) UnpackNonceSequenceNumber(data []byte) (*big.Int, error) {
	out, err := nonceManager.abi.Unpack("nonceSequenceNumber", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PanicMetaData contains all meta data concerning the Panic contract.
var PanicMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "d7d9ad298c70ed58e3eba1456b2d564fae",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220c9819e52b8d01db07ac18154a6cc103a795038d9beb2bf74d25bab0113d7face64736f6c634300081d0033",
}

// Panic is an auto generated Go binding around an Ethereum contract.
type Panic struct {
	abi abi.ABI
}

// NewPanic creates a new instance of Panic.
func NewPanic() *Panic {
	parsed, err := PanicMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Panic{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Panic) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// ReentrancyGuardTransientMetaData contains all meta data concerning the ReentrancyGuardTransient contract.
var ReentrancyGuardTransientMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"}]",
	ID:  "cdea1d85643a6387453325a8843f34a328",
}

// ReentrancyGuardTransient is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuardTransient struct {
	abi abi.ABI
}

// NewReentrancyGuardTransient creates a new instance of ReentrancyGuardTransient.
func NewReentrancyGuardTransient() *ReentrancyGuardTransient {
	parsed, err := ReentrancyGuardTransientMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ReentrancyGuardTransient{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ReentrancyGuardTransient) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (reentrancyGuardTransient *ReentrancyGuardTransient) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], reentrancyGuardTransient.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return reentrancyGuardTransient.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ReentrancyGuardTransientReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the ReentrancyGuardTransient contract.
type ReentrancyGuardTransientReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func ReentrancyGuardTransientReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (reentrancyGuardTransient *ReentrancyGuardTransient) UnpackReentrancyGuardReentrantCallError(raw []byte) (*ReentrancyGuardTransientReentrancyGuardReentrantCall, error) {
	out := new(ReentrancyGuardTransientReentrancyGuardReentrantCall)
	if err := reentrancyGuardTransient.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntToUint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintToInt\",\"type\":\"error\"}]",
	ID:  "ea30ab87e1a3c2e0ff3c21e804c0b5d55f",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea26469706673582212203f1a009fcc7702087d87122afad8bf35f6335eedff824b438081909ac18829ee64736f6c634300081d0033",
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	abi abi.ABI
}

// NewSafeCast creates a new instance of SafeCast.
func NewSafeCast() *SafeCast {
	parsed, err := SafeCastMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &SafeCast{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *SafeCast) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (safeCast *SafeCast) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], safeCast.abi.Errors["SafeCastOverflowedIntDowncast"].ID.Bytes()[:4]) {
		return safeCast.UnpackSafeCastOverflowedIntDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], safeCast.abi.Errors["SafeCastOverflowedIntToUint"].ID.Bytes()[:4]) {
		return safeCast.UnpackSafeCastOverflowedIntToUintError(raw[4:])
	}
	if bytes.Equal(raw[:4], safeCast.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return safeCast.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], safeCast.abi.Errors["SafeCastOverflowedUintToInt"].ID.Bytes()[:4]) {
		return safeCast.UnpackSafeCastOverflowedUintToIntError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// SafeCastSafeCastOverflowedIntDowncast represents a SafeCastOverflowedIntDowncast error raised by the SafeCast contract.
type SafeCastSafeCastOverflowedIntDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func SafeCastSafeCastOverflowedIntDowncastErrorID() common.Hash {
	return common.HexToHash("0x327269a7f29c3c5436f42eeed1c1adf0d4d525f36360483f4e83ab79e98f9089")
}

// UnpackSafeCastOverflowedIntDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func (safeCast *SafeCast) UnpackSafeCastOverflowedIntDowncastError(raw []byte) (*SafeCastSafeCastOverflowedIntDowncast, error) {
	out := new(SafeCastSafeCastOverflowedIntDowncast)
	if err := safeCast.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeCastSafeCastOverflowedIntToUint represents a SafeCastOverflowedIntToUint error raised by the SafeCast contract.
type SafeCastSafeCastOverflowedIntToUint struct {
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntToUint(int256 value)
func SafeCastSafeCastOverflowedIntToUintErrorID() common.Hash {
	return common.HexToHash("0xa8ce4432b175c373e5f41aba830358e5361584f628450fd436c066323ad91ac2")
}

// UnpackSafeCastOverflowedIntToUintError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntToUint(int256 value)
func (safeCast *SafeCast) UnpackSafeCastOverflowedIntToUintError(raw []byte) (*SafeCastSafeCastOverflowedIntToUint, error) {
	out := new(SafeCastSafeCastOverflowedIntToUint)
	if err := safeCast.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntToUint", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeCastSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the SafeCast contract.
type SafeCastSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func SafeCastSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (safeCast *SafeCast) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*SafeCastSafeCastOverflowedUintDowncast, error) {
	out := new(SafeCastSafeCastOverflowedUintDowncast)
	if err := safeCast.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeCastSafeCastOverflowedUintToInt represents a SafeCastOverflowedUintToInt error raised by the SafeCast contract.
type SafeCastSafeCastOverflowedUintToInt struct {
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintToInt(uint256 value)
func SafeCastSafeCastOverflowedUintToIntErrorID() common.Hash {
	return common.HexToHash("0x24775e0629ae69d78c11bae050651b81820407f300ff750ff2be51e4ce75c37f")
}

// UnpackSafeCastOverflowedUintToIntError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintToInt(uint256 value)
func (safeCast *SafeCast) UnpackSafeCastOverflowedUintToIntError(raw []byte) (*SafeCastSafeCastOverflowedUintToInt, error) {
	out := new(SafeCastSafeCastOverflowedUintToInt)
	if err := safeCast.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintToInt", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SenderCreatorMetaData contains all meta data concerning the SenderCreator contract.
var SenderCreatorMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"createSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCallData\",\"type\":\"bytes\"}],\"name\":\"initEip7702Sender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "968f0093411ddc152f9da67991f7a9cf5d",
	Bin: "0x60a0604052348015600e575f5ffd5b503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250506080516108c26100685f395f818160b0015281816101de015261020201526108c25ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063570e1a3614610043578063b0d691fe14610073578063c09ad0d914610091575b5f5ffd5b61005d6004803603810190610058919061039e565b6100ad565b60405161006a9190610428565b60405180910390f35b61007b6101dc565b6040516100889190610428565b60405180910390f35b6100ab60048036038101906100a691906105a3565b610200565b005b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461013c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013390610657565b60405180910390fd5b5f83835f906014926101509392919061067d565b9061015b91906106f8565b60601c90505f848460149080926101749392919061067d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090505f60205f8351602085015f875af1905080156101d3575f5193505b50505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461028e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028590610657565b60405180910390fd5b5f5f5f8351602085015f875af19050806102ef575f6102ae6108006102f4565b90505f816040517f65c8fd4d0000000000000000000000000000000000000000000000000000000081526004016102e692919061084b565b60405180910390fd5b505050565b60603d5f83111561030c578281111561030b578290505b5b604051602082018101604052818152815f602083013e8092505050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261035e5761035d61033d565b5b8235905067ffffffffffffffff81111561037b5761037a610341565b5b60208301915083600182028301111561039757610396610345565b5b9250929050565b5f5f602083850312156103b4576103b3610335565b5b5f83013567ffffffffffffffff8111156103d1576103d0610339565b5b6103dd85828601610349565b92509250509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610412826103e9565b9050919050565b61042281610408565b82525050565b5f60208201905061043b5f830184610419565b92915050565b61044a81610408565b8114610454575f5ffd5b50565b5f8135905061046581610441565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6104b58261046f565b810181811067ffffffffffffffff821117156104d4576104d361047f565b5b80604052505050565b5f6104e661032c565b90506104f282826104ac565b919050565b5f67ffffffffffffffff8211156105115761051061047f565b5b61051a8261046f565b9050602081019050919050565b828183375f83830152505050565b5f610547610542846104f7565b6104dd565b9050828152602081018484840111156105635761056261046b565b5b61056e848285610527565b509392505050565b5f82601f83011261058a5761058961033d565b5b813561059a848260208601610535565b91505092915050565b5f5f604083850312156105b9576105b8610335565b5b5f6105c685828601610457565b925050602083013567ffffffffffffffff8111156105e7576105e6610339565b5b6105f385828601610576565b9150509250929050565b5f82825260208201905092915050565b7f414139372073686f756c642063616c6c2066726f6d20456e747279506f696e745f82015250565b5f6106416020836105fd565b915061064c8261060d565b602082019050919050565b5f6020820190508181035f83015261066e81610635565b9050919050565b5f5ffd5b5f5ffd5b5f5f858511156106905761068f610675565b5b838611156106a1576106a0610679565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b5f82821b905092915050565b5f61070383836106b7565b8261070e81356106c1565b9250601482101561074e576107497fffffffffffffffffffffffffffffffffffffffff000000000000000000000000836014036008026106ec565b831692505b505092915050565b5f819050919050565b5f819050919050565b5f819050919050565b5f61078b61078661078184610756565b610768565b61075f565b9050919050565b61079b81610771565b82525050565b7f4141313320454950373730322073656e64657220696e6974206661696c6564005f82015250565b5f6107d5601f836105fd565b91506107e0826107a1565b602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61081d826107eb565b61082781856107f5565b9350610837818560208601610805565b6108408161046f565b840191505092915050565b5f60608201905061085e5f830185610792565b818103602083015261086f816107c9565b905081810360408301526108838184610813565b9050939250505056fea26469706673582212200747fe7ca42bac0b159be2c8536919a8e04cd78351433270e46514c67063e45064736f6c634300081d0033",
}

// SenderCreator is an auto generated Go binding around an Ethereum contract.
type SenderCreator struct {
	abi abi.ABI
}

// NewSenderCreator creates a new instance of SenderCreator.
func NewSenderCreator() *SenderCreator {
	parsed, err := SenderCreatorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &SenderCreator{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *SenderCreator) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCreateSender is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x570e1a36.
//
// Solidity: function createSender(bytes initCode) returns(address sender)
func (senderCreator *SenderCreator) PackCreateSender(initCode []byte) []byte {
	enc, err := senderCreator.abi.Pack("createSender", initCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateSender is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x570e1a36.
//
// Solidity: function createSender(bytes initCode) returns(address sender)
func (senderCreator *SenderCreator) UnpackCreateSender(data []byte) (common.Address, error) {
	out, err := senderCreator.abi.Unpack("createSender", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackEntryPoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (senderCreator *SenderCreator) PackEntryPoint() []byte {
	enc, err := senderCreator.abi.Pack("entryPoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEntryPoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (senderCreator *SenderCreator) UnpackEntryPoint(data []byte) (common.Address, error) {
	out, err := senderCreator.abi.Unpack("entryPoint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInitEip7702Sender is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc09ad0d9.
//
// Solidity: function initEip7702Sender(address sender, bytes initCallData) returns()
func (senderCreator *SenderCreator) PackInitEip7702Sender(sender common.Address, initCallData []byte) []byte {
	enc, err := senderCreator.abi.Pack("initEip7702Sender", sender, initCallData)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (senderCreator *SenderCreator) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], senderCreator.abi.Errors["FailedOpWithRevert"].ID.Bytes()[:4]) {
		return senderCreator.UnpackFailedOpWithRevertError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// SenderCreatorFailedOpWithRevert represents a FailedOpWithRevert error raised by the SenderCreator contract.
type SenderCreatorFailedOpWithRevert struct {
	OpIndex *big.Int
	Reason  string
	Inner   []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedOpWithRevert(uint256 opIndex, string reason, bytes inner)
func SenderCreatorFailedOpWithRevertErrorID() common.Hash {
	return common.HexToHash("0x65c8fd4dd32f2bb83f7d57728e1afa231e9956aaf4bdeea76c78b967426acb8c")
}

// UnpackFailedOpWithRevertError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedOpWithRevert(uint256 opIndex, string reason, bytes inner)
func (senderCreator *SenderCreator) UnpackFailedOpWithRevertError(raw []byte) (*SenderCreatorFailedOpWithRevert, error) {
	out := new(SenderCreatorFailedOpWithRevert)
	if err := senderCreator.abi.UnpackIntoInterface(out, "FailedOpWithRevert", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ShortStringsMetaData contains all meta data concerning the ShortStrings contract.
var ShortStringsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"}]",
	ID:  "5921c89c729d378db4f8f90b83d636f316",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea264697066735822122059bfdd3bd1edb67b8c18e42ab272da76324d104d18ce2e31a9091fe17b141e7e64736f6c634300081d0033",
}

// ShortStrings is an auto generated Go binding around an Ethereum contract.
type ShortStrings struct {
	abi abi.ABI
}

// NewShortStrings creates a new instance of ShortStrings.
func NewShortStrings() *ShortStrings {
	parsed, err := ShortStringsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ShortStrings{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ShortStrings) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (shortStrings *ShortStrings) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], shortStrings.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return shortStrings.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], shortStrings.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return shortStrings.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ShortStringsInvalidShortString represents a InvalidShortString error raised by the ShortStrings contract.
type ShortStringsInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ShortStringsInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (shortStrings *ShortStrings) UnpackInvalidShortStringError(raw []byte) (*ShortStringsInvalidShortString, error) {
	out := new(ShortStringsInvalidShortString)
	if err := shortStrings.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ShortStringsStringTooLong represents a StringTooLong error raised by the ShortStrings contract.
type ShortStringsStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ShortStringsStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (shortStrings *ShortStrings) UnpackStringTooLongError(raw []byte) (*ShortStringsStringTooLong, error) {
	out := new(ShortStringsStringTooLong)
	if err := shortStrings.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SignedMathMetaData contains all meta data concerning the SignedMath contract.
var SignedMathMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "bcf64d9c1468375457906244955c165c41",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220fe19893cb2fcff1d2297568063d34e82795c5d7e5fe664c658829039603d1d8e64736f6c634300081d0033",
}

// SignedMath is an auto generated Go binding around an Ethereum contract.
type SignedMath struct {
	abi abi.ABI
}

// NewSignedMath creates a new instance of SignedMath.
func NewSignedMath() *SignedMath {
	parsed, err := SignedMathMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &SignedMath{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *SignedMath) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// StakeManagerMetaData contains all meta data concerning the StakeManager contract.
var StakeManagerMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "8f1bbbff4d1f99845320143260dea329d8",
}

// StakeManager is an auto generated Go binding around an Ethereum contract.
type StakeManager struct {
	abi abi.ABI
}

// NewStakeManager creates a new instance of StakeManager.
func NewStakeManager() *StakeManager {
	parsed, err := StakeManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &StakeManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *StakeManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (stakeManager *StakeManager) PackAddStake(unstakeDelaySec uint32) []byte {
	enc, err := stakeManager.abi.Pack("addStake", unstakeDelaySec)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (stakeManager *StakeManager) PackBalanceOf(account common.Address) []byte {
	enc, err := stakeManager.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (stakeManager *StakeManager) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := stakeManager.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDepositTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (stakeManager *StakeManager) PackDepositTo(account common.Address) []byte {
	enc, err := stakeManager.abi.Pack("depositTo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetDepositInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (stakeManager *StakeManager) PackGetDepositInfo(account common.Address) []byte {
	enc, err := stakeManager.abi.Pack("getDepositInfo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDepositInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (stakeManager *StakeManager) UnpackGetDepositInfo(data []byte) (IStakeManagerDepositInfo, error) {
	out, err := stakeManager.abi.Unpack("getDepositInfo", data)
	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}
	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)
	return out0, err
}

// PackUnlockStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (stakeManager *StakeManager) PackUnlockStake() []byte {
	enc, err := stakeManager.abi.Pack("unlockStake")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (stakeManager *StakeManager) PackWithdrawStake(withdrawAddress common.Address) []byte {
	enc, err := stakeManager.abi.Pack("withdrawStake", withdrawAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (stakeManager *StakeManager) PackWithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) []byte {
	enc, err := stakeManager.abi.Pack("withdrawTo", withdrawAddress, withdrawAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// StakeManagerDeposited represents a Deposited event raised by the StakeManager contract.
type StakeManagerDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeManagerDepositedEventName = "Deposited"

// ContractEventName returns the user-defined event name.
func (StakeManagerDeposited) ContractEventName() string {
	return StakeManagerDepositedEventName
}

// UnpackDepositedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (stakeManager *StakeManager) UnpackDepositedEvent(log *types.Log) (*StakeManagerDeposited, error) {
	event := "Deposited"
	if log.Topics[0] != stakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeManagerDeposited)
	if len(log.Data) > 0 {
		if err := stakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeManagerStakeLocked represents a StakeLocked event raised by the StakeManager contract.
type StakeManagerStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const StakeManagerStakeLockedEventName = "StakeLocked"

// ContractEventName returns the user-defined event name.
func (StakeManagerStakeLocked) ContractEventName() string {
	return StakeManagerStakeLockedEventName
}

// UnpackStakeLockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (stakeManager *StakeManager) UnpackStakeLockedEvent(log *types.Log) (*StakeManagerStakeLocked, error) {
	event := "StakeLocked"
	if log.Topics[0] != stakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeManagerStakeLocked)
	if len(log.Data) > 0 {
		if err := stakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeManagerStakeUnlocked represents a StakeUnlocked event raised by the StakeManager contract.
type StakeManagerStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeManagerStakeUnlockedEventName = "StakeUnlocked"

// ContractEventName returns the user-defined event name.
func (StakeManagerStakeUnlocked) ContractEventName() string {
	return StakeManagerStakeUnlockedEventName
}

// UnpackStakeUnlockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (stakeManager *StakeManager) UnpackStakeUnlockedEvent(log *types.Log) (*StakeManagerStakeUnlocked, error) {
	event := "StakeUnlocked"
	if log.Topics[0] != stakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeManagerStakeUnlocked)
	if len(log.Data) > 0 {
		if err := stakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeManagerStakeWithdrawn represents a StakeWithdrawn event raised by the StakeManager contract.
type StakeManagerStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const StakeManagerStakeWithdrawnEventName = "StakeWithdrawn"

// ContractEventName returns the user-defined event name.
func (StakeManagerStakeWithdrawn) ContractEventName() string {
	return StakeManagerStakeWithdrawnEventName
}

// UnpackStakeWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (stakeManager *StakeManager) UnpackStakeWithdrawnEvent(log *types.Log) (*StakeManagerStakeWithdrawn, error) {
	event := "StakeWithdrawn"
	if log.Topics[0] != stakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeManagerStakeWithdrawn)
	if len(log.Data) > 0 {
		if err := stakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeManagerWithdrawn represents a Withdrawn event raised by the StakeManager contract.
type StakeManagerWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const StakeManagerWithdrawnEventName = "Withdrawn"

// ContractEventName returns the user-defined event name.
func (StakeManagerWithdrawn) ContractEventName() string {
	return StakeManagerWithdrawnEventName
}

// UnpackWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (stakeManager *StakeManager) UnpackWithdrawnEvent(log *types.Log) (*StakeManagerWithdrawn, error) {
	event := "Withdrawn"
	if log.Topics[0] != stakeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeManagerWithdrawn)
	if len(log.Data) > 0 {
		if err := stakeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StorageSlotMetaData contains all meta data concerning the StorageSlot contract.
var StorageSlotMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "18119f39eca68ff67e22d3b1bd81ca3df1",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220d0663ae3330dc09ed9579fbe24f4b5d29b595c829fdef2bf34e700ac27875ad264736f6c634300081d0033",
}

// StorageSlot is an auto generated Go binding around an Ethereum contract.
type StorageSlot struct {
	abi abi.ABI
}

// NewStorageSlot creates a new instance of StorageSlot.
func NewStorageSlot() *StorageSlot {
	parsed, err := StorageSlotMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &StorageSlot{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *StorageSlot) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StringsInvalidAddressFormat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StringsInvalidChar\",\"type\":\"error\"}]",
	ID:  "9feb222b0160bac7e1e71a58884bbcf822",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220663621fd716cd85299c1a0cc651ad2c369c5a34e1835de725610a2b355044e5064736f6c634300081d0033",
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	abi abi.ABI
}

// NewStrings creates a new instance of Strings.
func NewStrings() *Strings {
	parsed, err := StringsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Strings{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Strings) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (strings *Strings) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], strings.abi.Errors["StringsInsufficientHexLength"].ID.Bytes()[:4]) {
		return strings.UnpackStringsInsufficientHexLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], strings.abi.Errors["StringsInvalidAddressFormat"].ID.Bytes()[:4]) {
		return strings.UnpackStringsInvalidAddressFormatError(raw[4:])
	}
	if bytes.Equal(raw[:4], strings.abi.Errors["StringsInvalidChar"].ID.Bytes()[:4]) {
		return strings.UnpackStringsInvalidCharError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// StringsStringsInsufficientHexLength represents a StringsInsufficientHexLength error raised by the Strings contract.
type StringsStringsInsufficientHexLength struct {
	Value  *big.Int
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringsInsufficientHexLength(uint256 value, uint256 length)
func StringsStringsInsufficientHexLengthErrorID() common.Hash {
	return common.HexToHash("0xe22e27eb9593f9947dc34771120a3349dd201c662753f0b60502fc1d8e422233")
}

// UnpackStringsInsufficientHexLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringsInsufficientHexLength(uint256 value, uint256 length)
func (strings *Strings) UnpackStringsInsufficientHexLengthError(raw []byte) (*StringsStringsInsufficientHexLength, error) {
	out := new(StringsStringsInsufficientHexLength)
	if err := strings.abi.UnpackIntoInterface(out, "StringsInsufficientHexLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StringsStringsInvalidAddressFormat represents a StringsInvalidAddressFormat error raised by the Strings contract.
type StringsStringsInvalidAddressFormat struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringsInvalidAddressFormat()
func StringsStringsInvalidAddressFormatErrorID() common.Hash {
	return common.HexToHash("0x1d15ae44cf5601ace2ebdfa1451a862de1975935c0c8ba6788ae598e5e29a6dd")
}

// UnpackStringsInvalidAddressFormatError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringsInvalidAddressFormat()
func (strings *Strings) UnpackStringsInvalidAddressFormatError(raw []byte) (*StringsStringsInvalidAddressFormat, error) {
	out := new(StringsStringsInvalidAddressFormat)
	if err := strings.abi.UnpackIntoInterface(out, "StringsInvalidAddressFormat", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StringsStringsInvalidChar represents a StringsInvalidChar error raised by the Strings contract.
type StringsStringsInvalidChar struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringsInvalidChar()
func StringsStringsInvalidCharErrorID() common.Hash {
	return common.HexToHash("0x94e2737eaa44cfdde863f17909ef1e0595339c0eb63c5453ecb27449d2dcd443")
}

// UnpackStringsInvalidCharError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringsInvalidChar()
func (strings *Strings) UnpackStringsInvalidCharError(raw []byte) (*StringsStringsInvalidChar, error) {
	out := new(StringsStringsInvalidChar)
	if err := strings.abi.UnpackIntoInterface(out, "StringsInvalidChar", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TransientSlotMetaData contains all meta data concerning the TransientSlot contract.
var TransientSlotMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "e0e7521746c1dbaec6bed9e99e7d317443",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea26469706673582212206e679e516a8916cd33426cc8c088438b3f608aaf5982a3374e4700842dc8759364736f6c634300081d0033",
}

// TransientSlot is an auto generated Go binding around an Ethereum contract.
type TransientSlot struct {
	abi abi.ABI
}

// NewTransientSlot creates a new instance of TransientSlot.
func NewTransientSlot() *TransientSlot {
	parsed, err := TransientSlotMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TransientSlot{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TransientSlot) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UserOperationLibMetaData contains all meta data concerning the UserOperationLib contract.
var UserOperationLibMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PAYMASTER_DATA_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAYMASTER_POSTOP_GAS_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAYMASTER_VALIDATION_GAS_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "bc385f996edff2b3ce290ff387d17a7859",
	Bin: "0x61010a61004d600b8282823980515f1a6073146041577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106046575f3560e01c806325093e1b14604a578063b29a8ff4146064578063ede3150214607e575b5f5ffd5b60506098565b604051605b919060bd565b60405180910390f35b606a609d565b6040516075919060bd565b60405180910390f35b608460a2565b604051608f919060bd565b60405180910390f35b602481565b601481565b603481565b5f819050919050565b60b78160a7565b82525050565b5f60208201905060ce5f83018460b0565b9291505056fea264697066735822122091e3bc019d9813d01310b138d23f0f9856f63af1a3a5ca7c8458dfcdcef73ec664736f6c634300081d0033",
}

// UserOperationLib is an auto generated Go binding around an Ethereum contract.
type UserOperationLib struct {
	abi abi.ABI
}

// NewUserOperationLib creates a new instance of UserOperationLib.
func NewUserOperationLib() *UserOperationLib {
	parsed, err := UserOperationLibMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &UserOperationLib{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *UserOperationLib) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackPAYMASTERDATAOFFSET is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xede31502.
//
// Solidity: function PAYMASTER_DATA_OFFSET() view returns(uint256)
func (userOperationLib *UserOperationLib) PackPAYMASTERDATAOFFSET() []byte {
	enc, err := userOperationLib.abi.Pack("PAYMASTER_DATA_OFFSET")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPAYMASTERDATAOFFSET is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xede31502.
//
// Solidity: function PAYMASTER_DATA_OFFSET() view returns(uint256)
func (userOperationLib *UserOperationLib) UnpackPAYMASTERDATAOFFSET(data []byte) (*big.Int, error) {
	out, err := userOperationLib.abi.Unpack("PAYMASTER_DATA_OFFSET", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPAYMASTERPOSTOPGASOFFSET is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25093e1b.
//
// Solidity: function PAYMASTER_POSTOP_GAS_OFFSET() view returns(uint256)
func (userOperationLib *UserOperationLib) PackPAYMASTERPOSTOPGASOFFSET() []byte {
	enc, err := userOperationLib.abi.Pack("PAYMASTER_POSTOP_GAS_OFFSET")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPAYMASTERPOSTOPGASOFFSET is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x25093e1b.
//
// Solidity: function PAYMASTER_POSTOP_GAS_OFFSET() view returns(uint256)
func (userOperationLib *UserOperationLib) UnpackPAYMASTERPOSTOPGASOFFSET(data []byte) (*big.Int, error) {
	out, err := userOperationLib.abi.Unpack("PAYMASTER_POSTOP_GAS_OFFSET", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPAYMASTERVALIDATIONGASOFFSET is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb29a8ff4.
//
// Solidity: function PAYMASTER_VALIDATION_GAS_OFFSET() view returns(uint256)
func (userOperationLib *UserOperationLib) PackPAYMASTERVALIDATIONGASOFFSET() []byte {
	enc, err := userOperationLib.abi.Pack("PAYMASTER_VALIDATION_GAS_OFFSET")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPAYMASTERVALIDATIONGASOFFSET is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb29a8ff4.
//
// Solidity: function PAYMASTER_VALIDATION_GAS_OFFSET() view returns(uint256)
func (userOperationLib *UserOperationLib) UnpackPAYMASTERVALIDATIONGASOFFSET(data []byte) (*big.Int, error) {
	out, err := userOperationLib.abi.Unpack("PAYMASTER_VALIDATION_GAS_OFFSET", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}
