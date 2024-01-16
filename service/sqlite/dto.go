package sqlite

type utxo struct {
	utxoSearchArgument
	Vout   int
	Amount float64
}

type utxoSearchArgument struct {
	ID      string
	Address string
}
