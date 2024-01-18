package scanner

import (
	OurChainRpc "github.com/leon123858/go-aid/service/rpc"
)

func addBlocksToSame(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind, height uint64) (err error) {

	return
}

func addBlocksCoder(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind, height uint64) (*[]command, error) {
	var err error
	commandList := make([]command, 0)
	for i := curLocalChain.Length + 1; i < height; i++ {
		var blockHash string
		blockHash, err = rpc.GetBlockHash(i)
		if err != nil {
			return nil, err
		}
		var blockInfo OurChainRpc.BlockInfo
		blockInfo, err = rpc.GetBlock(blockHash)
		if err != nil {
			return nil, err
		}
		// args: height, blockHash
		commandList = append(commandList, *newCommand(ADD_BLOCK, i, blockHash))
		for _, tx := range blockInfo.Tx {
			var txInfo OurChainRpc.Transaction
			txInfo, err = rpc.GetRawTransaction(tx)
			if err != nil {
				return nil, err
			}
			for _, vin := range txInfo.Vin {
				if vin.Coinbase != "" {
					// coinbase do not need do anything
					continue
				} else if vin.Txid != "" {
					// args: txid, vout
					commandList = append(commandList, *newCommand(UPDATA_UTXO, vin.Txid, vin.Vout))
					// args: txid, pre_txid, pre_vout
					commandList = append(commandList, *newCommand(ADD_PREUTXO, tx, vin.Txid, vin.Vout))
				}
			}
			for _, vout := range txInfo.Vout {
				if vout.ScriptPubKey.Type == "pubkey" && vout.ScriptPubKey.Addresses != nil {
					// args: txid, vout, address, amount, is_spent, block_height
					commandList = append(commandList, *newCommand(ADD_UTXO, tx, vout.N, vout.ScriptPubKey.Addresses[0], vout.Value, 0, i))
				}
			}
		}
	}
	return &commandList, err
}

func minusBlocksToSame(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind, remoteHeight uint64) (err error) {
	// TODO: delete utxo and block from db to targetHeight
	return
}

func minusBlocksCoder(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind, remoteHeight uint64) (*[]command, error) {
	var err error
	commandList := make([]command, 0)
	targetHeight := min(curLocalChain.Length, remoteHeight)
	for {
		if targetHeight == 0 {
			break
		}
		var blockHash string
		blockHash, err = rpc.GetBlockHash(targetHeight)
		if err != nil {
			return nil, err
		}
		var localHash string
		localHash, err = curLocalChain.Client.GetBlockHash(targetHeight)
		if err != nil {
			return nil, err
		}
		if blockHash == localHash {
			break
		}
		targetHeight--
	}
	// args: height
	commandList = append(commandList, *newCommand(REMOVE_PREUTXO, targetHeight))
	return &commandList, err
}
