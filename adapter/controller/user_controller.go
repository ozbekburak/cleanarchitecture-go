package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ozbekburak/cleanarchitecture-go/adapter/gateway"
	"github.com/ozbekburak/cleanarchitecture-go/domain"
	"github.com/ozbekburak/cleanarchitecture-go/pkg/logger"
	"github.com/ozbekburak/cleanarchitecture-go/usecase/interactor"
)

type UserController struct {
	Interactor interactor.UserInteractor
}

func NewUserController(conn *sql.DB) *UserController {
	return &UserController{
		Interactor: interactor.UserInteractor{
			UserRepository: &gateway.UserRepository{
				Conn: conn,
			},
		},
	}
}

func (uc *UserController) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		var u domain.User

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			logger.Errorf("%s", err)

			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		userID, err := uc.Interactor.Create(u)
		if err != nil {
			logger.Errorf("%s", err)

			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusCreated, userID)
	}
}
