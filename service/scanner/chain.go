package scanner

import (
	"errors"
	ourchainrpc "github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/sqlite"
)

type clientWrapper struct {
	ChainType ChainType
	RPC       *ourchainrpc.Bitcoind
	DB        *sqlite.Client
}

type block struct {
	Height uint64
	Hash   string
}

type chainInfo struct {
	Length uint64
	clientWrapper
	FirstHash string
}

type ChainType string

const (
	LOCAL  ChainType = "local"
	REMOTE ChainType = "remote"
)

type abstractChain interface {
	GetBlockChainInfo() (result chainInfo)
	InitChainStep() (err error)
	GetName() (result ChainType)
	SyncLength(remote abstractChain) (err error)
	GetUnspent(addressList []string, confirm int) (result []ourchainrpc.Unspent, err error)
}

type remoteChain struct {
	Chain  []block
	Length uint64
	Client *ourchainrpc.Bitcoind
}

type localChain struct {
	Chain  []block
	Length uint64
	Client *sqlite.Client
}

func newChain(client clientWrapper) (result abstractChain) {
	blocks := make([]block, 0)
	if client.ChainType == REMOTE {
		result = &remoteChain{
			Chain:  blocks,
			Length: 0,
			Client: client.RPC,
		}
		return
	} else if client.ChainType == LOCAL {
		result = &localChain{
			Chain:  blocks,
			Length: 0,
			Client: client.DB,
		}
		return
	}
	panic("invalid chain type")
}

func (chain *remoteChain) GetName() (result ChainType) {
	return REMOTE
}

func (chain *localChain) GetName() (result ChainType) {
	return LOCAL
}

func (chain *remoteChain) GetBlockChainInfo() (result chainInfo) {
	return chainInfo{Length: chain.Length, FirstHash: chain.Chain[0].Hash, clientWrapper: clientWrapper{ChainType: REMOTE, RPC: chain.Client}}
}

func (chain *localChain) GetBlockChainInfo() (result chainInfo) {
	return chainInfo{Length: chain.Length, FirstHash: chain.Chain[0].Hash, clientWrapper: clientWrapper{ChainType: LOCAL, DB: chain.Client}}
}

func (chain *remoteChain) InitChainStep() (err error) {
	rpcClient := chain.Client
	if rpcClient == nil {
		return errors.New("rpc is nil")
	}
	chainInfo, err := rpcClient.GetBlockChainInfo()
	if err != nil {
		return
	}
	chain.Length = uint64(chainInfo.Blocks)
	// get first block
	blockHash, err := rpcClient.GetBlockHash(chain.Length)
	if err != nil {
		return
	}
	chain.Chain = append(chain.Chain, block{chain.Length, blockHash})
	return nil
}

func (chain *localChain) InitChainStep() (err error) {
	dbClient := chain.Client
	if dbClient.Instance == nil {
		return errors.New("db is nil")
	}
	blocks, err := dbClient.GetBlocks(1, 0)
	if err != nil {
		return err
	}
	if len(blocks) == 0 {
		return nil
	}
	chain.Chain = append(chain.Chain, block{blocks[0].Height, blocks[0].Hash})
	chain.Length = blocks[0].Height
	return nil
}

func (chain *localChain) SyncLength(remote abstractChain) (err error) {
	info := remote.GetBlockChainInfo()
	rpc := info.RPC
	if rpc == nil || chain.Client == nil {
		return errors.New("rpc or db is nil")
	}
	if chain.Length == info.Length && chain.Chain[0].Hash == info.FirstHash {
		return nil
	}
	if e := minusBlocksToSame(chain, rpc); e != nil {
		return e
	}
	if chain.Length < info.Length {
		if e := addBlocksToSame(chain, rpc); e != nil {
			return e
		}
	}
	return nil
}

func (chain *remoteChain) SyncLength(remote abstractChain) (err error) {
	panic("should not sync length for remote chain")
}

func (chain *remoteChain) GetUnspent(addressList []string, confirm int) (result []ourchainrpc.Unspent, err error) {
	panic("should not get unspent for remote chain")
}

func (chain *localChain) GetUnspent(addressList []string, confirm int) (result []ourchainrpc.Unspent, err error) {
	// TODO: get unspent from local db
	return
}
