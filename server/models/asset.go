package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionAssets holds the name of the assets collection
	CollectionAssets = "assets"
)

// Assets model
type Asset struct {
	Id             bson.ObjectId     `json:"_id,omitempty" bson:"_id,omitempty"`
	Data           string            `json:"data" bson:"data"`
	Hash           []byte            `json:"hash" bson:"hash"`
	CreatedOn      int64             `json:"created_on" bson:"created_on"`
	UpdatedOn      int64             `json:"updated_on" bson:"updated_on"`
	AssetId        string            `json:"assetId" bson:"assetId"`
	TxNumber       int64             `json:"txNumber" bson:"txNumber"`
	Assets         map[string][]byte `json:"assets" bson:"assets"`
	AssetTimeStamp map[string]int64 `json:"assetsTimestamp" bson:"assetsTimestamp"`
}
