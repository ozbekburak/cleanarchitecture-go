package db

import "sync"

// InMemoryDB is an in-memory implementation of DB
type InMemoryDB struct {
	kv map[string]string
	mu *sync.Mutex
}

// NewInMemoryDB returns a new InMemoryDB instance
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		kv: make(map[string]string),
		mu: &sync.Mutex{},
	}
}

// Set sets a key-value pair in the DB
func (db *InMemoryDB) Set(key, value string) {
	db.mu.Lock()
	db.kv[key] = value
	db.mu.Unlock()
}

// Get gets a value from the in-memory DB with the given key
func (db *InMemoryDB) Get(key string) (string, string, bool) {
	db.mu.Lock()
	value, ok := db.kv[key]
	db.mu.Unlock()

	return key, value, ok
}
