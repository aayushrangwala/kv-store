package inmemory

import (
	"sync"

	"cryptowatch/backend-go/internal/store"
)

var _ store.Store = &inMemory{}

type inMemory struct {
	transactionStack *stack
	mu               sync.RWMutex
}

type transaction struct {
	localStore map[string]string
	parent     *transaction
}

func NewStore() store.Store {
	return &inMemory{
		transactionStack: &stack{},
	}
}

// Get value of key from Store
func (store *inMemory) Get(key string) string {
	st := store.transactionStack.GlobalStore

	ActiveTransaction := store.transactionStack.latest
	if ActiveTransaction != nil {
		st = ActiveTransaction.localStore
	}

	if val, present := st[key]; present {
		return val
	}

	return "not set"
}

// Set key to value
func (store *inMemory) Set(key, value string) {
	// Get key:value store from active transaction
	ActiveTransaction := store.transactionStack.latest

	if ActiveTransaction == nil {
		store.transactionStack.GlobalStore[key] = value

		return
	}

	ActiveTransaction.localStore[key] = value
}

// Delete value from Store
func (store *inMemory) Delete(key string) {
	ActiveTransaction := store.transactionStack.latest
	if ActiveTransaction == nil {
		delete(store.transactionStack.GlobalStore, key)

		return
	}

	delete(ActiveTransaction.localStore, key)
}

func (store *inMemory) Start() {
	store.transactionStack.Start()
}

func (store *inMemory) Commit() {
	store.transactionStack.Commit()
}

func (store *inMemory) Abort() {
	store.transactionStack.Abort()
}

func (store *inMemory) Quit() {
	store.transactionStack.Quit()
}
