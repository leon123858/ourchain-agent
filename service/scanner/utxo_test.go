package scanner

import (
	"github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/sqlite"
	"github.com/leon123858/go-aid/utils"
	"log"
	"testing"
)

func initChain() *our_chain_rpc.Bitcoind {
	println("> Setup Test")
	utils.LoadConfig("../../config.toml")
	chain, err := our_chain_rpc.New(
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

func TestListUnspent(t *testing.T) {
	chain := initChain()
	db := sqlite.Client{}
	if sqlite.New(&db) != nil {
		log.Fatal("sqlite init failed")
	}
	_, err := ListUnspent(chain, &db, []string{"mjynSV5ztd1qKmqR2npEpp4HDUL9xvW5WU"}, 100)
	if err != nil {
		t.Fatal(err)
	}
	//t.Logf("List Unspent: %+v", list)
}
