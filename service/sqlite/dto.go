package sqlite

type Utxo struct {
	ID          string // txid
	Address     string
	Vout        int
	Amount      float64
	IsSpent     bool
	IsCoinBase  bool
	PreTxID     string
	PreVout     int
	BlockHeight uint64
}

type Block struct {
	Height uint64
	Hash   string
}
