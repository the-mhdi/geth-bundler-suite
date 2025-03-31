package mempool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/the-mhdi/geth-bundler-suite/bundler/bindings/entrypoint"
	"github.com/the-mhdi/geth-bundler-suite/bundler/types"
)

//implementation of the canonical public ERC-4337 mempool as mentioned in erc-7562

type Monitor struct {
	eth        *ethclient.Client
	EntryPoint map[common.Address]*entrypoint.EntryPoint
}

type ValidationManager interface {
	GREP() (bool, error) //General Reputation Rules
	OP() (bool, error)   //Opcode Rules
	COD() (bool, error)  ///Code Rule
	STO() (bool, error)  //Storage Rules
	SREP() (bool, error) //Staked Entities Reputation Rules
	EREP() (bool, error) //Entity Specific Rules
	UREP() (bool, error) //Unstaked Paymasters Reputation Rules
}

/**Entity-specific Rules**/

/*	For each paymaster, the bundler must track the total gas UserOperations using this paymaster may consume.
Bundler should not accept a new UserOperation with a paymaster to the mempool if the maximum total gas cost of all userops in the mempool, including this new UserOperation, is above the deposit of that paymaster at the current gas price.
*/

func EREP010(uo *types.UserOperation, eth *ethclient.Client) {

	ep := entrypoint.NewEntryPoint()
	ep.PackGetDepositInfo()

}

func (m *Monitor) getDepositInfo(account common.Address) {

	if _, ok := m.EntryPoint[account]; !ok {
		m.EntryPoint[account] = entrypoint.NewEntryPoint()
	}
	contract := m.EntryPoint[account].Instance(m.eth, account)
	if contract == nil {
		return
	} // handle error
	// Call the getDepositInfo function on the EntryPoint contract
	var depositInfo []any
	contract.Call(nil, &depositInfo, "getDepositInfo", account)

	// Process the depositInfo as needed
	// For example, you can extract the deposit amount and other relevant information
	// depositAmount := depositInfo[0] // Assuming the first element is the deposit amount
	// otherInfo := depositInfo[1:]    // Other relevant information
	// You can now use the deposit amount and other information for further processing
	// Example: Print the deposit amount
}
