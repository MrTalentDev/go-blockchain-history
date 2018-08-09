package tree

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../../models"
)

func RebuildOrCreateTree(c *gin.Context) {
	res := checkThatTreeIs(c)
	switch res {
	case true:
		rebuildTree(c)
	case false:
		createTree(c)
	}

}

func GetSpecificProof(c *gin.Context) {

}

func GetTotalProof(c *gin.Context) {

}

func GetMerkleRoot(c *gin.Context) {

}

////////////////////////////////////
//////// Internal functions ////////
////////////////////////////////////

const TREE_ID = "1"

func checkThatTreeIs(c *gin.Context) bool {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"TreeId": TREE_ID}
	tree := models.Tree{}
	err := db.C(models.CollectionAssets).Find(query).One(&tree)
	if err != nil {
		println("checkThatTreeIs mistake 1", err)
	}
	println(tree.Having)
	return tree.Having
}

func rebuildTree(c *gin.Context) {
	println("hhhhhhh")
}

func createTree(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	tree := models.Tree{}
	tree.Having = true
	tree.TreeId = TREE_ID
	err := db.C(models.CollectionTree).Insert(tree)
	if err != nil {
		println("InitAsset mistake 2")
	}
	println("first")
}
