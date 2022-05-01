package inmemory

import (
	"sync"

	"cryptowatch/backend-go/internal/store"
)

var _ store.Store = &inMemory{}

// inMemory is the concrete type implementing store interface in a synchronized way.
type inMemory struct {
	transactionStack *stack
	mu               sync.RWMutex
}

// transaction represents the inmemory transaction object having the link to its parent transaction if any.
type transaction struct {
	localStore map[string]string
	parent     *transaction
}

// NewStore is the constructor for the inmemory data store.
func NewStore() store.Store {
	return &inMemory{
		transactionStack: &stack{
			globalStore: make(map[string]string),
		},
	}
}

// Get value of key from Store
func (store *inMemory) Get(key string) string {
	store.mu.RLock()
	defer store.mu.RUnlock()

	st := store.transactionStack.globalStore

	activeTransaction := store.transactionStack.latest
	if activeTransaction != nil {
		st = activeTransaction.localStore
	}

	if val, present := st[key]; present {
		return val
	}

	return "not set"
}

// Set key to value
func (store *inMemory) Set(key, value string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	// Get key:value store from active transaction
	activeTransaction := store.transactionStack.latest

	if activeTransaction == nil {
		store.transactionStack.globalStore[key] = value

		return
	}

	activeTransaction.localStore[key] = value
}

// Delete value from Store
func (store *inMemory) Delete(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	activeTransaction := store.transactionStack.latest
	if activeTransaction == nil {
		delete(store.transactionStack.globalStore, key)

		return
	}

	delete(activeTransaction.localStore, key)
}

// Start is a synchronized wrapper of transaction Start.
func (store *inMemory) Start() {
	store.mu.RLock()
	defer store.mu.RUnlock()

	store.transactionStack.Start()
}

// Commit is a synchronized wrapper of transaction Commit.
func (store *inMemory) Commit() {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.transactionStack.Commit()
}

// Abort is a synchronized wrapper of transaction Abort.
func (store *inMemory) Abort() {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.transactionStack.Abort()
}
