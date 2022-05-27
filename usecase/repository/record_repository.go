package repository

import "github.com/ozbekburak/go-clean-api/domain"

// RecordRepository interface defines methods for filter operations on Record
type RecordRepository interface {
	Filter(domain.Filter) (domain.Records, error)
}
