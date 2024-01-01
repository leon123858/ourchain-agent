package scanner

import (
	"github.com/leon123858/go-aid/service/rpc"
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
	list, err := ListUnspent(chain, []string{"mvehVE6vb5yqoZ4FSeNmJpjacddSdWhh3A"}, 6)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("List Unspent: %+v", list)
}
