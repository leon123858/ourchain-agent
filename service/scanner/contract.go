package scanner

import (
	ourchainrpc "github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/sqlite"
)

func ListContract(chain *ourchainrpc.Bitcoind, db *sqlite.Client, protocol string) (list *[]sqlite.Contract, err error) {
	curLocalChain := newChain(clientWrapper{
		ChainType: LOCAL,
		DB:        db,
	})
	curRemoteChain := newChain(clientWrapper{
		ChainType: REMOTE,
		RPC:       chain,
	})
	err = curLocalChain.InitChainStep()
	if err != nil {
		return
	}
	err = curRemoteChain.InitChainStep()
	if err != nil {
		return
	}
	err = curLocalChain.SyncLength(curRemoteChain)
	if err != nil {
		return
	}
	list, err = curLocalChain.GetContractList(protocol)
	return
}
