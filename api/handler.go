package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/practice2017/photo-server/model"
	"github.com/renstrom/shortuuid"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!")
}

func getAllPhotoHandler(w http.ResponseWriter, r *http.Request) {
	photos, err := model.GetPhotos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err = json.NewEncoder(w).Encode(photos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(http.StatusOK)
}

func newPhotoHandler(w http.ResponseWriter, r *http.Request) {
	var photo model.Photo
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
	if err := json.Unmarshal(body, &photo); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	photo, err = model.NewPhoto(photo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(photo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
func deletePhotoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := model.DeletePhoto(vars["guid"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// the FormFile function takes in the POST input id file
	file, _, err := r.FormFile("file")
	u := shortuuid.New()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	out, err := os.Create("./upload/" + u + ".jpg")
	if err != nil {
		fmt.Fprintf(w, "Скорее всего не создана папка.")
		return
	}
	var urlphotostr string = "./upload/" + u + ".jpg"
	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	//fmt.Fprintf(w, "File uploaded successfully : ")
	//fmt.Fprintf(w, header.Filename)
	fmt.Fprintf(w, urlphotostr)
}
