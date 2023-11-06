package our_chain_rpc

type ContractDeployResult struct {
	TransactionId   string `json:"txid"`
	ContractAddress string `json:"contract address"`
}

type RawTransactionCreateResult struct {
	Hex             string `json:"hex"`
	ContractAddress string `json:"contractAddress"`
}

type Unspent struct {
	Txid          string  `json:"txid"`
	Vout          int     `json:"vout"`
	Address       string  `json:"address"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Confirmations int     `json:"confirmations"`
	Spendable     bool    `json:"spendable"`
	Solvable      bool    `json:"solvable"`
	Safe          bool    `json:"safe"`
}

type TxInput struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}

type TxOutput struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

type SignedTx struct {
	Hex      string `json:"hex"`
	Complete bool   `json:"complete"`
}

type Transaction struct {
	Amount            float64             `json:"amount"`
	Fee               float64             `json:"fee"`
	Confirmations     int                 `json:"confirmations"`
	Generated         bool                `json:"generated"`
	Trusted           bool                `json:"trusted"`
	BlockHash         string              `json:"blockhash"`
	BlockHeight       int                 `json:"blockheight"`
	BlockIndex        int                 `json:"blockindex"`
	BlockTime         float64             `json:"blocktime"`
	TxID              string              `json:"txid"`
	WalletConflicts   []string            `json:"walletconflicts"`
	Time              float64             `json:"time"`
	TimeReceived      float64             `json:"timereceived"`
	Comment           string              `json:"comment"`
	Bip125Replaceable string              `json:"bip125-replaceable"`
	Details           []TransactionDetail `json:"details"`
	Hex               string              `json:"hex"`
	Decoded           interface{}         `json:"decoded"`
}

type TransactionDetail struct {
	InvolvesWatchonly bool    `json:"involvesWatchonly"`
	Address           string  `json:"address"`
	Category          string  `json:"category"`
	Amount            float64 `json:"amount"`
	Label             string  `json:"label"`
	Vout              int     `json:"vout"`
	Fee               float64 `json:"fee"`
	Abandoned         bool    `json:"abandoned"`
}

type ContractAction int

// enum for action in ContractMessage
const (
	ContractNotExist ContractAction = iota
	ContractActionDeploy
	ContractActionCall
)

type ContractMessage struct {
	Action  ContractAction `json:"action"`
	Code    string         `json:"code"`
	Address string         `json:"address"`
	Args    []string       `json:"args"`
}
