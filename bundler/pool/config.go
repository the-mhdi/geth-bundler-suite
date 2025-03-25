package mempool

import (
	"github.com/holiman/uint256"
)

//userops go through these phases : 1. local simulation of validateOp()

// erc-7562 constants
var (
	MIN_UNSTAKE_DELAY                              = 86400
	MIN_STAKE_VALUE                    uint256.Int = *uint256.NewInt(1000000000000000000) //default
	SAME_SENDER_MEMPOOL_COUNT                      = 4                                    //Maximum number of allowed userops in the mempool from a single sender
	SAME_UNSTAKED_ENTITY_MEMPOOL_COUNT             = 10                                   //Maximum number allowed in the mempool of UserOperations referencing the same unstaked entity
	THROTTLED_ENTITY_MEMPOOL_COUNT                 = 4                                    //Number of UserOperations with a throttled entity that can stay in the mempool
	THROTTLED_ENTITY_LIVE_BLOCKS                   = 10                                   // Number of blocks a UserOperations with a throttled entity can stay in mempool
	THROTTLED_ENTITY_BUNDLE_COUNT                  = 4                                    //Number of UserOperations with a throttled entity that can be added in a single bundle
	MIN_INCLUSION_RATE_DENOMINATOR     int64       = 10                                   //A denominator of a formula for entity reputation calculation
	THROTTLING_SLACK                   int64       = 10                                   //Part of a reputation formula that allows entities to legitimately reject some transactions without being throttled
	BAN_SLACK                          int64       = 50
	BAN_OPS_SEEN_PENALTY                           = 10000 // A value to put into the opsSeen counter of entity to declare as banned
	MAX_OPS_ALLOWED_UNSTAKED_ENTITY                = 10000
	PRE_VERIFICATION_OVERHEAD_GAS                  = 50000  //Gas used by the EntryPoint per UserOp that cannot be tracked on-chain
	MAX_VERIFICATION_GAS                           = 500000 //Maximum gas verification functions may use
	MAX_USEROP_SIZE                                = 8192   //Maximum size of a single packed and ABI-encoded UserOperation in bytes
	MAX_CONTEXT_SIZE                               = 2048   //Maximum size of a context byte array returned by a paymaster in a single UserOperation in bytes
	MAX_BUNDLE_SIZE                                = 262144 // Maximum size of an ABI-encoded bundle call to the handleOps function in bytes
	MAX_BUNDLE_CONTEXT_SIZE                        = 65536  //Maximum total size of all context byte arrays returned by all paymasters in all UserOperations in a bundle in bytes
	VALIDATION_GAS_SLACK                           = 4000   // An amount of gas that must be added to the estimations of verificationGasLimit and paymasterVerificationGasLimit
)

type poolConf struct {
}
