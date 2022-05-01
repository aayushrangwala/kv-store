package cmd

import (
	"fmt"
	"os"

	"cryptowatch/backend-go/internal/store"
)

var _ Executor = &quit{}

type quit struct {
	ds store.Store
}

func newQuitCommand(datastore store.Store) Executor {
	return &quit{
		ds: datastore,
	}
}

func (cmd *quit) Execute() error {
	fmt.Println("Execution Stopped. Thank you for using REPL!!")
	os.Exit(0)

	return nil
}
