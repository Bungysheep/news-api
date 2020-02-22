package routes

import (
	"github.com/bungysheep/news-api/pkg/controllers/v1/defaultcontroller"
	"github.com/gorilla/mux"
)

// APIV1RouteHandler builds Api v1 routes
func APIV1RouteHandler() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	v1Router := router.PathPrefix("/v1").Subrouter()

	dftCtl := defaultcontroller.NewDefaultController()
	v1Router.HandleFunc("/", dftCtl.GetHome).Methods("GET")

	return router
}
