package database

import (
	"crypto/tls"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"net"
)

func MongoDBConnect(url string) *mgo.Session {

	if len(url) == 0 {
		panic("Can`t start because mongo uri is empty")
	}

	dialInfo, err := mgo.ParseURL(url)

	if err != nil {
		panic(err)
	}

	tlsConfig := tls.Config{}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), &tlsConfig)
		return conn, err
	}

	panic(url)
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err)
	} else {
		logrus.Info(fmt.Sprintf("Connected with MongoDB in database:[%s]", dialInfo.Database))
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}