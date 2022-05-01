package cmd

import (
	"errors"
	"testing"

	"github.com/onsi/gomega"

	"cryptowatch/backend-go/internal/factory"
	"cryptowatch/backend-go/internal/store"
)

func TestGetCommandExecutor(t *testing.T) {
	g := gomega.NewWithT(t)

	datastore := factory.NewStoreFactory().GetInmemoryStore()

	tests := []struct {
		desc          string
		datastore     store.Store
		args          []string
		expectedCmd   Executor
		expectedError error
	}{
		{
			desc:          "Should return error when no arg passed",
			expectedError: errors.New("invalid argument. No operation passed"),
		},
		{
			desc:          "Should return error when nil datastore passed",
			args:          []string{Start},
			expectedError: errors.New("invalid argument. Nil datastore passed"),
		},
		{
			desc:          "Should return error when invalid arg passed",
			args:          []string{"invalid-cmd"},
			datastore:     datastore,
			expectedError: errors.New("unrecognised operation: invalid-cmd"),
		},
		{
			desc:        "Should return start command",
			args:        []string{Start},
			datastore:   datastore,
			expectedCmd: newStartCommand(datastore),
		},
		{
			desc:        "Should return start command",
			args:        []string{Abort},
			datastore:   datastore,
			expectedCmd: newAbortCommand(datastore),
		},
		{
			desc:        "Should return start command",
			args:        []string{Commit},
			datastore:   datastore,
			expectedCmd: newCommitCommand(datastore),
		},
		{
			desc:        "Should return start command",
			args:        []string{Quit},
			datastore:   datastore,
			expectedCmd: newQuitCommand(datastore),
		},
		{
			desc:          "Should return error with invalid args for write command",
			args:          []string{Write, "a"},
			datastore:     datastore,
			expectedError: errors.New("invalid argument. Need key and value"),
		},
		{
			desc:          "Should return error with invalid args for read command",
			args:          []string{Read},
			datastore:     datastore,
			expectedError: errors.New("invalid argument. Need key"),
		},
		{
			desc:          "Should return error with invalid args for delete command",
			args:          []string{Delete},
			datastore:     datastore,
			expectedError: errors.New("invalid argument. Need key"),
		},
		{
			desc:        "Should return write cmd with args",
			args:        []string{Write, "a", "value"},
			datastore:   datastore,
			expectedCmd: newWriteCommand(datastore, "a", "value"),
		},
		{
			desc:        "Should return read cmd with args",
			args:        []string{Read, "a"},
			datastore:   datastore,
			expectedCmd: newReadCommand(datastore, "a"),
		},
		{
			desc:        "Should return delete cmd with args",
			args:        []string{Delete, "a"},
			datastore:   datastore,
			expectedCmd: newDeleteCommand(datastore, "a"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			cmd, err := GetCommandExecutor(test.datastore, test.args...)

			if test.expectedError != nil {
				g.Expect(err).To(gomega.Equal(test.expectedError))
				g.Expect(cmd).To(gomega.BeNil())

				return
			}

			g.Expect(err).To(gomega.BeNil())
			g.Expect(cmd).To(gomega.Equal(test.expectedCmd))
		})
	}
}
