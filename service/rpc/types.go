package our_chain_rpc

type ContractDeployResult struct {
	TransactionId   string `json:"txid"`
	ContractAddress string `json:"contract address"`
}

type RawTransactionCreateResult struct {
	Hex             string `json:"hex"`
	ContractAddress string `json:"contractAddress"`
}

type Softfork struct {
	ID      string `json:"id"`
	Version int    `json:"version"`
	Reject  struct {
		Status bool `json:"status"`
	} `json:"reject"`
}

type BIP9Softforks struct {
	CSV    Softfork `json:"csv"`
	Segwit Softfork `json:"segwit"`
}

type ChainInfo struct {
	Chain                string        `json:"chain"`
	Blocks               int           `json:"blocks"`
	Headers              int           `json:"headers"`
	BestBlockHash        string        `json:"bestblockhash"`
	Difficulty           float64       `json:"difficulty"`
	MedianTime           int           `json:"mediantime"`
	VerificationProgress float64       `json:"verificationprogress"`
	ChainWork            string        `json:"chainwork"`
	Pruned               bool          `json:"pruned"`
	Softforks            []Softfork    `json:"softforks"`
	BIP9Softforks        BIP9Softforks `json:"bip9_softforks"`
}

type BlockInfo struct {
	Hash              string   `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	StrippedSize      int      `json:"strippedsize"`
	Size              int      `json:"size"`
	Weight            int      `json:"weight"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	VersionHex        string   `json:"versionHex"`
	MerkleRoot        string   `json:"merkleroot"`
	ContractState     string   `json:"contractstate"`
	Tx                []string `json:"tx"`
	Time              int      `json:"time"`
	MedianTime        int      `json:"mediantime"`
	Nonce             int      `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	ChainWork         string   `json:"chainwork"`
	PreviousBlockHash string   `json:"previousblockhash"`
}

type Unspent struct {
	Txid          string  `json:"txid"`
	Vout          int     `json:"vout"`
	Address       string  `json:"address"`
	Amount        float64 `json:"amount"`
	Confirmations int     `json:"confirmations"`
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

type ScriptPubKey struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int      `json:"reqSigs"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses"`
}

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type Vin struct {
	Coinbase    string    `json:"coinbase,omitempty"`
	Txid        string    `json:"txid,omitempty"`
	Vout        int       `json:"vout,omitempty"`
	ScriptSig   ScriptSig `json:"scriptSig,omitempty"`
	TxinWitness []string  `json:"txinwitness,omitempty"`
	Sequence    int       `json:"sequence"`
}

type Vout struct {
	Value        float64      `json:"value"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type Transaction struct {
	TxID          string         `json:"txid"`
	Hash          string         `json:"hash"`
	Version       int            `json:"version"`
	Size          int            `json:"size"`
	Vsize         int            `json:"vsize"`
	Locktime      int            `json:"locktime"`
	Vin           []Vin          `json:"vin"`
	Vout          []Vout         `json:"vout"`
	Hex           string         `json:"hex"`
	BlockHash     string         `json:"blockhash"`
	Confirmations int            `json:"confirmations"`
	Time          int            `json:"time"`
	BlockTime     int            `json:"blocktime"`
	Action        ContractAction `json:"contractAction"`
	Contract      string         `json:"contractAddress"`
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

type ContractGeneralInterface struct {
	Protocol string `json:"name"`
	Version  string `json:"version"`
}
