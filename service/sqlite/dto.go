package sqlite

import our_chain_rpc "github.com/leon123858/go-aid/service/rpc"

type Utxo struct {
	ID          string // txid
	Address     string
	Vout        int
	Amount      float64
	IsSpent     bool
	BlockHeight uint64
}

type Block struct {
	Height uint64
	Hash   string
}

type PreUtxo struct {
	TxID    string
	PreTxID string
	PreVout int
}

type Contract struct {
	TxID            string
	ContractAddress string
	ContractAction  our_chain_rpc.ContractAction
}
