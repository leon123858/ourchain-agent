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
	//chain := initChain()
	//_, err := chain.GetBlock("4cfdbf0980085181684b7ea50ab3c4206116acc20491befe42f1ab47a3a15689")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Logf("Block Info: %+v", blockInfo)
	//print("block hash: ", blockInfo.Hash, "\n")
	//print("block txs: ", blockInfo.Tx[0], "\n")
}

func TestBitcoind_GetRawTransaction(t *testing.T) {
	//chain := initChain()
	//_, err := chain.GetRawTransaction("82e828404ba561e12483c97d486cad7fc6fa1f1beaf450c9d7028efcc5304dbd")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Logf("Tx Info: %+v", tx)
	//print("tx hash: ", tx.TxID, "\n")
	//print("tx address: ", tx.Vout[0].ScriptPubKey.Addresses[0], "\n")
}

func TestBitcoind_GetBlockHash(t *testing.T) {
	chain := initChain()
	blockHash, err := chain.GetBlockHash(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Block Hash: %+v", blockHash)
	print("block hash: ", blockHash, "\n")
}

func TestBitcoind_GenerateBlock(t *testing.T) {
	chain := initChain()
	blockHash, err := chain.GenerateBlock(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Block Hash: %+v", blockHash)
}

func TestBitcoind_GetNewAddress(t *testing.T) {
	chain := initChain()
	address, err := chain.GetNewAddress()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Address: %+v", address)
}

func TestBitcoind_GenerateToAddress(t *testing.T) {
	chain := initChain()
	blockHash, err := chain.GenerateToAddress(3, "2N8hwP1WmJrFF5QWABn38y63uYLhnJYJYTF")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Block Hash: %+v", blockHash)
}
