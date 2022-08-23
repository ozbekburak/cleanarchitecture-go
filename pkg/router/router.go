package router

import (
	"net/http"

	"github.com/ozbekburak/cleanarch-mongo-inmem/adapter/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

// Initialize function initializes the router
func Initialize(conn *mongo.Client) *http.ServeMux {
	recordController := controller.NewRecordController(conn)

	router := http.NewServeMux()

	router.Handle("/records", recordController.Filter())

	return router
}
