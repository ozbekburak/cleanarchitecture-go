package router

import (
	"net/http"

	"github.com/ozbekburak/go-clean-api/adapter/controller"
	"github.com/ozbekburak/go-clean-api/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

// Initialize function initializes the router
func Initialize(conn *mongo.Client, inMem *db.InMemoryDB) *http.ServeMux {
	recordController := controller.NewRecordController(conn)
	inMemoryController := controller.NewInMemoryController(inMem)

	router := http.NewServeMux()

	router.Handle("/records", recordController.Filter())
	// In-memory DB handlers go here
	router.Handle("/in-memory/set", inMemoryController.Store())
	router.Handle("/in-memory", inMemoryController.Get())

	return router
}
