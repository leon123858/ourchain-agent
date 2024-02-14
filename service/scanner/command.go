package scanner

import (
	"database/sql"
	our_chain_rpc "github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/sqlite"
	"log"
)

type commandType string
type rawCommandType string

type command struct {
	Type commandType
	args []interface{}
}

type rawCommand struct {
	Type rawCommandType
	Stmt *sql.Stmt
	item interface{}
}

const (
	RAW_ADD_TX         rawCommandType = "ADD_TX"
	RAW_DELETE_TX      rawCommandType = "DELETE_TX"
	RAW_ADD_BLOCK      rawCommandType = "ADD_BLOCK"
	RAW_DELETE_BLOCK   rawCommandType = "DELETE_BLOCK"
	RAW_ADD_UTXO       rawCommandType = "ADD_UTXO"
	RAW_UPDATE_UTXO    rawCommandType = "UPDATE_UTXO"
	RAW_DELETE_UTXO    rawCommandType = "DELETE_UTXO"
	RAW_ADD_PREUTXO    rawCommandType = "ADD_PREUTXO"
	RAW_DELETE_PREUTXO rawCommandType = "DELETE_PREUTXO"
)

const (
	ADD_TX         commandType = "ADD_TX"
	ADD_BLOCK      commandType = "ADD_BLOCK"
	ADD_UTXO       commandType = "ADD_UTXO"
	ADD_PREUTXO    commandType = "ADD_PREUTXO"
	REMOVE_PREUTXO commandType = "REMOVE_PREUTXO"
)

func newCommand(t commandType, args ...interface{}) *command {
	return &command{
		Type: t,
		args: args,
	}
}

func (c *command) Print() {
	// print command like: ADD_BLOCK ...args
	log.Println("Command: ", c.Type, c.args)
}

func (rc *rawCommand) Exec() (err error) {
	switch rc.Type {
	case RAW_ADD_BLOCK:
		_, err = sqlite.BlockCreateExec(rc.Stmt, rc.item.(sqlite.Block))
	case RAW_DELETE_BLOCK:
		_, err = sqlite.BlockDeleteExec(rc.Stmt, rc.item.(sqlite.Block))
	case RAW_ADD_UTXO:
		_, err = sqlite.UtxoCreateExec(rc.Stmt, rc.item.(sqlite.Utxo))
	case RAW_UPDATE_UTXO:
		_, err = sqlite.UtxoUpdateExec(rc.Stmt, rc.item.(sqlite.Utxo))
	case RAW_DELETE_UTXO:
		_, err = sqlite.UtxoDeleteExec(rc.Stmt, rc.item.(sqlite.Utxo))
	case RAW_ADD_PREUTXO:
		_, err = sqlite.TxCreateExec(rc.Stmt, rc.item.(sqlite.PreUtxo))
	case RAW_DELETE_PREUTXO:
		_, err = sqlite.TxDeleteExec(rc.Stmt, rc.item.(sqlite.PreUtxo))
	case RAW_ADD_TX:
		_, err = sqlite.ContractCreateExec(rc.Stmt, rc.item.(sqlite.Contract))
	case RAW_DELETE_TX:
		_, err = sqlite.ContractDeleteExec(rc.Stmt, rc.item.(sqlite.Contract))
	}
	return
}

func compileAdd(sqlTx *sql.Tx, commandList *[]command) (rawCommandList *[]rawCommand, err error) {
	realRawCommandList := make([]rawCommand, 0)
	rawCommandList = &realRawCommandList
	utxoCreatStmt, err := sqlite.UtxoCreatePrepare(sqlTx)
	if err != nil {
		return
	}
	utxoUpdateStmt, err := sqlite.UtxoUpdatePrepare(sqlTx)
	if err != nil {
		return
	}
	blockCreateStmt, err := sqlite.BlockCreatePrepare(sqlTx)
	if err != nil {
		return
	}
	txCreateStmt, err := sqlite.TxCreatePrepare(sqlTx)
	if err != nil {
		return
	}
	contractCreateStmt, err := sqlite.ContractCreatePrepare(sqlTx)
	if err != nil {
		return
	}
	for _, cmd := range *commandList {
		switch cmd.Type {
		case ADD_BLOCK:
			rawCmd := rawCommand{
				Type: RAW_ADD_BLOCK,
				Stmt: blockCreateStmt,
				item: sqlite.Block{
					Height: cmd.args[0].(uint64),
					Hash:   cmd.args[1].(string),
				},
			}
			*rawCommandList = append(*rawCommandList, rawCmd)
		case ADD_TX:
			rawCmd := rawCommand{
				Type: RAW_ADD_TX,
				Stmt: contractCreateStmt,
				item: sqlite.Contract{
					TxID:            cmd.args[0].(string),
					ContractAction:  cmd.args[1].(our_chain_rpc.ContractAction),
					ContractAddress: cmd.args[2].(string),
				},
			}
			*rawCommandList = append(*rawCommandList, rawCmd)
		case ADD_UTXO:
			rawCmd := rawCommand{
				Type: RAW_ADD_UTXO,
				Stmt: utxoCreatStmt,
				item: sqlite.Utxo{
					ID:          cmd.args[0].(string),
					Vout:        cmd.args[1].(int),
					Address:     cmd.args[2].(string),
					Amount:      cmd.args[3].(float64),
					IsSpent:     cmd.args[4].(bool),
					BlockHeight: cmd.args[5].(uint64),
				},
			}
			*rawCommandList = append(*rawCommandList, rawCmd)
		case ADD_PREUTXO:
			rawCmd := rawCommand{
				Type: RAW_ADD_PREUTXO,
				Stmt: txCreateStmt,
				item: sqlite.PreUtxo{
					TxID:    cmd.args[0].(string),
					PreTxID: cmd.args[1].(string),
					PreVout: cmd.args[2].(int),
				},
			}
			*rawCommandList = append(*rawCommandList, rawCmd)
			rawCmd = rawCommand{
				Type: RAW_UPDATE_UTXO,
				Stmt: utxoUpdateStmt,
				item: sqlite.Utxo{
					ID:      cmd.args[1].(string),
					Vout:    cmd.args[2].(int),
					IsSpent: true,
				},
			}
			*rawCommandList = append(*rawCommandList, rawCmd)
		}
	}
	return
}

func compileMinus(sqlClient *sqlite.Client, sqlTx *sql.Tx, commandList *[]command) (rawCommandList *[]rawCommand, err error) {
	realRawCommandList := make([]rawCommand, 0)
	rawCommandList = &realRawCommandList
	utxoUpdateStmt, err := sqlite.UtxoUpdatePrepare(sqlTx)
	if err != nil {
		return
	}
	utxoDeleteStmt, err := sqlite.UtxoDeletePrepare(sqlTx)
	if err != nil {
		return
	}
	blockDeleteStmt, err := sqlite.BlockDeletePrepare(sqlTx)
	if err != nil {
		return
	}
	txDeleteStmt, err := sqlite.TxDeletePrepare(sqlTx)
	if err != nil {
		return
	}
	contractDeleteStmt, err := sqlite.ContractDeletePrepare(sqlTx)
	if err != nil {
		return
	}
	for _, cmd := range *commandList {
		switch cmd.Type {
		case REMOVE_PREUTXO:
			removeHeight := cmd.args[0].(uint64)
			// search txs by height
			txList, err := sqlClient.GetBlockTxList(removeHeight)
			if err != nil {
				return nil, err
			}
			// search preOut by txs
			preUtxoList := make([]sqlite.PreUtxo, 0)
			for _, tx := range *txList {
				preUtxo, err := sqlClient.GetPreUtxo(tx)
				if err != nil {
					return nil, err
				}
				preUtxoList = append(preUtxoList, *preUtxo...)
			}
			// update utxo by preOut
			for _, preUtxo := range preUtxoList {
				rawCmd := rawCommand{
					Type: RAW_UPDATE_UTXO,
					Stmt: utxoUpdateStmt,
					item: sqlite.Utxo{
						ID:      preUtxo.PreTxID,
						Vout:    preUtxo.PreVout,
						IsSpent: false,
					},
				}
				*rawCommandList = append(*rawCommandList, rawCmd)
			}
			// delete preOut
			for _, tx := range *txList {
				*rawCommandList = append(*rawCommandList, rawCommand{
					Type: RAW_DELETE_PREUTXO,
					Stmt: txDeleteStmt,
					item: sqlite.PreUtxo{
						TxID: tx,
					},
				})
				*rawCommandList = append(*rawCommandList, rawCommand{
					Type: RAW_DELETE_TX,
					Stmt: contractDeleteStmt,
					item: sqlite.Contract{
						TxID: tx,
					},
				})
			}
			*rawCommandList = append(*rawCommandList, rawCommand{
				Type: RAW_DELETE_UTXO,
				Stmt: utxoDeleteStmt,
				item: sqlite.Utxo{
					BlockHeight: removeHeight,
				},
			})
			*rawCommandList = append(*rawCommandList, rawCommand{
				Type: RAW_DELETE_BLOCK,
				Stmt: blockDeleteStmt,
				item: sqlite.Block{
					Height: removeHeight,
				},
			})
		}
	}
	return
}
