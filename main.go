package main

import (
	"github.com/ozbekburak/cleanarch-mongo-inmem/pkg/config"
	"github.com/ozbekburak/cleanarch-mongo-inmem/pkg/db"
	"github.com/ozbekburak/cleanarch-mongo-inmem/pkg/logger"
	"github.com/ozbekburak/cleanarch-mongo-inmem/pkg/router"
)

func main() {
	err := config.Load()
	if err != nil {
		logger.Errorf("Failed to load config error: %v", err)
		return
	}

	db := db.NewPostgres()

	conn, err := db.Conn()
	if err != nil {
		logger.Errorf("Failed to connect DB error: %v", err)
		return
	}

	defer conn.Close()

	router.Initialize(conn)
}
