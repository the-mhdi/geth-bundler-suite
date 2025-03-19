package bundler

import (
	"github.com/ethereum/go-ethereum/p2p"
)

type Bundler struct {
	server        p2p.Server
	rpcAPIs       []rpc.API   // List of APIs currently provided by the node
	http          *httpServer //
	ws            *httpServer //
	httpAuth      *httpServer //
	wsAuth        *httpServer //
	ipc           *ipcServer  // Stores information about the ipc http server
	inprocHandler *rpc.Server // In-process RPC request handler to process the API requests
}
