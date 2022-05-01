package factory

import (
	"cryptowatch/backend-go/internal/store"
	"cryptowatch/backend-go/internal/store/inmemory"
)

// storeFactory implements the factory pattern in go to return the desired type of
type storeFactory struct{}

// NewStoreFactory constructor for the store factory.
func NewStoreFactory() *storeFactory {
	return &storeFactory{}
}

// GetInmemoryStore is the factory wrapper to return inmemory type of datastore.
func (factory *storeFactory) GetInmemoryStore() store.Store {
	return inmemory.NewStore()
}
