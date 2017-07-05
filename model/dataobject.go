package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

//Person - Пользователь, который добавил фото
type Photo struct {
	ID   string `json:"id",gorethink:"id"`
	Name string `json:"name",gorethink:"name"`
	Url  string `json:"url",gorethink:"url"`
}

var session *r.Session

func InitSesson() error {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	return err
}

func GetPhotos() ([]Photo, error) {
	res, err := r.DB("instagram").Table("photo").Run(session)
	if err != nil {
		return nil, err
	}

	var response []Photo
	err = res.All(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewPhoto(p Photo) (Photo, error) {
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

	res, err = r.DB("instagram").Table("photo").Insert(p).Run(session)
	if err != nil {
		return p, err
	}

	return p, nil
}

func DeletePhoto(id string) error {
	_, err := r.DB("instagram").Table("photo").Get(id).Delete().Run(session)
	if err != nil {
		return err
	}
	return nil
}
