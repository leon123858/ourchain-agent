package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type Client struct {
	DbPath   string
	Instance *sql.DB
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

func (client *Client) createUtxo(item utxo) (err error) {
	if client.Instance == nil {
		return err
	}
	_, err = client.Instance.Exec("INSERT INTO utxo(id, vout, address, amount) VALUES(?, ?, ?, ?)", item.ID, item.Vout, item.Address, item.Amount)
	return err
}

func (client *Client) deleteUtxo(id string) (err error) {
	if client.Instance == nil {
		return err
	}
	_, err = client.Instance.Exec("DELETE FROM utxo WHERE id=?", id)
	return err
}

func (client *Client) createUtxoList(utxoList []utxo) (err error) {
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

func (client *Client) deleteUtxoList(idList []string) (err error) {
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

func (client *Client) getUtxoList(arg utxoSearchArgument) (result []utxo, err error) {
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
		var item utxo
		err = rows.Scan(&item.ID, &item.Vout, &item.Address, &item.Amount)
		if err != nil {
			return result, err
		}
		result = append(result, item)
	}
	return result, err
}
