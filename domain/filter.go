package domain

import (
	"time"
)

// Filter represents our request payload for filtering records
type Filter struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

// FormatDate formats the date string to time.Time
func (f Filter) FormatDate(date string) (time.Time, error) {
	layout := "2006-01-02"

	t, err := time.Parse(layout, date)
	return t, err
}
