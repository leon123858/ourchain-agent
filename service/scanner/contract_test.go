package scanner

import (
	"github.com/leon123858/go-aid/service/sqlite"
	"log"
	"testing"
)

func TestListContract(t *testing.T) {
	chain := initChain()
	db := sqlite.Client{}
	if sqlite.New(&db) != nil {
		log.Fatal("sqlite init failed")
	}
	list, err := ListContract(chain, &db, "undefined")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("List Contract: %+v", list)
}
