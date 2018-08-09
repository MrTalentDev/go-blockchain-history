package models

const (
	// CollectionAssets holds the name of the assets collection
	CollectionTree = "tree"
)

// Assets model
type Tree struct {
	Tree	  []byte		`json:"tree" bson:"tree"`
	Having    bool			`json:"having" bson:"having"`
	TreeId	  string		`json:"TreeId" bson:"TreeId"`
}

