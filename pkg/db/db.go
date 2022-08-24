package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DB interface {
	Conn() (conn *sql.DB, err error)
}

// PostgresDB represents Postgresql database connection
type PostgresDB struct{}

// NewDB returns a new DB
func NewPostgres() DB {
	return &PostgresDB{}
}

// Guarantee that PostgresDB implements DB
var _ DB = &PostgresDB{}

func (d *PostgresDB) Conn() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	conn, err := sql.Open("postgres", dsn) // It takes postgres receiver, so not needed to get driver from config file
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
