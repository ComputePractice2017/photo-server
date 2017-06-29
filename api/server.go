package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/photo-server/model"
)

//Run runs the server
func Run() {
	log.Println("Connecting to rethinkDB on localhost...")
	err := model.InitSesson()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected!")

	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods("GET")
	r.HandleFunc("/persons", getAllPersonHandler).Methods("GET")

	log.Println("Runnung the server on port: 8000")
	http.ListenAndServe(":8000", r)
}
