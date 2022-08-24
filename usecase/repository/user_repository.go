package repository

import "github.com/ozbekburak/cleanarchitecture-go/domain"

// UserRepository interface defines methods for CRUD operations on User
type UserRepository interface {
	Create(domain.User) (int64, error)
}
