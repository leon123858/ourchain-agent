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
    "is_coinbase" INTEGER DEFAULT 0,
    "pre_txid" TEXT,
    "pre_vout" INTEGER,
    "block_height" INTEGER,
    PRIMARY KEY("id", "vout"),
    FOREIGN KEY("pre_txid", "pre_vout") REFERENCES utxo("id", "vout"),
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

func clearTables(db *sql.DB) (err error) {
	if db == nil {
		return errors.New("db is nil")
	}
	if _, err = db.Exec("DELETE FROM utxo"); err != nil {
		return err
	}
	if _, err = db.Exec("DELETE FROM block"); err != nil {
		return err
	}
	return nil
}
