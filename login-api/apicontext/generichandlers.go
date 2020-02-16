package apicontext

import (
	"fmt"
	"io/ioutil"
	"net/http"

	util "github.com/serdarkalayci/go-grpc/login-api/util"
)

// Index swagger:route GET / Index
//
// Handler to show welcome message.
//
// Responses:
//        200: OK
func (apiContext *APIContext) Index(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("./static/version.txt")
	if err != nil {
		dat = append(dat, '0')
	}
	fmt.Fprintf(w, "Welcome to Cookie Monster Web API! Version:%s", dat)
	util.Logger.TRACE.Log("Index called")
}

// CorsHandler swagger:route OPTIONS /
//
// Handler to respond to CORS preflight requests
//
// Responses:
//        200: OK
func CorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, lang")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, OPTIONS")
	w.WriteHeader(200)
}
