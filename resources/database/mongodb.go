package database

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

func MongoDBConnect(url string) *mgo.Session {

	if len(url) == 0 {
		panic("Can`t start because mongo uri is empty")
	}

	session, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	} else {
		logrus.Info(fmt.Sprintf("Connected with MongoDB"))
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}