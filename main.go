package main

import (
	"github.com/ozbekburak/cleanarchitecture-go/pkg/config"
	"github.com/ozbekburak/cleanarchitecture-go/pkg/db"
	"github.com/ozbekburak/cleanarchitecture-go/pkg/logger"
	"github.com/ozbekburak/cleanarchitecture-go/pkg/router"
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
