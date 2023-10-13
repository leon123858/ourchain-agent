package main

import (
	"log"

	ourChain "github.com/leon123858/go-aid/utils/rpc"
)

const (
	SERVER_HOST       = "127.0.0.1"
	SERVER_PORT       = 8332
	USER              = "test"
	PASSWD            = "test"
	USESSL            = false
	WALLET_PASSPHRASE = "WalletPassphrase"
)

func main() {
	ourChain, err := ourChain.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Fatal(err)
	}

	// Get balance
	balance, err := ourChain.GetBalance("", 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Balance: %f", balance)

	// Deploy contract
	contractPath := "/root/Desktop/ourchain/sample.cpp"
	address, err := ourChain.DeployContract(contractPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Contract address: %s", address)

	// Generate block
	blockHash, err := ourChain.GenerateBlock(1)
	if err != nil {
		log.Fatal(err)
	}
	for _, hash := range blockHash {
		log.Printf("Block hash: %s", hash)
	}

	// Call contract
	_, err = ourChain.CallContract(address, []string{})
	if err != nil {
		log.Fatal(err)
	}

	// Generate block
	blockHash, err = ourChain.GenerateBlock(1)
	if err != nil {
		log.Fatal(err)
	}
	for _, hash := range blockHash {
		log.Printf("Block hash: %s", hash)
	}

	// Dump contract message
	message, err := ourChain.DumpContractMessage(address, []string{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Contract message: %s", message)
}
