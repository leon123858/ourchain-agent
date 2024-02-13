package sqlite

import (
	"database/sql"
	"errors"
)

func initTables(db *sql.DB) (err error) {
	if db == nil {
		return errors.New("db is nil")
	}
	if err = initUTXOTable(db); err != nil {
		return err
	}
	if err = initBlockTable(db); err != nil {
		return err
	}
	if err = initTxTable(db); err != nil {
		return err
	}
	if err = initContractTable(db); err != nil {
		return err
	}
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}
	return nil
}

func initUTXOTable(db *sql.DB) (err error) {
	creatTable := `
    CREATE TABLE IF NOT EXISTS utxo(
    "id" TEXT,
    "vout" INTEGER,
    "address" TEXT,
    "amount" REAL,
    "is_spent" INTEGER DEFAULT 0,
    "block_height" INTEGER,
    PRIMARY KEY("id", "vout"),
    FOREIGN KEY("block_height") REFERENCES block("height")
    );`
	_, err = db.Exec(creatTable)
	return err
}

func initBlockTable(db *sql.DB) (err error) {
	creatTable := `
	CREATE TABLE IF NOT EXISTS block(
	"height" INTEGER PRIMARY KEY,
	"hash" TEXT
	);`
	_, err = db.Exec(creatTable)
	return err
}

func initTxTable(db *sql.DB) (err error) {
	creatTable := `
	CREATE TABLE IF NOT EXISTS tx(
	"txid" TEXT,
	"pre_txid" TEXT,
	"pre_vout" INTEGER,
	PRIMARY KEY("pre_txid", "pre_vout"),
	FOREIGN KEY("pre_txid", "pre_vout") REFERENCES utxo("id", "vout")
	);`
	_, err = db.Exec(creatTable)
	return err
}

func initContractTable(db *sql.DB) (err error) {
	creatTable := `
	CREATE TABLE IF NOT EXISTS contract(
    "txid" TEXT PRIMARY KEY,
    "contract_address" TEXT UNIQUE,
    "contract_action" TEXT
	);`
	_, err = db.Exec(creatTable)
	return err
}

func ClearTables(db *sql.DB) (err error) {
	if db == nil {
		return errors.New("db is nil")
	}
	if _, err = db.Exec("DELETE FROM contract"); err != nil {
		return err
	}
	if _, err = db.Exec("DELETE FROM tx"); err != nil {
		return err
	}
	if _, err = db.Exec("DELETE FROM utxo"); err != nil {
		return err
	}
	if _, err = db.Exec("DELETE FROM block"); err != nil {
		return err
	}
	return nil
}
