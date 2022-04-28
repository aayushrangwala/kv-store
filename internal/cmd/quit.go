package cmd

import "cryptowatch/backend-go/internal/store"

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
	cmd.ds.Quit()

	return nil
}
