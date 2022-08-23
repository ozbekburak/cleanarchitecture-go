package main

import (
	"context"
	"net/http"

	"github.com/ozbekburak/cleanarch-mongo-inmem/pkg/config"
	"github.com/ozbekburak/cleanarch-mongo-inmem/pkg/db"
	"github.com/ozbekburak/cleanarch-mongo-inmem/pkg/logger"
	"github.com/ozbekburak/cleanarch-mongo-inmem/pkg/router"
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

	router := router.Initialize(conn)

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Errorf("Error when starting server %s", err)
	}
}
