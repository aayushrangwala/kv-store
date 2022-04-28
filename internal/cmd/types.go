package cmd

import (
	"cryptowatch/backend-go/internal/store"
	"errors"
	"fmt"
)

type CommandType string

const (
	Read = "READ"

	Write = "WRITE"

	Start = "START"

	Commit = "COMMIT"

	Delete = "DELETE"

	Abort = "ABORT"

	Quit = "QUIT"
)

type Executor interface {
	Execute() error
}

func GetCommandExecutor(datastore store.Store, args ...string) (Executor, error) {
	if len(args) == 0 {
		return nil, errors.New("invalid argument. No operation passed")
	}

	cmdType := CommandType(args[0])

	switch cmdType {
	case Start:
		return newStartCommand(datastore), nil
	case Abort:
		return newAbortCommand(datastore), nil
	case Commit:
		return newCommitCommand(datastore), nil
	case Quit:
		return newQuitCommand(datastore), nil
	case Write:
		return newWriteCommand(datastore, args[1], args[2]), nil
	case Read:
		return newReadCommand(datastore, args[1]), nil
	case Delete:
		return newDeleteCommand(datastore, args[1]), nil
	default:
		return nil, fmt.Errorf("unrecognised operation: %s", cmdType)
	}
}
