package cmd

import (
	"errors"
	"fmt"

	"cryptowatch/backend-go/internal/store"
)

// CommandType is the ENUM to represent the supported type of commands.
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

// Executor is the interface to be implemented by all the commands, which represents an object which is executable.
type Executor interface {
	// Execute is the execution operation on the object which will be implemented in its own way depending on the command.
	Execute() error
}

// GetCommandExecutor is the constructor of the command type based on the arguments passed.
// It returns error when arguments are not as expected or the datastore used internally is nil.
func GetCommandExecutor(datastore store.Store, args ...string) (Executor, error) {
	if len(args) == 0 {
		return nil, errors.New("invalid argument. No operation passed")
	}

	if datastore == nil {
		return nil, errors.New("invalid argument. Nil datastore passed")
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
		if len(args) < 3 {
			return nil, errors.New("invalid argument. Need key and value")
		}

		return newWriteCommand(datastore, args[1], args[2]), nil
	case Read:
		if len(args) < 2 {
			return nil, errors.New("invalid argument. Need key")
		}

		return newReadCommand(datastore, args[1]), nil
	case Delete:
		if len(args) < 2 {
			return nil, errors.New("invalid argument. Need key")
		}

		return newDeleteCommand(datastore, args[1]), nil
	default:
		return nil, fmt.Errorf("unrecognised operation: %s", cmdType)
	}
}
