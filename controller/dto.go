package controller

import (
	OurChain "github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/sqlite"
)

type RepositoryDTO struct {
	Chain    *OurChain.Bitcoind
	Database *sqlite.Client
}
