package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

//Person - Пользователь, который добавил фото
type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var session *r.Session

func InitSesson() error {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	return err
}

func GetPersons() ([]Person, error) {
	res, err := r.DB("instagram").Table("person").Run(session)
	if err != nil {
		return nil, err
	}

	var response []Person
	err = res.All(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
