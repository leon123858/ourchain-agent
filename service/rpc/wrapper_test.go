package our_chain_rpc

import (
	"testing"
)

func TestBitcoind_GetContractGeneralInterface(t *testing.T) {
	chain := initChain()
	// deploy a contract
	contractAddr, err := chain.DeployContract("/root/Desktop/ourchain/sample.cpp")
	if err != nil {
		t.Fatal(err)
		return
	}
	// mine a block
	_, err = chain.GenerateBlock(1)
	if err != nil {
		t.Fatal(err)
		return
	}
	generalInterface, err := chain.GetContractGeneralInterface(contractAddr)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("Info: %+v", generalInterface)
	print("protocol: ", generalInterface.Protocol, "\n")
	println("version: ", generalInterface.Version, "\n")
}
