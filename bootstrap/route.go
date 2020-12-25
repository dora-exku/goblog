package bootstrap

import (
	"goblog/routes"

	"github.com/gorilla/mux"
)

func SetopRoute() *mux.Router {
	router := mux.Router
	routes.RegisterWebRoutes(router)
	return router
}
