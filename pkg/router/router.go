package router

import (
	"database/sql"
	"net/http"

	"github.com/ozbekburak/cleanarch-mongo-inmem/adapter/controller"
)

// Initialize function initializes the router
func Initialize(conn *sql.DB) *http.ServeMux {
	userController := controller.NewUserController(conn)

	router := http.NewServeMux()

	router.Handle("/user", userController.Store())

	return router
}
