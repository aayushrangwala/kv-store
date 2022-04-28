package inmemory

import (
	"fmt"

	"cryptowatch/backend-go/internal/store"
)

var _ store.TransactionStack = &stack{}

type stack struct {
	latest      *transaction
	size        int32
	limit       int32
	GlobalStore map[string]string
}

func (stack *stack) Start() {
	transctn := &transaction{
		localStore: make(map[string]string),
	}

	if stack.latest == nil {
		stack.latest = transctn

		return
	}

	transctn.parent = stack.latest
	stack.latest = transctn

	stack.size++
}

func (stack *stack) Quit() {
	if stack.latest == nil {
		fmt.Println("ERROR: No Active transaction")

		return
	}

	stack.latest = stack.latest.parent
	stack.size--
}

func (stack *stack) Commit() {
	ActiveTransaction := stack.latest
	if ActiveTransaction == nil {
		fmt.Println("INFO: No Active transaction")

		return
	}

	for key, value := range ActiveTransaction.localStore {
		stack.GlobalStore[key] = value

		if ActiveTransaction.parent != nil {
			// update the parent transaction
			ActiveTransaction.parent.localStore[key] = value
		}
	}
}

func (stack *stack) Abort() {
	if stack.latest == nil {
		return
	}

	for key := range stack.latest.localStore {
		delete(stack.latest.localStore, key)
	}
}
