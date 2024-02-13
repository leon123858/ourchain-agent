package sqlite

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func setUp() Client {
	dbClient := Client{}
	if New(&dbClient) != nil {
		panic("New dbClient failed")
	}
	if err := ClearTables(dbClient.Instance); err != nil {
		println(err.Error())
		panic("Clear tables failed")
	}
	return dbClient
}

func tearDown(dbClient Client) {
	if dbClient.Close() != nil {
		panic("Close dbClient failed")
	}
}

func TestNew(t *testing.T) {
	var err error
	dbClient := Client{}
	if New(&dbClient) != nil {
		t.Fatal("New dbClient failed")
	}
	if dbClient.Instance == nil {
		t.Fatal("dbClient.Instance is nil")
	}
	err = dbClient.Instance.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Close(t *testing.T) {
	dbClient := Client{}
	if New(&dbClient) != nil {
		t.Fatal("New dbClient failed")
	}
	if dbClient.Instance == nil {
		t.Fatal("dbClient.Instance is nil")
	}
	err := dbClient.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetFirstBlockInfo(t *testing.T) {
	client := setUp()

	var err error
	// get first block info and get error because of no data
	blocks, err := client.GetFirstBlockInfo()
	if err != nil {
		t.Fatal("should not error when no data")
	}
	assert.Equal(t, len(blocks), 0)
	// insert data
	_, err = client.Instance.Exec("INSERT INTO block(height, hash) VALUES(?, ?)", 1, "hash1")
	assert.Equal(t, err, nil)
	// get first block info and get data now
	blocks, err = client.GetFirstBlockInfo()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, blocks[0].Height, uint64(1))
	assert.Equal(t, blocks[0].Hash, "hash1")
	// insert more data
	_, err = client.Instance.Exec("INSERT INTO block(height, hash) VALUES(?, ?)", 2, "hash2")
	assert.Equal(t, err, nil)
	// get first block info and get data now
	blocks, err = client.GetFirstBlockInfo()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, blocks[0].Height, uint64(2))
	assert.Equal(t, blocks[0].Hash, "hash2")

	tearDown(client)
}

func TestClient_GetPreUtxo(t *testing.T) {
	client := setUp()

	var err error
	// get first block info and get error because of no data
	utxos, err := client.GetPreUtxo("emptyTxID")
	if err != nil {
		t.Fatal("should not error when no data")
	}
	assert.Equal(t, len(*utxos), 0)
	// insert data

	tearDown(client)
}
