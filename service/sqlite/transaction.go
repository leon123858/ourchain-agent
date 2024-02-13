package sqlite

import (
	"database/sql"
	"errors"
)

func BeginTx(client *Client) (tx *sql.Tx, err error) {
	if client.Instance == nil {
		return nil, client.Error()
	}
	return client.Instance.Begin()
}

func CommitTx(tx *sql.Tx) (err error) {
	if tx == nil {
		return errors.New("tx is nil")
	}
	return tx.Commit()
}

func RollbackTx(tx *sql.Tx) (err error) {
	if tx == nil {
		return errors.New("tx is nil")
	}
	return tx.Rollback()
}

func PrepareTx(tx *sql.Tx, sql string) (stmt *sql.Stmt, err error) {
	if tx == nil {
		return nil, errors.New("tx is nil")
	}
	return tx.Prepare(sql)
}

func ExecPrepare(stmt *sql.Stmt, args ...interface{}) (result sql.Result, err error) {
	if stmt == nil {
		return nil, errors.New("stmt is nil")
	}
	return stmt.Exec(args...)
}

func UtxoCreatePrepare(tx *sql.Tx) (stmt *sql.Stmt, err error) {
	return PrepareTx(tx, "INSERT INTO utxo(id, vout, address, amount, is_spent, block_height) VALUES(?, ?, ?, ?, ?, ?)")
}

func UtxoUpdatePrepare(tx *sql.Tx) (stmt *sql.Stmt, err error) {
	return PrepareTx(tx, "UPDATE utxo SET is_spent = ? WHERE id = ? AND vout = ?")
}

func UtxoDeletePrepare(tx *sql.Tx) (stmt *sql.Stmt, err error) {
	return PrepareTx(tx, "DELETE FROM utxo WHERE block_height = ?")
}

func UtxoCreateExec(stmt *sql.Stmt, item Utxo) (result sql.Result, err error) {
	return ExecPrepare(stmt, item.ID, item.Vout, item.Address, item.Amount, item.IsSpent, item.BlockHeight)
}

// UtxoUpdateExec only need item.IsSpent, ID, Vout now
func UtxoUpdateExec(stmt *sql.Stmt, item Utxo) (result sql.Result, err error) {
	return ExecPrepare(stmt, item.IsSpent, item.ID, item.Vout)
}

// UtxoDeleteExec only need item.BlockHeight now
func UtxoDeleteExec(stmt *sql.Stmt, item Utxo) (result sql.Result, err error) {
	return ExecPrepare(stmt, item.BlockHeight)
}

func BlockCreatePrepare(tx *sql.Tx) (stmt *sql.Stmt, err error) {
	return PrepareTx(tx, "INSERT INTO block(height, hash) VALUES(?, ?)")
}

func BlockCreateExec(stmt *sql.Stmt, item Block) (result sql.Result, err error) {
	return ExecPrepare(stmt, item.Height, item.Hash)
}

func BlockDeletePrepare(tx *sql.Tx) (stmt *sql.Stmt, err error) {
	return PrepareTx(tx, "DELETE FROM block WHERE height = ?")
}

// BlockDeleteExec only need item.Height now
func BlockDeleteExec(stmt *sql.Stmt, item Block) (result sql.Result, err error) {
	return ExecPrepare(stmt, item.Height)
}

func TxCreatePrepare(tx *sql.Tx) (stmt *sql.Stmt, err error) {
	return PrepareTx(tx, "INSERT INTO tx(txid, pre_txid, pre_vout) VALUES(?, ?, ?)")
}

func TxCreateExec(stmt *sql.Stmt, pre PreUtxo) (result sql.Result, err error) {
	return ExecPrepare(stmt, pre.TxID, pre.PreTxID, pre.PreVout)
}

func TxDeletePrepare(tx *sql.Tx) (stmt *sql.Stmt, err error) {
	return PrepareTx(tx, "DELETE FROM tx WHERE txid = ?")
}

// TxDeleteExec only need item.TxID now
func TxDeleteExec(stmt *sql.Stmt, item PreUtxo) (result sql.Result, err error) {
	return ExecPrepare(stmt, item.TxID)
}

func ContractCreatePrepare(tx *sql.Tx) (stmt *sql.Stmt, err error) {
	return PrepareTx(tx, "INSERT INTO contract(txid, contract_address, contract_action) VALUES(?, ?, ?)")
}

func ContractCreateExec(stmt *sql.Stmt, item Contract) (result sql.Result, err error) {
	return ExecPrepare(stmt, item.TxID, item.ContractAddress, item.ContractAction)
}

func ContractDeletePrepare(tx *sql.Tx) (stmt *sql.Stmt, err error) {
	return PrepareTx(tx, "DELETE FROM contract WHERE txid = ?")
}

func ContractDeleteExec(stmt *sql.Stmt, item Contract) (result sql.Result, err error) {
	return ExecPrepare(stmt, item.TxID)
}
