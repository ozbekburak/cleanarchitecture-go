package repository

import "github.com/ozbekburak/cleanarch-mongo-inmem/domain"

// RecordRepository interface defines methods for filter operations on Record
type RecordRepository interface {
	Filter(domain.Filter) (domain.Records, error)
}
