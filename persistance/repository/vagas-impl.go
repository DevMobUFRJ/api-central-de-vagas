package repository

import (
	"github.com/globalsign/mgo"
)

type Resource struct {
	mongoSession *mgo.Session
}

func NewRepository(mongoSession *mgo.Session) Vagas {
	return &Resource{
		mongoSession: mongoSession,
	}
}

func (r *Resource) getUsersCollection() *mgo.Collection {
	collection := r.mongoSession.DB("central-de-vagas").C("users")
	return collection
}