package factory

import (
	"testing"

	"github.com/onsi/gomega"

	"cryptowatch/backend-go/internal/store/inmemory"
)

func TestNewStoreFactory(t *testing.T) {
	g := gomega.NewWithT(t)

	g.Expect(NewStoreFactory()).To(gomega.Equal(&storeFactory{}))
}

func TestStoreFactory_GetInmemoryStore(t *testing.T) {
	g := gomega.NewWithT(t)

	g.Expect(NewStoreFactory().GetInmemoryStore()).To(gomega.Equal(inmemory.NewStore()))
}
