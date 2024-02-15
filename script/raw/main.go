package main

import (
	"github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/scanner"
	"github.com/leon123858/go-aid/service/sqlite"
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
	db := sqlite.Client{}
	if sqlite.New(&db) != nil {
		log.Fatal("sqlite init failed")
	}

	// Get balance
	balance, err := chain.GetBalance("", 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Balance: %f", balance)

	// Get unspent
	unspentList, err := scanner.ListUnspent(chain, &db, []string{}, 2)
	if err != nil {
		log.Fatal(err)
	}
	if len(*unspentList) == 0 {
		unspentList, _ = scanner.ListUnspent(chain, &db, []string{}, 2)
	}
	fee := 0.001
	var unspent our_chain_rpc.Unspent
	for _, item := range *unspentList {
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
		Action: our_chain_rpc.ContractActionDeploy,
		Code: `
#include <ourcontract.h>
#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

extern "C" int contract_main(int argc, char **argv) {
  // pure mode
  if (!check_runtime_can_write_db()) {
    std::cerr << "runtime is pure mode" << std::endl;
    json j = state_read();
    std::cerr << "get state: " << j.dump() << std::endl;
    std::cerr << "pre txid: " << get_pre_txid() << std::endl;
    // some operation
    j.push_back("pure click: " + std::to_string((size_t)j.size()));
    state_write(j);
    return 0;
  }
  // call contract state
  if (state_exist()) {
    json j = state_read();
    std::cerr << "get state: " << j.dump() << std::endl;
    std::cerr << "pre txid: " << get_pre_txid() << std::endl;
    // some operation
    j.push_back("more click: " + std::to_string((size_t)j.size()));
    state_write(j);
    return 0;
  }
  // init state
  std::cerr << "read state error" << std::endl;
  std::cerr << "pre txid: " << get_pre_txid() << std::endl;
  json j;
  j.push_back("baby cute");
  j.push_back(1);
  j.push_back(true);
  state_write(j);
  return 0;
}`,
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
	log.Printf("should output Contract state: [\"baby cute\",1,true,\"pure click: 3\"]\n")
}
