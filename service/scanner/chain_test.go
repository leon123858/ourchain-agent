package scanner

import (
	"testing"
)

func Test_localChain_SyncLength(t *testing.T) {
	chain := initChain()
	db := initDB()
	var err error
	curLocalChain := newChain(clientWrapper{
		ChainType: LOCAL,
		DB:        db,
	})
	curRemoteChain := newChain(clientWrapper{
		ChainType: REMOTE,
		RPC:       chain,
	})
	err = curLocalChain.InitChainStep()
	if err != nil {
		t.Fatal(err)
	}
	err = curRemoteChain.InitChainStep()
	if err != nil {
		t.Fatal(err)
	}
	err = curLocalChain.SyncLength(curRemoteChain)
	if err != nil {
		t.Fatal(err)
	}
}
