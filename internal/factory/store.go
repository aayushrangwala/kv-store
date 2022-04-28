package factory

import (
	"cryptowatch/backend-go/internal/store"
	"cryptowatch/backend-go/internal/store/inmemory"
)

type storeFactory struct{}

func NewStoreFactory() *storeFactory {
	return &storeFactory{}
}

func (factory *storeFactory) GetInmemoryStore() store.Store {
	return inmemory.NewStore()
}
