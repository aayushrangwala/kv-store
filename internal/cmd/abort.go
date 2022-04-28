package cmd

import "cryptowatch/backend-go/internal/store"

var _ Executor = &abort{}

type abort struct {
	ds store.Store
}

func newAbortCommand(datastore store.Store) Executor {
	return &abort{ds: datastore}
}

func (cmd *abort) Execute() error {
	cmd.ds.Abort()

	return nil
}
