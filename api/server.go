package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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
	r.HandleFunc("/photos", firstOptionsHandler).Methods("OPTIONS")
	r.HandleFunc("/photos/{guid}", deletePhotoHandler).Methods("DELETE")
	r.HandleFunc("/photos/{guid}", secondOptionsHandler).Methods("OPTIONS")

	log.Println("Runnung the server on port: 8000")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r))

}
