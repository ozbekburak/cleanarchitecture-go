package repository

import "github.com/ozbekburak/cleanarch-mongo-inmem/domain"

// InMemoryRepository interface defines methods for create, read operations on in-memory db
type InMemoryRepository interface {
	Create(domain.KeyValue) (domain.KeyValue, error)
	Get(string) (string, string, bool)
}
