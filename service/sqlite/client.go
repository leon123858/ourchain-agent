package sqlite

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type Client struct {
	DbPath   string
	Instance *sql.DB
}

func (client *Client) Error() string {
	panic("sql client error")
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

func (client *Client) CreateUtxo(item Utxo) (err error) {
	if client.Instance == nil {
		return err
	}
	_, err = client.Instance.Exec("INSERT INTO utxo(id, vout, address, amount) VALUES(?, ?, ?, ?)", item.ID, item.Vout, item.Address, item.Amount)
	return err
}

func (client *Client) DeleteUtxo(id string) (err error) {
	if client.Instance == nil {
		return err
	}
	_, err = client.Instance.Exec("DELETE FROM utxo WHERE id=?", id)
	return err
}

func (client *Client) CreateUtxoList(utxoList []Utxo) (err error) {
	tx, err := client.Instance.Begin()
	if err != nil {
		return err
	}
	prepare, err := tx.Prepare("INSERT INTO utxo(id, vout, address, amount) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	for _, item := range utxoList {
		_, err = prepare.Exec(item.ID, item.Vout, item.Address, item.Amount)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (client *Client) DeleteUtxoList(idList []string) (err error) {
	tx, err := client.Instance.Begin()
	if err != nil {
		return err
	}
	prepare, err := tx.Prepare("DELETE FROM utxo WHERE id=?")
	if err != nil {
		return err
	}
	for _, id := range idList {
		_, err = prepare.Exec(id)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (client *Client) GetUtxoList(arg UtxoSearchArgument) (result []Utxo, err error) {
	var args []interface{}
	var conditions []string
	queryStr := "SELECT id, vout, address, amount FROM utxo"
	if arg.ID != "" {
		conditions = append(conditions, "id = ?")
		args = append(args, arg.ID)
	}
	if arg.Address != "" {
		conditions = append(conditions, "address = ?")
		args = append(args, arg.Address)
	}
	if len(conditions) > 0 {
		queryStr += " WHERE " + strings.Join(conditions, " AND ")
	}
	rows, err := client.Instance.Query(queryStr, args...)
	if err != nil {
		return result, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var item Utxo
		err = rows.Scan(&item.ID, &item.Vout, &item.Address, &item.Amount)
		if err != nil {
			return result, err
		}
		result = append(result, item)
	}
	return result, err
}

func (client *Client) CreateBlock(item Block) (err error) {
	if client.Instance == nil {
		return errors.New("db is nil")
	}
	_, err = client.Instance.Exec("INSERT INTO block(height, hash) VALUES(?, ?)", item.Height, item.Hash)
	return err
}

func (client *Client) DeleteBlock(height int) (err error) {
	if client.Instance == nil {
		return errors.New("db is nil")
	}
	_, err = client.Instance.Exec("DELETE FROM block WHERE height=?", height)
	return err
}

func (client *Client) GetBlocks(length, offset int) (result []Block, err error) {
	if client.Instance == nil {
		return result, errors.New("db is nil")
	}
	rows, err := client.Instance.Query("SELECT height, hash FROM block ORDER BY height DESC LIMIT ? OFFSET ?", length, offset)
	if err != nil {
		return result, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var item Block
		err = rows.Scan(&item.Height, &item.Hash)
		if err != nil {
			return result, err
		}
		result = append(result, item)
	}
	return result, err
}
