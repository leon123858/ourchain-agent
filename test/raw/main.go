package main

import (
	"github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/scanner"
	"github.com/leon123858/go-aid/utils"
	"log"
)

func main() {
	utils.LoadConfig()
	chain, err := our_chain_rpc.New(
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

	// Get unspent
	unspentList, err := scanner.ListUnspent(chain, []string{}, 100)
	if err != nil {
		log.Fatal(err)
	}
	fee := 0.001
	var unspent our_chain_rpc.Unspent
	for _, item := range unspentList {
		if item.Amount > fee {
			unspent = item
			break
		}
	}
	targetUtxo := unspent
	log.Printf("Unspent: %v", targetUtxo.Amount)

	// Create raw transaction
	inputs := []our_chain_rpc.TxInput{{
		Txid: targetUtxo.Txid,
		Vout: targetUtxo.Vout,
	}}
	outputs := []our_chain_rpc.TxOutput{{
		Address: targetUtxo.Address,
		Amount:  targetUtxo.Amount - fee,
	}}
	println("Contract Action", our_chain_rpc.ContractNotExist, our_chain_rpc.ContractActionDeploy, our_chain_rpc.ContractActionCall)
	contract := our_chain_rpc.ContractMessage{
		Action:  our_chain_rpc.ContractActionDeploy,
		Code:    "#include <ourcontract.h>\n#include <iostream>\n#include <stdio.h>\n#include <stdlib.h>\n#include <string.h>\n#include <sys/wait.h>\n#include <unistd.h>\n\nextern \"C\" int contract_main(int argc, char **argv) {\n  // pure mode\n  if (!check_runtime_can_write_db()) {\n    std::cerr << \"runtime is pure mode\" << std::endl;\n    json j = state_read();\n    std::cerr << \"get state: \" << j.dump() << std::endl;\n    std::cerr << \"pre txid: \" << get_pre_txid() << std::endl;\n    // some operation\n    j.push_back(\"pure click: \" + std::to_string((size_t)j.size()));\n    state_write(j);\n    return 0;\n  }\n  // call contract state\n  if (state_exist()) {\n    json j = state_read();\n    std::cerr << \"get state: \" << j.dump() << std::endl;\n    std::cerr << \"pre txid: \" << get_pre_txid() << std::endl;\n    // some operation\n    j.push_back(\"more click: \" + std::to_string((size_t)j.size()));\n    state_write(j);\n    return 0;\n  }\n  // init state\n  std::cerr << \"read state error\" << std::endl;\n  std::cerr << \"pre txid: \" << get_pre_txid() << std::endl;\n  json j;\n  j.push_back(\"baby cute\");\n  j.push_back(1);\n  j.push_back(true);\n  state_write(j);\n  return 0;\n}\n",
		Address: "",
		Args:    []string{},
	}
	rawTx, err := chain.CreateRawTransaction(inputs, outputs, contract)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Raw Contract Address: %s", rawTx.ContractAddress)

	// Dump private key
	privateKey, err := chain.DumpPrivKey(targetUtxo.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Private key: %s", privateKey)

	//Sign raw transaction
	signedTx, err := chain.SignRawTransaction(rawTx.Hex, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Signed transaction: %s", signedTx.Hex)

	//Send raw transaction
	txid, err := chain.SendRawTransaction(signedTx.Hex)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction id: %s", txid)

	// Generate block
	blockHash, err := chain.GenerateBlock(2)
	if err != nil {
		log.Fatal(err)
	}
	for _, hash := range blockHash {
		log.Printf("Block hash: %s", hash)
	}

	// get transaction
	transaction, err := chain.GetRawTransaction(txid)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction: %v", transaction.Confirmations)

	// get contract state
	state, err := chain.DumpContractMessage(rawTx.ContractAddress, []string{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Contract state: %s", state)
}
