package our_chain_rpc

type ContractDeployResult struct {
	TransactionId   string `json:"txid"`
	ContractAddress string `json:"contract address"`
}
