package scanner

import (
	our_chain_rpc "github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/sqlite"
	"testing"
)

func initDB() *sqlite.Client {
	dbClient := sqlite.Client{}
	if sqlite.New(&dbClient) != nil {
		panic("New dbClient failed")
	}
	return &dbClient
}

func setUp() (*our_chain_rpc.Bitcoind, *sqlite.Client) {
	rpc := initChain()
	db := initDB()
	if err := sqlite.ClearTables(db.Instance); err != nil {
		return nil, nil
	}
	return rpc, db
}

func tearDown(db *sqlite.Client) {
	if db.Close() != nil {
		panic("Close dbClient failed")
	}
}

func Test_addBlocksCoder(t *testing.T) {
	rpc, db := setUp()
	defer tearDown(db)
	chain := localChain{
		Chain:  make([]block, 0),
		Length: 0,
		Client: db,
	}
	cmds, err := addBlocksCoder(&chain, rpc, 200)
	if err != nil {
		t.Fatal(err)
	}
	for _, cmd := range *cmds {
		cmd.Print()
	}
}

func Test_minusBlocksCoder(t *testing.T) {
	rpc, db := setUp()
	defer tearDown(db)
	chain := localChain{
		Chain:  make([]block, 0),
		Length: 0,
		Client: db,
	}
	cmds, err := minusBlocksCoder(&chain, rpc, 200)
	if err != nil {
		t.Fatal(err)
	}
	for _, cmd := range *cmds {
		cmd.Print()
	}
}
