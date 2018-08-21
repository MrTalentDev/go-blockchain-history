package tree

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../../ethercrypto/tree/customsmt"
	"../../models"
	"encoding/hex"
)

//RebuildOrCreateTree check's if there is any tree exists
//if it exists - rebuild it
//if not - create it
func RebuildOrCreateTree(c *gin.Context) {
	res := checkThatTreeIs(c)
	switch res {
	case true:
		d := getContent(c)
		rebuildTree(c, d)
	case false:
		createTreeContent(c)
	}

}

//GetTotalProof returns a total proof for all assets
func GetTotalProof(c *gin.Context) ([]string, string) {
	cont := getContent(c)
	t := customsmt.CreateTree(customsmt.CreateContent(cont))
	root := customsmt.GetMerkleRoot(t)
	s := hex.EncodeToString(root)
	return customsmt.Strings(t), s
}

//GetMerkleRoot returns main merkle root for the tree
func GetMerkleRoot(c *gin.Context) []byte {
	cont := getContent(c)
	t := customsmt.CreateTree(customsmt.CreateContent(cont))
	return customsmt.GetMerkleRoot(t)
}

const TREE_ID = "1"

func checkThatTreeIs(c *gin.Context) bool {
	db := c.MustGet("test").(*mgo.Database)
	query := bson.M{"TreeId": TREE_ID}
	tree := models.Tree{}
	err := db.C(models.CollectionTree).Find(query).One(&tree)
	if err != nil {
		println("checkThatTreeIs mistake 1", err)
	}
	//println(tree.Having)
	return tree.Having
}

func rebuildTree(c *gin.Context, content [][]byte) {
	db := c.MustGet("test").(*mgo.Database)
	query := bson.M{"TreeId": TREE_ID}
	data, _ := hex.DecodeString(c.Param("hash"))
	content = append(content, data)
	var result bson.M
	changeInDocument := mgo.Change{
		Update: bson.M{"$set": bson.M{"TreeContent": content}},
	}
	_, err := db.C(models.CollectionTree).Find(query).Apply(changeInDocument, &result)
	if err != nil {
		println("rebuildTree mistake 1")
	}
}

func createTreeContent(c *gin.Context) {
	db := c.MustGet("test").(*mgo.Database)
	var m [][]byte
	data, _ := hex.DecodeString(c.Param("hash"))
	m = append(m, data)
	tree := models.Tree{}
	tree.Having = true
	tree.TreeId = TREE_ID
	tree.TreeContent = m
	err := db.C(models.CollectionTree).Insert(tree)
	if err != nil {
		println("createTreeContent mistake 1")
		println(err)
	}
}

func getContent(c *gin.Context) [][]byte {
	db := c.MustGet("test").(*mgo.Database)
	query := bson.M{"TreeId": TREE_ID}
	tree := models.Tree{}
	err := db.C(models.CollectionTree).Find(query).One(&tree)
	if err != nil {
		println("getContent mistake 1")
		println(err)
	}
	return tree.TreeContent
}

//TODO make it work with DB without memory at all
