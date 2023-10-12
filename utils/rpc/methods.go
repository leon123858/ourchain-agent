package our_chain_rpc

import (
	"encoding/json"
	"errors"
	"strconv"
)

// GetBalance return the balance of the server or of a specific account
// If [account] is "", returns the server's total available balance.
// If [account] is specified, returns the balance in the account
func (b *Bitcoind) GetBalance(account string, minconf uint64) (balance float64, err error) {
	r, err := b.client.call("getbalance", []interface{}{account, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	balance, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

// DeployContract deploy a smart contract in the ourChain, return the contract address.
// contractPath is the path of the contract file in remote device, should be absolute path.
// example: /root/Desktop/ourchain/sample.cpp
func (b *Bitcoind) DeployContract(contractPath string) (address string, err error) {
	rpcResponse, err := b.client.call("deploycontract", []interface{}{contractPath})
	if err = handleError(err, &rpcResponse); err != nil {
		return "", err
	}
	result := ContractDeployResult{}
	// convert []byte to json
	err = json.Unmarshal(rpcResponse.Result, &result)
	if err != nil {
		return "", err
	}
	return result.ContractAddress, nil
}

// CallContract call a smart contract in the ourChain, return txid of the contract.
// contractAddress is the address of the contract.
// contractData is the parameter of the contract.
func (b *Bitcoind) CallContract(contractAddress string, contractData []string) (result string, err error) {
	// create string slice like [contractAddress, ...contractData]
	if len(contractData) == 0 {
		contractData = []string{""}
	}
	args := append([]string{contractAddress}, contractData...)
	rpcResponse, err := b.client.call("callcontract", args)
	if err = handleError(err, &rpcResponse); err != nil {
		return "", err
	}
	result = string(rpcResponse.Result)
	return result, nil
}

// DumpContractMessage dump a smart contract message in the ourChain, return the message(jsonString).
// contractAddress is the address of the contract.
// contractData is the parameter of the contract.
func (b *Bitcoind) DumpContractMessage(contractAddress string, contractData []string) (result string, err error) {
	// create string slice like [contractAddress, ...contractData]
	if len(contractData) == 0 {
		contractData = []string{""}
	}
	args := append([]string{contractAddress}, contractData...)
	rpcResponse, err := b.client.call("dumpcontractmessage", args)
	if err = handleError(err, &rpcResponse); err != nil {
		return "", err
	}
	result = string(rpcResponse.Result)
	return result, nil
}

// GenerateBlock generate a block in the ourChain, return the block hash.
// blockData is the data of the block.
func (b *Bitcoind) GenerateBlock(count uint64) (blockHash []string, err error) {
	if count > 101 {
		return []string{}, errors.New("count should not be greater than 101")
	}
	rpcResponse, err := b.client.call("generate", []uint64{count})
	if err = handleError(err, &rpcResponse); err != nil {
		return []string{}, err
	}
	result := []string{}
	// convert []byte to json
	err = json.Unmarshal(rpcResponse.Result, &result)
	if err != nil {
		return []string{}, err
	}
	return result, nil
}
