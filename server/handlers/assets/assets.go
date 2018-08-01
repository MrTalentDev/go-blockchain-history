package assets

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"../../tokeccak"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"encoding/hex"
	)

// New asset
func New(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	asset := models.Asset{}
	assetId := c.Param("assetId")
	asset.AssetId = assetId
	asset.Hash = hex.EncodeToString(hashing.StringToKeccak(assetId))
	err := c.Bind(&asset)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionAssets).Insert(asset)

	if err != nil {
		c.Error(err)
	}
	println("success")
	//c.Redirect(http.StatusMovedPermanently, "/articles")
}

// List all assets
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	assets := []models.Asset{}
	err := db.C(models.CollectionAssets).Find(nil).All(&assets)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"assets": assets,
	})
}

// Returns txNumber for assetId
func ReturnAssetTx(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"txNumber": getTxNumber(c),
	})
}

// Returns incremented txNumber for assetId and safe it to db
func IncrementAssetTx(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	var result bson.M
	changeInDocument := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"txNumber": 1}},
		ReturnNew: true,
	}
	_, err:= db.C(models.CollectionAssets).Find(query).Apply(changeInDocument, &result)
	if err != nil {
		panic(err)
	}
	newID := result["txNumber"]
	c.JSON(http.StatusOK, gin.H{
		"txNumber": newID,
	})
}

func getTxNumber(c *gin.Context) int64 {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	asset := models.Asset{}
	err := c.Bind(&asset)
	if err != nil {
		c.Error(err)
	}
	err = db.C(models.CollectionAssets).Find(query).One(&asset)
	if err != nil {
		c.Error(err)
	}
	return asset.TxNumber
}

