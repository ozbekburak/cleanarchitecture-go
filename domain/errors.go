package domain

import "errors"

var (
	ErrMaxCountLessThanMin  = errors.New("Min count must be less than max count")
	ErrStartDateLessThanEnd = errors.New("Start date must be before end date")
	ErrEmptyKey             = errors.New("Key cannot be empty")
)
