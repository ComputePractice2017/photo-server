package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

//Person - Пользователь, который добавил фото
type Person struct {
	ID    string `json:"id",gorethink:"id"`
	Name  string `json:"name",gorethink:"name"`
	Photo string `json:"photo",gorethink:"photo"`
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

func NewPerson(p Person) (Person, error) {
	res, err := r.UUID().Run(session)
	if err != nil {
		return p, err
	}

	var UUID string
	err = res.One(&UUID)
	if err != nil {
		return p, err
	}

	p.ID = UUID
	//p.Name = namephoto
	//p.Photo = urlphoto

	res, err = r.DB("instagram").Table("person").Insert(p).Run(session)
	if err != nil {
		return p, err
	}

	return p, nil
}
