package interactor

import (
	"github.com/ozbekburak/cleanarch-mongo-inmem/domain"
	"github.com/ozbekburak/cleanarch-mongo-inmem/usecase/repository"
)

// InMemoryInteractor is the struct that implements the InMemoryRepository interface
type InMemoryInteractor struct {
	InMemoryRepository repository.InMemoryRepository
}

// Create send key-value pair to InMemoryDB
func (imr *InMemoryInteractor) Create(kv domain.KeyValue) (domain.KeyValue, error) {
	if kv.Key == "" {
		return domain.KeyValue{}, domain.ErrEmptyKey
	}
	return imr.InMemoryRepository.Create(kv)
}

// Get returns a value from InMemoryDB with the given key
func (imr *InMemoryInteractor) Get(key string) (string, string, bool) {
	return imr.InMemoryRepository.Get(key)
}
