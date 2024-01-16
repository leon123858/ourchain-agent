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
	return dbClient
}

func tearDown(dbClient Client) {
	_, err := dbClient.Instance.Exec("DELETE FROM utxo")
	if err != nil {
		panic("Delete utxo in teardown failed")
	}
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
	err := dbClient.createUtxo(utxo{
		utxoSearchArgument: utxoSearchArgument{
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
	err := dbClient.createUtxoList([]utxo{
		{
			utxoSearchArgument: utxoSearchArgument{
				ID:      "test1",
				Address: "test",
			},
			Vout:   11,
			Amount: 20,
		},
		{
			utxoSearchArgument: utxoSearchArgument{
				ID:      "test2",
				Address: "test",
			},
			Vout:   12,
			Amount: 20,
		},
		{
			utxoSearchArgument: utxoSearchArgument{
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
	err := client.createUtxo(utxo{
		utxoSearchArgument: utxoSearchArgument{
			ID:      "test",
			Address: "test",
		},
		Vout:   10,
		Amount: 20,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = client.deleteUtxo("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_deleteUtxoList(t *testing.T) {
	client := setUp()
	err := client.createUtxoList([]utxo{
		{
			utxoSearchArgument: utxoSearchArgument{
				ID:      "test1",
				Address: "test",
			},
			Vout:   11,
			Amount: 20,
		},
		{
			utxoSearchArgument: utxoSearchArgument{
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
	result, _ := client.getUtxoList(utxoSearchArgument{})
	if len(result) != 2 {
		t.Fatal("createUtxoList failed")
	}
	err = client.deleteUtxoList([]string{"test1", "test2"})
	if err != nil {
		t.Fatal(err)
	}
	tearDown(client)
}

func TestClient_getUtxoList(t *testing.T) {
	client := setUp()
	err := client.createUtxo(utxo{
		utxoSearchArgument: utxoSearchArgument{
			ID:      "test",
			Address: "test",
		},
		Vout:   10,
		Amount: 20,
	})
	if err != nil {
		return
	}
	utxoList, err := client.getUtxoList(utxoSearchArgument{})
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
