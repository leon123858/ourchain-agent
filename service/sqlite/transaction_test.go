package sqlite

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBeginTx(t *testing.T) {
	client := setUp()
	tx, err := BeginTx(&client)
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil {
		t.Fatal("tx is nil")
	}
	err = RollbackTx(tx)
	if err != nil {
		return
	}
	err = CommitTx(tx)
	if err != nil {
		return
	}
	tearDown(client)
}

func TestBlockCreate(t *testing.T) {
	client := setUp()

	tx, err := BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err := BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	item := Block{
		Height: 1,
		Hash:   "hash",
	}
	_, err = BlockCreateExec(stmt, item)
	assert.Equal(t, err, nil)
	blocks, err := client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 0)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	blocks, err = client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 1)

	tearDown(client)
}

func TestBlockCreateList(t *testing.T) {
	client := setUp()

	tx, err := BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err := BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	for i := 1; i < 10; i++ {
		item := Block{
			Height: uint64(i),
			Hash:   "hash",
		}
		_, err = BlockCreateExec(stmt, item)
		assert.Equal(t, err, nil)
	}
	blocks, err := client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 0)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	blocks, err = client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 1)
	assert.Equal(t, blocks[0].Height, uint64(9))

	tearDown(client)
}

func TestBlockDelete(t *testing.T) {
	client := setUp()
	// insert
	tx, err := BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err := BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	for i := 1; i < 10; i++ {
		item := Block{
			Height: uint64(i),
			Hash:   "hash",
		}
		_, err = BlockCreateExec(stmt, item)
		assert.Equal(t, err, nil)
	}
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	// delete
	tx, err = BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err = BlockDeletePrepare(tx)
	assert.Equal(t, err, nil)
	item := Block{
		Height: 9,
	}
	_, err = BlockDeleteExec(stmt, item)
	assert.Equal(t, err, nil)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	// check
	blocks, err := client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 1)
	assert.Equal(t, blocks[0].Height, uint64(8))
	tearDown(client)
}

func TestUtxoConstrain(t *testing.T) {
	client := setUp()
	// block_height constrain
	tx, err := BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err := BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	stmt, err = UtxoCreatePrepare(tx)
	assert.Equal(t, err, nil)
	utxo := Utxo{
		ID:          "txhash",
		Vout:        0,
		BlockHeight: 1,
		Address:     "address",
		Amount:      1,
		IsSpent:     false,
	}
	_, err = UtxoCreateExec(stmt, utxo)
	assert.Equal(t, err.Error(), "FOREIGN KEY constraint failed")
	err = CommitTx(tx)
	// pre_txid, pre_vout constrain
	tx, err = BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err = BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	block := Block{
		Height: 1,
		Hash:   "hash",
	}
	_, err = BlockCreateExec(stmt, block)
	assert.Equal(t, err, nil)
	stmt, err = UtxoCreatePrepare(tx)
	assert.Equal(t, err, nil)
	utxo = Utxo{
		ID:          "txhash",
		Vout:        0,
		BlockHeight: 1,
		Address:     "address",
		Amount:      1,
		IsSpent:     false,
	}
	_, err = UtxoCreateExec(stmt, utxo)
	stmt, err = TxCreatePrepare(tx)
	assert.Equal(t, err, nil)
	txItem := PreUtxo{
		TxID:    "txhash2",
		PreTxID: "pre_txhash",
		PreVout: 0,
	}
	_, err = TxCreateExec(stmt, txItem)
	assert.Equal(t, err.Error(), "FOREIGN KEY constraint failed")
	txItem.PreTxID = "txhash"
	_, err = TxCreateExec(stmt, txItem)
	assert.Equal(t, err, nil)
	err = CommitTx(tx)
	// primary key constrain
	tx, err = BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err = UtxoCreatePrepare(tx)
	assert.Equal(t, err, nil)
	utxo = Utxo{
		ID:          "txhash",
		Vout:        0,
		BlockHeight: 1,
		Address:     "address",
		Amount:      1,
		IsSpent:     false,
	}
	_, err = UtxoCreateExec(stmt, utxo)
	assert.Equal(t, err.Error(), "UNIQUE constraint failed: utxo.id, utxo.vout")
	err = CommitTx(tx)

	tearDown(client)
}

func TestUtxo(t *testing.T) {
	client := setUp()
	// insert
	tx, err := BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err := BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	item := Block{
		Height: uint64(1),
		Hash:   "hash",
	}
	_, err = BlockCreateExec(stmt, item)
	assert.Equal(t, err, nil)
	stmt, err = UtxoCreatePrepare(tx)
	assert.Equal(t, err, nil)
	utxo := Utxo{
		ID:          "txhash",
		Vout:        0,
		BlockHeight: 1,
		Address:     "address",
		Amount:      1,
		IsSpent:     false,
	}
	_, err = UtxoCreateExec(stmt, utxo)
	assert.Equal(t, err, nil)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	utxoList, err := client.GetAddressUtxo("address", 1)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*utxoList), 1)
	assert.Equal(t, (*utxoList)[0].ID, "txhash")
	assert.Equal(t, (*utxoList)[0].Vout, 0)
	assert.Equal(t, (*utxoList)[0].BlockHeight, uint64(1))
	assert.Equal(t, (*utxoList)[0].Address, "address")
	assert.Equal(t, (*utxoList)[0].Amount, float64(1))
	assert.Equal(t, (*utxoList)[0].IsSpent, false)
	// update
	tx, err = BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err = UtxoUpdatePrepare(tx)
	assert.Equal(t, err, nil)
	utxo = Utxo{
		ID:      "txhash",
		Vout:    0,
		IsSpent: true,
	}
	_, err = UtxoUpdateExec(stmt, utxo)
	assert.Equal(t, err, nil)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	utxoList, err = client.GetAddressUtxo("address", 1)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*utxoList), 0)
	// delete
	tx, err = BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err = UtxoDeletePrepare(tx)
	assert.Equal(t, err, nil)
	utxo = Utxo{
		ID:   "txhash",
		Vout: 0,
	}
	_, err = UtxoDeleteExec(stmt, utxo)
	assert.Equal(t, err, nil)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	utxoList, err = client.GetAddressUtxo("address", 1)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*utxoList), 0)

	tearDown(client)
}

func TestPreUtxo(t *testing.T) {
	client := setUp()

	tx, err := BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err := BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	block := Block{
		Height: 1,
		Hash:   "hash",
	}
	_, err = BlockCreateExec(stmt, block)
	assert.Equal(t, err, nil)
	stmt, err = UtxoCreatePrepare(tx)
	assert.Equal(t, err, nil)
	utxo := Utxo{
		ID:          "txhash",
		Vout:        0,
		BlockHeight: 1,
		Address:     "address",
		Amount:      1,
		IsSpent:     false,
	}
	_, err = UtxoCreateExec(stmt, utxo)
	assert.Equal(t, err, nil)
	stmt, err = TxCreatePrepare(tx)
	assert.Equal(t, err, nil)
	txItem := PreUtxo{
		TxID:    "txhash2",
		PreTxID: "txhash",
		PreVout: 0,
	}
	_, err = TxCreateExec(stmt, txItem)
	assert.Equal(t, err, nil)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	// check
	txItems, err := client.GetPreUtxo("txhash2")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*txItems), 1)
	assert.Equal(t, (*txItems)[0].TxID, "txhash2")
	assert.Equal(t, (*txItems)[0].PreTxID, "txhash")
	assert.Equal(t, (*txItems)[0].PreVout, 0)

	// delete
	tx, err = BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err = TxDeletePrepare(tx)
	assert.Equal(t, err, nil)
	txItem = PreUtxo{
		TxID: "txhash2",
	}
	_, err = TxDeleteExec(stmt, txItem)
	assert.Equal(t, err, nil)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	// check
	txItems, err = client.GetPreUtxo("txhash2")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*txItems), 0)

	tearDown(client)
}
