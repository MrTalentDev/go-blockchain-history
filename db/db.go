package db

import (
	"fmt"
	"github.com/BANKEX/poa-history/config"
	"gopkg.in/mgo.v2"
)

type InstanceDB struct {
	Info    *mgo.DialInfo
	Session *mgo.Session
}

var GlobalDB InstanceDB

func Connect(cfg *config.Config) error {
	uri := ("mongodb://" + cfg.DatabaseLogin + ":" + cfg.DatabasePassword + "@" + cfg.DatabaseIP + ":27017/" + "admin")
	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		return err
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)

	i := InstanceDB{
		Session: s,
		Info:    mongo,
	}
	GlobalDB = i
	return nil
}
