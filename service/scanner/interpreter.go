package scanner

type CommandType string

const (
	ADD    CommandType = "add"
	MINUS  CommandType = "minus"
	UPDATE CommandType = "update"
)

type command struct {
	Type CommandType
	Args []string
}

func interpretCommand(commandList []command) (err error) {
	return
}
