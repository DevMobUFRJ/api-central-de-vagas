package repository

import (
	"api-central-de-vagas/model"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"io/ioutil"
	"os"
	"time"
)

type Resource struct {
	mongoSession *mgo.Session
}

func NewRepository(mongoSession *mgo.Session) Vagas {
	return &Resource{
		mongoSession: mongoSession,
	}
}

func (r *Resource) CreateUser(user *model.User) error {
	if err := r.getUsersCollection().Insert(user); err != nil {
		return err
	}

	return nil
}

func (r *Resource) UpdateUser(user *model.User) error {
	update := bson.M{"$set": &user}

	if err := r.getUsersCollection().UpdateId(user.Id, update); err != nil {
		return err
	}

	return nil
}

func (r *Resource) SendCurriculum(curriculum *os.File, createdAt time.Time) (interface{}, error) {
	file, err := r.mongoSession.DB("central-de-vagas").GridFS("fs").Create(fmt.Sprintf("%s - %s.pdf", curriculum.Name(), createdAt))
	if err != nil {
		return "", err
	}

	byteSlice, err := ioutil.ReadAll(curriculum)
	if err != nil {
		return "", err
	}

	if _, err := file.Write(byteSlice); err != nil {
		return "", err
	}

	if err := file.Close(); err != nil {
		return "", err
	}

	return file.Id(), nil
}

func (r *Resource) GetUserByUID(uid string) (*model.User, error) {
	query := bson.M{
		"uid": uid,
	}

	var user model.User

	if err := r.getUsersCollection().Find(query).One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

/*
func (r *Resource) GetCurriculum(photo *model.Photo) error {

}
*/

func (r *Resource) getUsersCollection() *mgo.Collection {
	collection := r.mongoSession.DB("central-de-vagas").C("users")
	return collection
}