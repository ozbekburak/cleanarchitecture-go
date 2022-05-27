package domain

import "time"

// Records contains our array of Record struct
type Records struct {
	Records []Record `json:"records"`
}

// Record represents our response payload for filtered records
type Record struct {
	Key        string    `bson:"key" json:"key"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	TotalCount int32     `bson:"totalCount" json:"totalCount"`
}
