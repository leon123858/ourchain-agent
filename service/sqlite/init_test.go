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
	if err = New(&dbClient); err != nil {
		t.Fatal("New dbClient failed")
	}
}
