package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/practice2017/photo-server/model"
	"github.com/renstrom/shortuuid"
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

func newPersonHandler(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := json.Unmarshal(body, &person); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	person, err = model.NewPerson(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")
	u := shortuuid.New()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	out, err := os.Create("./upload/" + u + ".jpg")
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)
}
