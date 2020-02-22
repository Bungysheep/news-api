package routes

import (
	defaultcontrollerv1 "github.com/bungysheep/news-api/pkg/controllers/v1/defaultcontroller"
	newscontrollerv1 "github.com/bungysheep/news-api/pkg/controllers/v1/newscontroller"
	"github.com/bungysheep/news-api/pkg/protocols/rest/middlewares"
	"github.com/gorilla/mux"
)

// APIV1RouteHandler builds Api v1 routes
func APIV1RouteHandler() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.DefaultMiddleware)

	v1Router := router.PathPrefix("/v1").Subrouter()

	dftCtl := defaultcontrollerv1.NewDefaultController()
	v1Router.HandleFunc("/", dftCtl.GetHome).Methods("GET")

	newsCtl := newscontrollerv1.NewNewsController()
	v1Router.HandleFunc("/news", newsCtl.PostNews).Methods("POST")
	v1Router.HandleFunc("/news", newsCtl.GetNews).Methods("GET")

	return router
}
