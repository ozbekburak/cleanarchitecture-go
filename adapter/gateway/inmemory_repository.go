package gateway

import (
	"github.com/ozbekburak/go-clean-api/domain"
	"github.com/ozbekburak/go-clean-api/pkg/db"
)

type InMemoryRepository struct {
	Conn *db.InMemoryDB
}

// Create send key-value pair to InMemoryDB
func (imr *InMemoryRepository) Create(kv domain.KeyValue) (domain.KeyValue, error) {
	imr.Conn.Set(kv.Key, kv.Value)
	return kv, nil
}

// Get returns a value from InMemoryDB with the given key
func (imr *InMemoryRepository) Get(key string) (string, string, bool) {
	return imr.Conn.Get(key)
}
