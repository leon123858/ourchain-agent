package sqlite

import (
	"database/sql"
	"errors"
)

func initTables(db *sql.DB) (err error) {
	if db == nil {
		return errors.New("db is nil")
	}
	creatTable := `
    CREATE TABLE IF NOT EXISTS utxo(
    "id" TEXT PRIMARY KEY,
    "vout" INTEGER,
    "address" TEXT,
    "amount" REAL
    );`
	_, err = db.Exec(creatTable)
	return err
}
