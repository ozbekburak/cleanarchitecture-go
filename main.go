package main

import (
	"context"
	"net/http"

	"github.com/ozbekburak/go-clean-api/pkg/config"
	"github.com/ozbekburak/go-clean-api/pkg/db"
	"github.com/ozbekburak/go-clean-api/pkg/logger"
	"github.com/ozbekburak/go-clean-api/pkg/router"
)

func main() {
	err := config.Load()
	if err != nil {
		logger.Errorf("Loading config error: %s", err)
		return
	}

	mongo := db.NewMongoDB()

	conn, err := mongo.Conn()
	if err != nil {
		logger.Errorf("Connecting to db error: %s", err)
		return
	}
	defer func() {
		if err = conn.Disconnect(context.TODO()); err != nil {
			logger.Errorf("Disconnecting from db error: %s", err)
			return
		}
	}()

	inMemoryDB := db.NewInMemoryDB()

	router := router.Initialize(conn, inMemoryDB)

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Errorf("Error when starting server %s", err)
	}
}
