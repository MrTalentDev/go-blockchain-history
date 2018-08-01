package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionAssets holds the name of the assets collection
	CollectionAssets = "assets"
)

// Assets model
type Asset struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Data	  string		`json:"data" bson:"data"`
	Hash	  string		`json:"hash" bson:"hash"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
	AssetId   string		`json:"assetId" bson:"assetId"`
	TxNumber int64         `json:"txNumber" bson:"txNumber"`
}
