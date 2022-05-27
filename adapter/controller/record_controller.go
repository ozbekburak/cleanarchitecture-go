package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ozbekburak/cleanarch-mongo-inmem/adapter/gateway"
	"github.com/ozbekburak/cleanarch-mongo-inmem/domain"
	"github.com/ozbekburak/cleanarch-mongo-inmem/pkg/logger"
	"github.com/ozbekburak/cleanarch-mongo-inmem/usecase/interactor"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordController struct {
	Interactor interactor.RecordInteractor
}

// NewRecordController returns a new RecordController instance
func NewRecordController(conn *mongo.Client) *RecordController {
	return &RecordController{
		Interactor: interactor.RecordInteractor{
			RecordRepository: &gateway.RecordRepository{
				Conn: conn,
			},
		},
	}
}

// Filter triggers when a request is made to /records with startdate, enddate, mincount and maxcount query params
func (rc *RecordController) Filter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We only accept POST request, so any request other than POST will return 405
		switch r.Method {
		case http.MethodPost:
			logger.Infof("Endpoint hit %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
			var filter domain.Filter

			err := json.NewDecoder(r.Body).Decode(&filter)
			if err != nil {
				logger.Errorf("%s", err)

				RespondWithError(w, http.StatusBadRequest, err.Error())
				return
			}

			records, err := rc.Interactor.Filter(filter)
			if err != nil {
				logger.Errorf("%s", err)

				RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}

			RespondWithJSON(w, http.StatusOK, SuccessResponse{Code: 0, Msg: "success", Records: records.Records})

		default:
			logger.Errorf("Method %s not allowed %s %s\n", r.Method, r.RemoteAddr, r.URL)
			RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}
}
