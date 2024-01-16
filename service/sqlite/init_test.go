package sqlite

import (
	"testing"
)

func Test_initTables(t *testing.T) {
	var err error
	dbClient := Client{}
	err = initTables(dbClient.Instance)
	if err == nil {
		t.Fatal("initTables should failed")
	}
	if New(&dbClient) != nil {
		t.Fatal("New dbClient failed")
	}
}
