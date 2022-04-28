package cmd

import "cryptowatch/backend-go/internal/store"

var _ Executor = &deleteCmd{}

type deleteCmd struct {
	key string
	ds  store.Store
}

func newDeleteCommand(datastore store.Store, key string) Executor {
	return &deleteCmd{
		key: key,
		ds:  datastore,
	}
}

func (cmd *deleteCmd) Execute() error {
	cmd.ds.Delete(cmd.key)

	return nil
}
