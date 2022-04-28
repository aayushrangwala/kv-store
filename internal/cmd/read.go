package cmd

import (
	"fmt"

	"cryptowatch/backend-go/internal/store"
)

var _ Executor = &read{}

type read struct {
	key string
	ds  store.Store
}

func newReadCommand(datastore store.Store, key string) Executor {
	return &read{
		key: key,
		ds:  datastore,
	}
}

func (cmd *read) Execute() error {
	fmt.Println(cmd.ds.Get(cmd.key))

	return nil
}
