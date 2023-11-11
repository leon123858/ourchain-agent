package our_chain_rpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// GetBalance return the balance of the server or of a specific account
// If [account] is "", returns the server's total available balance.
// If [minconf] is min confirm for coin should be counted.
func (b *Bitcoind) GetBalance(account string, minconf uint64) (balance float64, err error) {
	r, err := b.client.call("getbalance", []interface{}{account, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	balance, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

// GetTransaction get a transaction in the ourChain, return the transaction.
// txid is the id of the transaction.
func (b *Bitcoind) GetTransaction(txid string) (result Transaction, err error) {
	rpcResponse, err := b.client.call("gettransaction", []interface{}{txid})
	if err = handleError(err, &rpcResponse); err != nil {
		return
	}
	err = json.Unmarshal(rpcResponse.Result, &result)
	if err != nil {
		return
	}
	return
}

// ListUnspent return the unspent transaction output of the node
func (b *Bitcoind) ListUnspent(addressList ...string) (result []Unspent, err error) {
	rpcResponse, err := b.client.call("listunspent", []interface{}{6, 9999999, addressList})
	if err = handleError(err, &rpcResponse); err != nil {
		return []Unspent{}, err
	}
	// convert []byte to json
	err = json.Unmarshal(rpcResponse.Result, &result)
	if err != nil {
		return []Unspent{}, err
	}
	return result, nil
}

// DumpPrivKey dump the private key of the address in the ourChain, return the private key.
// address is the address of the private key.
func (b *Bitcoind) DumpPrivKey(address string) (result string, err error) {
	rpcResponse, err := b.client.call("dumpprivkey", []interface{}{address})
	if err = handleError(err, &rpcResponse); err != nil {
		return "", err
	}
	result = strings.Trim(string(rpcResponse.Result), "\"")
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
	result = strings.Trim(string(rpcResponse.Result), "\"")
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
	result = strings.Trim(string(rpcResponse.Result), "\"")
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

// CreateRawTransaction create a raw transaction in the ourChain, return the raw transaction.
// input is the input of the transaction.
// output is the output of the transaction.
func (b *Bitcoind) CreateRawTransaction(input []TxInput, output []TxOutput, contract ContractMessage) (result RawTransactionCreateResult, err error) {
	// create a  map like {"address": "amount"}
	outputMap := make(map[string]string)
	outputMap[output[0].Address] = fmt.Sprintf("%.8f", output[0].Amount)
	rpcResponse, err := b.client.call("createrawtransaction", []interface{}{input, outputMap, contract})
	if err = handleError(err, &rpcResponse); err != nil {
		return
	}
	// convert []byte to json
	err = json.Unmarshal(rpcResponse.Result, &result)
	if err != nil {
		return
	}
	return
}

// SignRawTransaction sign a raw transaction in the ourChain, return the signed transaction.
// rawTx is the raw transaction.
func (b *Bitcoind) SignRawTransaction(rawTx string, privateKey string) (result SignedTx, err error) {
	rpcResponse, err := b.client.call("signrawtransaction", []interface{}{rawTx, []interface{}{}, []interface{}{privateKey}})
	if err = handleError(err, &rpcResponse); err != nil {
		return
	}
	err = json.Unmarshal(rpcResponse.Result, &result)
	if err != nil {
		return
	}
	return
}

// SendRawTransaction send a signed raw transaction in the ourChain, return the transaction id.
// rawTx is the signed raw transaction.
// return the transaction id.
func (b *Bitcoind) SendRawTransaction(rawTx string) (result string, err error) {
	rpcResponse, err := b.client.call("sendrawtransaction", []interface{}{rawTx})
	if err = handleError(err, &rpcResponse); err != nil {
		return "", err
	}
	result = strings.Trim(string(rpcResponse.Result), "\"")
	return
}
