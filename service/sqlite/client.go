package sqlite

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	DbPath   string
	Instance *sql.DB
}

func (client *Client) Error() error {
	return errors.New("dbClient native error")
}

func New(client *Client) (err error) {
	if client.DbPath == "" {
		client.DbPath = "./ourchain.db"
	}
	client.Instance, err = sql.Open("sqlite3", client.DbPath)
	if err != nil {
		return err
	}
	err = initTables(client.Instance)
	return err
}

func (client *Client) Close() (err error) {
	if client.Instance != nil {
		err = client.Instance.Close()
	}
	return err
}

func (client *Client) GetFirstBlockInfo() ([]Block, error) {
	row := client.Instance.QueryRow("SELECT * FROM block ORDER BY height DESC LIMIT 1")
	var height uint64
	var hash string
	err := row.Scan(&height, &hash)
	if errors.Is(err, sql.ErrNoRows) {
		return []Block{}, nil
	}
	if err != nil {
		return []Block{}, err
	}
	return []Block{{Height: height, Hash: hash}}, nil
}

func (client *Client) GetBlockHash(height uint64) (string, error) {
	row := client.Instance.QueryRow("SELECT hash FROM block WHERE height = ?", height)
	var hash string
	err := row.Scan(&hash)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (client *Client) GetAddressUtxo(address string, maxHeight int) (*[]Utxo, error) {
	rows, err := client.Instance.Query("SELECT * FROM utxo WHERE address = ? AND block_height <= ? AND is_spent = 0", address, maxHeight)
	if err != nil {
		return &[]Utxo{}, err
	}
	defer func(rows *sql.Rows) {
		e := rows.Close()
		if e != nil {
			panic(e)
		}
	}(rows)
	var result []Utxo
	for rows.Next() {
		var utxo Utxo
		err = rows.Scan(&utxo.ID, &utxo.Vout, &utxo.Address, &utxo.Amount, &utxo.IsSpent, &utxo.BlockHeight)
		if err != nil {
			return &[]Utxo{}, err
		}
		result = append(result, utxo)
	}
	return &result, nil
}

func (client *Client) GetAllUtxo(maxHeight int) (*[]Utxo, error) {
	rows, err := client.Instance.Query("SELECT * FROM utxo WHERE block_height <= ? AND is_spent = 0", maxHeight)
	if err != nil {
		return &[]Utxo{}, err
	}
	defer func(rows *sql.Rows) {
		e := rows.Close()
		if e != nil {
			panic(e)
		}
	}(rows)
	var result []Utxo
	for rows.Next() {
		var utxo Utxo
		err = rows.Scan(&utxo.ID, &utxo.Vout, &utxo.Address, &utxo.Amount, &utxo.IsSpent, &utxo.BlockHeight)
		if err != nil {
			return &[]Utxo{}, err
		}
		result = append(result, utxo)
	}
	return &result, nil
}

func (client *Client) GetUtxoByHeight(height uint64) (*[]Utxo, error) {
	rows, err := client.Instance.Query("SELECT * FROM utxo WHERE block_height = ?", height)
	if err != nil {
		return &[]Utxo{}, err
	}
	defer func(rows *sql.Rows) {
		e := rows.Close()
		if e != nil {
			panic(e)
		}
	}(rows)
	var result []Utxo
	for rows.Next() {
		var utxo Utxo
		err = rows.Scan(&utxo.ID, &utxo.Vout, &utxo.Address, &utxo.Amount, &utxo.IsSpent, &utxo.BlockHeight)
		if err != nil {
			return &[]Utxo{}, err
		}
		result = append(result, utxo)
	}
	return &result, nil
}

func (client *Client) GetPreUtxo(txid string) (*[]PreUtxo, error) {
	rows, err := client.Instance.Query("SELECT pre_txid, pre_vout FROM tx WHERE txid = ?", txid)
	if err != nil {
		return &[]PreUtxo{}, err
	}
	defer func(rows *sql.Rows) {
		e := rows.Close()
		if e != nil {
			panic(e)
		}
	}(rows)
	var result []PreUtxo
	for rows.Next() {
		var preTxID string
		var preVout int
		err := rows.Scan(&preTxID, &preVout)
		if err != nil {
			return &result, err
		}
		result = append(result, PreUtxo{PreTxID: preTxID, PreVout: preVout, TxID: txid})
	}
	return &result, nil
}

func (client *Client) GetBlockTxList(height uint64) (*[]string, error) {
	rows, err := client.Instance.Query("SELECT DISTINCT id FROM utxo WHERE block_height = ?", height)
	if err != nil {
		return &[]string{}, err
	}
	defer func(rows *sql.Rows) {
		e := rows.Close()
		if e != nil {
			panic(e)
		}
	}(rows)
	var result []string
	for rows.Next() {
		var txid string
		err := rows.Scan(&txid)
		if err != nil {
			return &result, err
		}
		result = append(result, txid)
	}
	return &result, nil
}
