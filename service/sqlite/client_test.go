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
	if err := clearTables(dbClient.Instance); err != nil {
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

func TestClient_createUtxo(t *testing.T) {
	dbClient := setUp()
	err := dbClient.CreateUtxo(Utxo{
		UtxoSearchArgument: UtxoSearchArgument{
			ID:      "test",
			Address: "test",
		},
		Vout:   10,
		Amount: 20,
	})
	if err != nil {
		t.Fatal(err)
	}
	tearDown(dbClient)
}

func TestClient_createUtxoList(t *testing.T) {
	dbClient := setUp()
	err := dbClient.CreateUtxoList([]Utxo{
		{
			UtxoSearchArgument: UtxoSearchArgument{
				ID:      "test1",
				Address: "test",
			},
			Vout:   11,
			Amount: 20,
		},
		{
			UtxoSearchArgument: UtxoSearchArgument{
				ID:      "test2",
				Address: "test",
			},
			Vout:   12,
			Amount: 20,
		},
		{
			UtxoSearchArgument: UtxoSearchArgument{
				ID:      "test3",
				Address: "test",
			},
			Vout:   10,
			Amount: 20,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	tearDown(dbClient)
}

func TestClient_deleteUtxo(t *testing.T) {
	client := setUp()
	err := client.CreateUtxo(Utxo{
		UtxoSearchArgument: UtxoSearchArgument{
			ID:      "test",
			Address: "test",
		},
		Vout:   10,
		Amount: 20,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = client.DeleteUtxo("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_deleteUtxoList(t *testing.T) {
	client := setUp()
	err := client.CreateUtxoList([]Utxo{
		{
			UtxoSearchArgument: UtxoSearchArgument{
				ID:      "test1",
				Address: "test",
			},
			Vout:   11,
			Amount: 20,
		},
		{
			UtxoSearchArgument: UtxoSearchArgument{
				ID:      "test2",
				Address: "test",
			},
			Vout:   12,
			Amount: 20,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	result, _ := client.GetUtxoList(UtxoSearchArgument{})
	if len(result) != 2 {
		t.Fatal("createUtxoList failed")
	}
	err = client.DeleteUtxoList([]string{"test1", "test2"})
	if err != nil {
		t.Fatal(err)
	}
	tearDown(client)
}

func TestClient_getUtxoList(t *testing.T) {
	client := setUp()
	err := client.CreateUtxo(Utxo{
		UtxoSearchArgument: UtxoSearchArgument{
			ID:      "test",
			Address: "test",
		},
		Vout:   10,
		Amount: 20,
	})
	if err != nil {
		return
	}
	utxoList, err := client.GetUtxoList(UtxoSearchArgument{})
	if err != nil {
		t.Fatal(err)
	}
	// print utxoList
	for _, item := range utxoList {
		assert.Equal(t, item.ID, "test")
		assert.Equal(t, item.Address, "test")
		assert.Equal(t, item.Vout, 10)
		assert.Equal(t, item.Amount, 20.0)
	}
	tearDown(client)
}

func TestClient_getBlocks(t *testing.T) {
	client := setUp()
	err := client.CreateBlock(Block{
		Height: 1,
		Hash:   "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	err = client.CreateBlock(Block{
		Height: 1,
		Hash:   "test",
	})
	if err == nil {
		t.Fatal("createBlock should failed when same key")
	}
	err = client.CreateBlock(Block{
		Height: 2,
		Hash:   "test2",
	})
	if err != nil {
		t.Fatal(err)
	}
	blockList, err := client.GetBlocks(1, 0)
	if err != nil {
		t.Fatal(err)
	}
	tmpInt := uint64(2)
	assert.Equal(t, len(blockList), 1)
	assert.Equal(t, blockList[0].Height, tmpInt)
	assert.Equal(t, blockList[0].Hash, "test2")

	blockList, err = client.GetBlocks(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	tmpInt = uint64(1)
	assert.Equal(t, len(blockList), 1)
	assert.Equal(t, blockList[0].Height, tmpInt)
	assert.Equal(t, blockList[0].Hash, "test")

	err = client.DeleteBlock(1)
	if err != nil {
		t.Fatal(err)
	}

	blockList, err = client.GetBlocks(2, 0)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(blockList), 1)

	tearDown(client)
}
