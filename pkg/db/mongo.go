package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB interface {
	Conn() (conn *mongo.Client, err error)
}

// MongoDB is a struct that implements DB interface
type MongoDB struct{}

// NewMongoDB returns a new MongoDB instance
func NewMongoDB() DB {
	return &MongoDB{}
}

// Conn returns a connection to the MongoDB instance
func (d *MongoDB) Conn() (*mongo.Client, error) {
	conn, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("mongo_uri")))
	if err != nil {
		return nil, err
	}

	// Ping the primary
	if err := conn.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return conn, nil
}
