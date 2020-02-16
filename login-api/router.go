package main

import (
	"github.com/gorilla/mux"
	apicontext "github.com/serdarkalayci/go-grpc/login-api/apicontext"
)

// NewRouter creates a new gorilla/mux router, adds handlers to it and returns
func NewRouter(apiContext *apicontext.APIContext) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router.Methods("OPTIONS").HandlerFunc(apicontext.CorsHandler)
	router.HandleFunc("/", apiContext.Index).Methods("GET")
	router.HandleFunc("/login", apiContext.Login).Methods("POST")
	router.HandleFunc("/login/", apiContext.Login).Methods("POST")
	router.HandleFunc("/login", apiContext.Refresh).Methods("GET")
	router.HandleFunc("/login/", apiContext.Refresh).Methods("GET")
	router.Use(apicontext.MonitoringMiddleware)
	return router
}
