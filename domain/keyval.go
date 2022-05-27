package domain

// KeyValue represent our data type for in-memory storage
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
