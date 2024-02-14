package scanner

import (
	OurChainRpc "github.com/leon123858/go-aid/service/rpc"
)

// local chain 加長到 high
func addBlocksToSame(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind, height uint64) (err error) {
	commands, err := addBlocksCoder(curLocalChain, rpc, height)
	if err != nil {
		return
	}
	newTx, err := curLocalChain.Client.Instance.Begin()
	if err != nil {
		return
	}
	rawCommandList, err := compileAdd(newTx, commands)
	if err != nil {
		return
	}
	for _, item := range *rawCommandList {
		err = item.Exec()
		if err != nil {
			originErr := err
			err = newTx.Rollback()
			if err != nil {
				return err
			}
			return originErr
		}
	}
	err = newTx.Commit()
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
			// args: txid, action, contract
			commandList = append(commandList, *newCommand(ADD_TX, tx, txInfo.Action, txInfo.Contract))
			for _, vout := range txInfo.Vout {
				if (vout.ScriptPubKey.Type == "pubkey" || vout.ScriptPubKey.Type == "pubkeyhash") && vout.ScriptPubKey.Addresses != nil {
					if vout.Value > 0 {
						// args: txid, vout, address, amount, is_spent, block_height
						commandList = append(commandList, *newCommand(ADD_UTXO, tx, vout.N, vout.ScriptPubKey.Addresses[0], vout.Value, false, i))
					}
				}
			}
			for _, vin := range txInfo.Vin {
				if vin.Coinbase != "" {
					// coinbase do not need do anything
					continue
				} else if vin.Txid != "" {
					// args: txid, pre_txid, pre_vout
					commandList = append(commandList, *newCommand(ADD_PREUTXO, tx, vin.Txid, vin.Vout))
				}
			}
		}
	}
	return &commandList, err
}

// local chain 減短到 min(遠端高度, 地端高度)
func minusBlocksToSame(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind, remoteHeight uint64) (err error) {
	commands, err := minusBlocksCoder(curLocalChain, rpc, remoteHeight)
	if err != nil {
		return
	}
	newTx, err := curLocalChain.Client.Instance.Begin()
	if err != nil {
		return
	}
	rawCommandList, err := compileMinus(curLocalChain.Client, newTx, commands)
	if err != nil {
		return
	}
	for _, item := range *rawCommandList {
		err = item.Exec()
		if err != nil {
			originErr := err
			err = newTx.Rollback()
			if err != nil {
				return err
			}
			return originErr
		}
	}
	err = newTx.Commit()
	return
}

func minusBlocksCoder(curLocalChain *localChain, rpc *OurChainRpc.Bitcoind, remoteHeight uint64) (*[]command, error) {
	var err error
	commandList := make([]command, 0)
	targetHeight := min(curLocalChain.Length, remoteHeight)
	if targetHeight < curLocalChain.Length {
		for i := curLocalChain.Length; i > targetHeight; i-- {
			// args: height
			commandList = append(commandList, *newCommand(REMOVE_PREUTXO, i))
		}
	}
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
		// args: height localHash remoteHash
		commandList = append(commandList, *newCommand(REMOVE_PREUTXO, targetHeight, localHash, blockHash))
		targetHeight--
	}
	return &commandList, err
}
