package gateway

import (
	"database/sql"

	"github.com/ozbekburak/cleanarchitecture-go/domain"
)

type UserRepository struct {
	Conn *sql.DB
}

func (ur *UserRepository) Create(u domain.User) (int64, error) {
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`

	var id int64
	err := ur.Conn.QueryRow(query, u.ID).Scan(&id)

	return id, err
}
