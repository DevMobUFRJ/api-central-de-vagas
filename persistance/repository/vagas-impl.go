package repository

import (
	"api-central-de-vagas/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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

func (r *Resource) getUsersCollection() *mgo.Collection {
	collection := r.mongoSession.DB("central-de-vagas").C("users")
	return collection
}