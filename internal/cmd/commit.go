package cmd

import "cryptowatch/backend-go/internal/store"

var _ Executor = &commit{}

type commit struct {
	ds store.Store
}

func newCommitCommand(datastore store.Store) Executor {
	return &commit{ds: datastore}
}

func (cmd *commit) Execute() error {
	cmd.ds.Commit()

	return nil
}
