package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/practice2017/photo-server/model"
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
	r.HandleFunc("/photos", getAllPhotoHandler).Methods("GET")
	r.HandleFunc("/photos", newPhotoHandler).Methods("POST")
	r.HandleFunc("/uploadfile", uploadHandler).Methods("POST")
	r.HandleFunc("/photos/{guid}", deletePhotoHandler).Methods("DELETE")

	log.Println("Runnung the server on port: 8000")
	http.ListenAndServe(":8000", r)
}
