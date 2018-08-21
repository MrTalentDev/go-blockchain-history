package models

const (
	// CollectionAssets holds the name of the assets collection
	CollectionTree = "tree"
)

// Assets model
type Tree struct {
	TreeContent [][]byte `json:"TreeContent" bson:"TreeContent"`
	Having      bool     `json:"having" bson:"having"`
	TreeId      string   `json:"TreeId" bson:"TreeId"`
}
