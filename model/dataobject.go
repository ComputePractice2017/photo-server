package model

import (
	"log"
	"os"

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
	dbaddress := os.Getenv("RETHINKDB_HOST")
	if dbaddress == "" {
		dbaddress = "192.168.99.100"
	}

	log.Printf("RETHINKDB_HOST: %s\n", dbaddress)
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: dbaddress,
	})
	if err != nil {
		return err
	}

	err = CreateDBIfNotExist()
	if err != nil {
		return err
	}

	err = CreateTableIfNotExist()

	return err
}

func CreateDBIfNotExist() error {
	res, err := r.DBList().Run(session)
	if err != nil {
		return err
	}

	var dbList []string
	err = res.All(&dbList)
	if err != nil {
		return err
	}

	for _, item := range dbList {
		if item == "instagram" {
			return nil
		}
	}

	_, err = r.DBCreate("instagram").Run(session)
	if err != nil {
		return err
	}

	return nil
}

func CreateTableIfNotExist() error {
	res, err := r.DB("instagram").TableList().Run(session)
	if err != nil {
		return err
	}

	var tableList []string
	err = res.All(&tableList)
	if err != nil {
		return err
	}

	for _, item := range tableList {
		if item == "photo" {
			return nil
		}
	}

	_, err = r.DB("instagram").TableCreate("photo", r.TableCreateOpts{PrimaryKey: "ID"}).Run(session)

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
