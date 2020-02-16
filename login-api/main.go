// GoGrpc Login API service.
//
// This service can be used to interact with cookie consent management backend
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	apicontext "github.com/serdarkalayci/go-grpc/login-api/apicontext"
)

const ()

func main() {

	apiContext := apicontext.PrepareContext()
	router := NewRouter(apiContext)
	fs := http.FileServer(http.Dir("./swagger/"))
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
	router.PathPrefix("/metrics").Handler(promhttp.Handler())
	prometheus.MustRegister(apicontext.RequestCounterVec)
	prometheus.MustRegister(apicontext.RequestDurationGauge)

	fmt.Println("Http Server started")
	log.Fatal(http.ListenAndServe(":6543", router))

}
