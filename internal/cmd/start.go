package cmd

import "cryptowatch/backend-go/internal/store"

var _ Executor = &start{}

type start struct {
	ds store.Store
}

func newStartCommand(datastore store.Store) Executor {
	return &start{
		ds: datastore,
	}
}

func (cmd *start) Execute() error {
	cmd.ds.Start()

	return nil
}
