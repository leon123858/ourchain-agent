package sqlite

type Utxo struct {
	UtxoSearchArgument
	Vout   int
	Amount float64
}

type UtxoSearchArgument struct {
	ID      string
	Address string
}

type Block struct {
	Height uint64
	Hash   string
}
