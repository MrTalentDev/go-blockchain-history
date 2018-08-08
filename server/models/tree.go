package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionAssets holds the name of the assets collection
	CollectionTree = "tree"
)

// Assets model
type Tree struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Tree	  []byte		`json:"tree" bson:"tree"`
	Having    bool			`json:"having" bson:"having"`
	treeId	  string		`json:"treeId" bson:"treeId"`
}

