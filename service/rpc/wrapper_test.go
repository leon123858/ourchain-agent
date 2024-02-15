package our_chain_rpc

import (
	"testing"
)

func TestBitcoind_GetContractGeneralInterface(t *testing.T) {
	chain := initChain()
	generalInterface, err := chain.GetContractGeneralInterface("2f6088e76e21457261d5a059825b37186f6cccbd48d8cc160de182c83a081ebf")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Info: %+v", generalInterface)
	print("protocol: ", generalInterface.Protocol, "\n")
	println("version: ", generalInterface.Version, "\n")
}
