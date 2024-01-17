package scanner

import (
	OurChainRpc "github.com/leon123858/go-aid/service/rpc"
)

func addBlocksToSame(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind, height uint64) (err error) {
	for i := curLocalChain.Length + 1; i < height; i++ {
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
					// TODO: add new utxo to db
				}
			}
			for _, vin := range txInfo.Vin {
				if vin.Coinbase != "" {
					// coinbase do not need do anything
					continue
				} else if vin.Txid != "" {
					// TODO: remove utxo from db
				}
			}
		}
	}
	return
}

func minusBlocksToSame(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind, remoteHeight uint64) (err error) {
	targetHeight := min(curLocalChain.Length, remoteHeight)
	for {
		var blockHash string
		blockHash, err = rpc.GetBlockHash(targetHeight)
		if err != nil {
			return
		}
		var localHash string
		localHash, err = curLocalChain.Client.GetBlockHash(targetHeight)
		if err != nil {
			return
		}
		if blockHash == localHash {
			return
		}
		targetHeight--
		if targetHeight == 0 {
			return
		}
	}
	// TODO: delete utxo and block from db to targetHeight
}
