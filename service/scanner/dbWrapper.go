package scanner

import "github.com/leon123858/go-aid/service/sqlite"

func operates(client *sqlite.Client) (err error) {
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
