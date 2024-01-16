package scanner

import (
	OurChainRpc "github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/sqlite"
)

func addBlocksToSame(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind) (err error) {
	info, err := rpc.GetBlockChainInfo()
	if err != nil {
		return
	}
	for i := curLocalChain.Length + 1; i < uint64(info.Blocks); i++ {
		var blockHash string
		blockHash, err = rpc.GetBlockHash(i)
		if err != nil {
			return
		}
		var blockInfo OurChainRpc.BlockInfo
		blockInfo, err = rpc.GetBlock(blockHash)
		if err != nil {
			return
		}
		for _, tx := range blockInfo.Tx {
			var txInfo OurChainRpc.Transaction
			txInfo, err = rpc.GetRawTransaction(tx)
			if err != nil {
				return
			}
			for _, vout := range txInfo.Vout {
				if vout.ScriptPubKey.Type == "pubkey" && vout.ScriptPubKey.Addresses != nil {
					err = curLocalChain.Client.CreateUtxo(sqlite.Utxo{UtxoSearchArgument: sqlite.UtxoSearchArgument{ID: tx, Address: vout.ScriptPubKey.Addresses[0]}, Vout: vout.N, Amount: vout.Value})
					if err != nil {
						return err
					}
					err = curLocalChain.Client.CreateBlock(sqlite.Block{Height: i, Hash: blockHash})
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return
}

func minusBlocksToSame(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind) (err error) {
	return
}
