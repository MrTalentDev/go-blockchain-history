package responses

import "gopkg.in/mgo.v2/bson"

type CreateResponse struct {
	assetId    string `json:"assetId" example:"a"`
	hash       string `json:"hash" example:"96e75810b7fe519dd92f6a3f72170b00c0a8a9553f9c765a3cc681eaf7eeab38"`
	merkleRoot []byte `json:"merkleRoot" example:"Vu14mZ91jlhkqHhjFwmgjXgxyhLjLADVQlqMSQA3Q3o="`
	timestamp  uint64 `json:"timestamp" example:"1536920750859"`
	txNumber   uint64 `json:"txNumber" example:"0"`
}

type CreateResponseError struct {
	Answer string `json:"Answer" example:"This assetId is already created"`
}

type UpdateResponse struct {
	assetId   string `json:"assetId" example:"a"`
	timestamp uint64 `json:"timestamp" example:"1536920750859"`
	txNumber  uint64 `json:"txNumber" example:"1"`
}
type AssetsResponse struct {
	assets []byte `json:"assets" example:"1c8d54df80c03a56b5470d164c49f823108f96a67d020e4c677810c9a10b1ca7"`
}

type ListResponse struct {
	Id             bson.ObjectId     `json:"_id,omitempty" example:"5b869ee5ca2985e06552a49d"`
	Data           string            `json:"data" example:""`
	Hash           []byte            `json:"hash" example:"qNCllA0uMdgEPSVQBYzD4JESEECY2NyjbJgGjy0NP6c="`
	CreatedOn      int64             `json:"created_on" example:"1535549157514"`
	UpdatedOn      int64             `json:"updated_on" example:"1535549157514"`
	AssetId        string            `json:"assetId" example:"a"`
	TxNumber       int64             `json:"txNumber" example:"0"`
	Assets         map[string][]byte `json:"assets" example:""`
	AssetTimeStamp map[string]int64  `json:"assetsTimestamp" example:""`
}
