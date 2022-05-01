package cmd

import (
	"cryptowatch/backend-go/internal/store"
)

var _ Executor = &write{}

type write struct {
	key   string
	value string
	ds    store.Store
}

func newWriteCommand(datastore store.Store, key, value string) Executor {
	return &write{
		key:   key,
		value: value,
		ds:    datastore,
	}
}

func (cmd *write) Execute() error {
	cmd.ds.Set(cmd.key, cmd.value)

	return nil
}
