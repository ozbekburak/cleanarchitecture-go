package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ozbekburak/go-clean-api/adapter/gateway"
	"github.com/ozbekburak/go-clean-api/domain"
	"github.com/ozbekburak/go-clean-api/pkg/db"
	"github.com/ozbekburak/go-clean-api/pkg/logger"
	"github.com/ozbekburak/go-clean-api/usecase/interactor"
)

type InMemoryController struct {
	Interactor interactor.InMemoryInteractor
}

// NewInMemoryController returns a new InMemoryController instance
func NewInMemoryController(conn *db.InMemoryDB) *InMemoryController {
	return &InMemoryController{
		Interactor: interactor.InMemoryInteractor{
			InMemoryRepository: &gateway.InMemoryRepository{
				Conn: conn,
			},
		},
	}
}

// Store triggers when a request is made to /in-memory/set
func (imc *InMemoryController) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("Endpoint hit %s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		var kv domain.KeyValue

		err := json.NewDecoder(r.Body).Decode(&kv)
		if err != nil {
			logger.Errorf("%s", err)

			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		kv, err = imc.Interactor.Create(kv)
		if err != nil {
			logger.Errorf("%s", err)

			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, kv)
	}
}

// Get triggers when a request is made to /in-memory?{key}={value}}
func (imc *InMemoryController) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("Endpoint hit %s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		key := r.URL.Query().Get("key")
		if key == "" {
			RespondWithError(w, http.StatusBadRequest, "key is required")
			return
		}

		key, value, ok := imc.Interactor.Get(key)
		if !ok {
			RespondWithError(w, http.StatusNotFound, "key not found")
			return
		}

		RespondWithJSON(w, http.StatusOK, domain.KeyValue{
			Key:   key,
			Value: value,
		})
	}
}
