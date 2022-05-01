package store

// TransactionStack represents the stack where the transactions will be stored.
type TransactionStack interface {
	// Start will start a new transaction, it will be either a first transaction or a child of a previous one.
	Start()

	// Commit will persist the key value pair to all the transactions in the stack along with the global store.
	Commit()

	// Abort will gracefully abort the current transaction.
	Abort()
}

// Store represents the actual datastore to store the key value pairs of this K-V store.
// This interface currently is being inmplemented by inMemory type of store,
// but can be extended to be implemented by any type of store.
type Store interface {
	// TransactionStack It will make sure that store also inherits the transaction stack capabilities.
	TransactionStack

	// Get returns the value of the corresponding key from the kv store and 'not set' if the key is not present.
	Get(key string) string

	// Set sets the value against the passed key.
	// If the transaction is going on, it will store to that otherwise will store to the global store.
	Set(key, value string)

	// Delete will delete the key-value either from the current transaction or from the global store.
	Delete(key string)
}
