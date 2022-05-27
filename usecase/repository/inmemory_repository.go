package repository

import "github.com/ozbekburak/go-clean-api/domain"

// InMemoryRepository interface defines methods for create, read operations on in-memory db
type InMemoryRepository interface {
	Create(domain.KeyValue) (domain.KeyValue, error)
	Get(string) (string, string, bool)
}
