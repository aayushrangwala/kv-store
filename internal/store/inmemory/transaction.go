package inmemory

import (
	"fmt"

	"cryptowatch/backend-go/internal/store"
)

var _ store.TransactionStack = &stack{}

// stack will implement the TransactionStack interface.
type stack struct {
	latest      *transaction
	size        int32
	globalStore map[string]string
}

// Start will start the transaction in memory by adding a transacation node to the latest and parent.
func (stack *stack) Start() {
	transctn := &transaction{
		localStore: make(map[string]string),
	}

	stack.size++

	if stack.latest == nil {
		stack.latest = transctn

		return
	}

	transctn.parent = stack.latest
	stack.latest = transctn
}

// Commit will persist all the data from transaction to global and parent transactions.
func (stack *stack) Commit() {
	ActiveTransaction := stack.latest
	if ActiveTransaction == nil {
		fmt.Println("ERROR: No Active transaction")

		return
	}

	for key, value := range ActiveTransaction.localStore {
		stack.globalStore[key] = value

		if ActiveTransaction.parent != nil {
			// update the parent transaction
			ActiveTransaction.parent.localStore[key] = value
		}
	}
}

// Abort will abort the current transaction and make the parent transaction active.
func (stack *stack) Abort() {
	if stack.latest == nil {
		fmt.Println("INFO: No Active transaction")

		return
	}

	stack.latest = stack.latest.parent
	stack.size--
}
