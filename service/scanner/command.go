package scanner

import "log"

type commandType string

type command struct {
	Type commandType
	args []interface{}
}

const (
	ADD_BLOCK      commandType = "ADD_BLOCK"
	MINUS_BLOCK    commandType = "MINUS_BLOCK"
	ADD_UTXO       commandType = "ADD_UTXO"
	MINUS_UTXO     commandType = "MINUS_UTXO"
	UPDATA_UTXO    commandType = "UPDATA_UTXO"
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
