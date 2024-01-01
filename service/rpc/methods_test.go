package our_chain_rpc

import (
	"github.com/leon123858/go-aid/utils"
	"log"
	"testing"
)

func initChain() *Bitcoind {
	println("> Setup Test")
	utils.LoadConfig("../../config.toml")
	chain, err := New(
		utils.OurChainConfigInstance.ServerHost,
		utils.OurChainConfigInstance.ServerPort,
		utils.OurChainConfigInstance.User,
		utils.OurChainConfigInstance.Passwd,
		utils.OurChainConfigInstance.UseSsl)
	if err != nil {
		log.Fatal(err)
	}
	return chain
}

func TestBitcoind_GetBlockChainInfo(t *testing.T) {
	chain := initChain()
	chainInfo, err := chain.GetBlockChainInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Chain Info: %+v", chainInfo)
	print("chain high: ", chainInfo.Blocks, "\n")
}

func TestBitcoind_GetBlock(t *testing.T) {
	chain := initChain()
	blockInfo, err := chain.GetBlock("16e4287f1928844facc298905654bc73766126bc47b98065475946feaaea223a")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Block Info: %+v", blockInfo)
	print("block hash: ", blockInfo.Hash, "\n")
	print("block txs: ", blockInfo.Tx[1], "\n")
}

func TestBitcoind_GetRawTransaction(t *testing.T) {
	chain := initChain()
	tx, err := chain.GetRawTransaction("26cf5c71bf9bf2784b86e8dd49f7269c0c30d74ecc78af172626962c8e07c7f7")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Tx Info: %+v", tx)
	print("tx hash: ", tx.TxID, "\n")
	print("tx address: ", tx.Vout[0].ScriptPubKey.Addresses[0], "\n")
}

func TestBitcoind_GetBlockHash(t *testing.T) {
	chain := initChain()
	blockHash, err := chain.GetBlockHash(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Block Hash: %+v", blockHash)
	print("block hash: ", blockHash, "\n")
}
