package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Run runs the server
func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods("GET")

	log.Println("Runnung the server on port: 8000")
	http.ListenAndServe(":8000", r)
}
