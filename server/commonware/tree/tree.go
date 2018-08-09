package tree

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../../ethercrypto/tree/customsmt"
	"../../models"
)

func RebuildOrCreateTree(c *gin.Context) {
	res := checkThatTreeIs(c)
	switch res {
	case true:
		return
	case false:
		createTreeContent(c)
	}

}

func GetSpecificProof(c *gin.Context) bool {
	cont := getContent(c)
	t := customsmt.CreateTree(customsmt.CreateContent(cont))
	var m []string
	m = append(m, c.Param("assets"))
	conts := customsmt.CreateContent(m)
	res, _ := t.VerifyContent(conts[0])
	return res
}

func GetTotalProof(c *gin.Context) []string {
	cont := getContent(c)
	t := customsmt.CreateTree(customsmt.CreateContent(cont))
	return customsmt.Strings(t)
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
	err := db.C(models.CollectionTree).Find(query).One(&tree)
	if err != nil {
		println("checkThatTreeIs mistake 1", err)
	}
	println(tree.Having)
	return tree.Having
}

//func rebuildTree(c *gin.Context) {
//	println("rebuilding tree")
//}

func createTreeContent(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var m []string
	m = append(m, c.Param("assets"))
	tree := models.Tree{}
	tree.Having = true
	tree.TreeId = TREE_ID
	tree.TreeContent = m
	err := db.C(models.CollectionTree).Insert(tree)
	if err != nil {
		println(err)
	}
}

func getContent(c *gin.Context) []string {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"TreeId": TREE_ID}
	tree := models.Tree{}
	db.C(models.CollectionTree).Find(query).One(&tree)
	return tree.TreeContent
}
