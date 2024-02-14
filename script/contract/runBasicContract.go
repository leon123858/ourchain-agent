package main

import (
	ourChain "github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/utils"
	"log"
)

func main() {
	utils.LoadConfig()
	chain, err := ourChain.New(
		utils.OurChainConfigInstance.ServerHost,
		utils.OurChainConfigInstance.ServerPort,
		utils.OurChainConfigInstance.User,
		utils.OurChainConfigInstance.Passwd,
		utils.OurChainConfigInstance.UseSsl)
	if err != nil {
		log.Fatal(err)
	}

	// Get balance
	balance, err := chain.GetBalance("", 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Balance: %f", balance)

	// Deploy contract
	contractPath := "/root/Desktop/ourchain/sample.cpp"
	address, err := chain.DeployContract(contractPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Contract address: %s", address)

	// Generate block
	blockHash, err := chain.GenerateBlock(2)
	if err != nil {
		log.Fatal(err)
	}
	for _, hash := range blockHash {
		log.Printf("Block hash: %s", hash)
	}

	// Call contract
	_, err = chain.CallContract(address, []string{})
	if err != nil {
		log.Fatal(err)
	}

	// Generate block
	blockHash, err = chain.GenerateBlock(2)
	if err != nil {
		log.Fatal(err)
	}
	for _, hash := range blockHash {
		log.Printf("Block hash: %s", hash)
	}

	// Dump contract message
	message, err := chain.DumpContractMessage(address, []string{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Contract message: %s", message)
}
