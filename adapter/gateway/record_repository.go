package gateway

import (
	"context"
	"os"
	"time"

	"github.com/ozbekburak/cleanarch-mongo-inmem/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordRepository struct {
	Conn *mongo.Client
}

// Filter gets filtered records from database and return filtered records
func (rr *RecordRepository) Filter(f domain.Filter) (domain.Records, error) {
	// We are checking formatting error in the usecase layer so we don't need to check here
	start, _ := f.FormatDate(f.StartDate)
	end, _ := f.FormatDate(f.EndDate)

	// Filtering our records by date range and counts using aggregation pipeline
	// We assumed that given date and count nums are greater/less or equal style
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gte": start,
					"$lte": end,
				},
			},
		},
		{
			"$addFields": bson.M{
				"totalCount": bson.M{
					"$sum": "$counts",
				},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gte": f.MinCount,
					"$lte": f.MaxCount,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": 1,
			},
		},
	}

	var records domain.Records

	collection := rr.Conn.Database(os.Getenv("database")).Collection(os.Getenv("collection"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Finding our records
	data, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return domain.Records{}, err
	}

	// Deserializing our records
	err = data.All(ctx, &records.Records)
	if err != nil {
		return domain.Records{}, err
	}

	return records, nil
}
