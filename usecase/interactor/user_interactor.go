package interactor

import (
	"github.com/ozbekburak/cleanarchitecture-go/domain"
	"github.com/ozbekburak/cleanarchitecture-go/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func (ui *UserInteractor) Create(u domain.User) (int64, error) {
	var id int64

	id, err := ui.UserRepository.Create(u)
	return id, err
}
