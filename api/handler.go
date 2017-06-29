package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/photo-server/model"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!")
}

func getAllPersonHandler(w http.ResponseWriter, r *http.Request) {
	persons, err := model.GetPersons()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err = json.NewEncoder(w).Encode(persons); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(http.StatusOK)
}
