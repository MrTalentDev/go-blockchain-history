// using https://github.com/madhums/go-gin-mgo-demo/blob/master/db/connect.go code snippet

package db

import (
	"fmt"
	"os"
	"gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo

	LOGIN_DB    = os.Getenv("LOGIN_DB")
	PASSWORD_DB = os.Getenv("PASSWORD_DB")
	IP          = os.Getenv("IP")
)

//const (
//// MongoDBUrl is the default mongodb url that will be used to connect to the
//// database.
////
////MongoDBUrl = "mongodb://0.0.0.0:27017/demo"
////MongoDBUrl = "mongodb://bankex2:ObshiDostup1@18.209.40.150:27017/admin"
//)

// Connect connects to mongodb
func Connect() {
	//println(LOGIN_DB, PASSWORD_DB, IP)
	uri := ("mongodb://" + LOGIN_DB + ":" + PASSWORD_DB + "@" + IP + ":27017/"+"admin")
	//println("Jpiwfhbiohwcoiuwbcouiwb!!!!!!!!!")
	//println(uri)
	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}
