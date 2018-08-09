package assets

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"../../ethercrypto/hashing"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
	"strconv"
)

func CheckAndReturn(c *gin.Context) string {
	_, err := GetAssetId(c)
	if err == nil {
		c.JSON(200, gin.H{
			"Answer": "This assetId is already created",
		})
		return ""
	}
	return InitAsset(c)
}

func GetAssetId(c *gin.Context) (string, error) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	asset := models.Asset{}
	err := c.Bind(&asset)
	if err != nil {
		println(err)
	}
	err = db.C(models.CollectionAssets).Find(query).One(&asset)
	if err != nil {
		return "", err
	}
	return asset.AssetId, nil
}

func InitAsset(c *gin.Context) string {
	db := c.MustGet("db").(*mgo.Database)
	asset := models.Asset{}
	assetId := c.Param("assetId")
	assets := c.Param("assets")
	asset.AssetId = assetId
	asset.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)
	asset.UpdatedOn = asset.CreatedOn
	asset.Hash = hashing.StringToKeccak(assetId)
	m := make(map[string][]byte)
	m["0"] = hashing.StringToKeccak(assets)
	asset.Assets = m
	err := c.Bind(&asset)
	if err != nil {
		println("InitAsset mistake 1")
		return ""
	}
	err = db.C(models.CollectionAssets).Insert(asset)
	if err != nil {
		println("InitAsset mistake 2")
	}
	return asset.AssetId
}

func FindALlAssets(c *gin.Context) []models.Asset {
	db := c.MustGet("db").(*mgo.Database)
	assets := []models.Asset{}
	err := db.C(models.CollectionAssets).Find(nil).All(&assets)
	if err != nil {
		println("FindALlAssets mistake 1")
	}
	return assets
}

func GetAssetsByAssetById(c *gin.Context) map[string][]byte {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	asset := models.Asset{}
	//err := c.Bind(&asset)
	//if err != nil {
	//	println("GetAssetsByAssetById mistake 1")
	//}
	err := db.C(models.CollectionAssets).Find(query).One(&asset)
	if err != nil {
		println("GetAssetsByAssetById mistake 2")
	}
	return asset.Assets
}

func UpdateAssetsByAssetId(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	txNumber := getTxNumber(c)
	txNumber++
	stringTx := strconv.FormatInt(txNumber, 10)
	m := GetAssetsByAssetById(c)
	m[stringTx] = hashing.StringToKeccak(c.Param("assets"))
	//println(newAssets)
	var result bson.M
	changeInDocument := mgo.Change{
		Update: bson.M{"$set": bson.M{"updated_on": time.Now().UnixNano() / int64(time.Millisecond), "assets": m}},
	}
	_, err := db.C(models.CollectionAssets).Find(query).Apply(changeInDocument, &result)
	if err != nil {
		println("UpdateAssetsByAssetId mistake 1")
	}
}

// Returns incremented txNumber for assetId and safe it to db
func IncrementAssetTx(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	var result bson.M
	changeInDocument := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"txNumber": 1, "updated_on": time.Now().UnixNano() / int64(time.Millisecond),}},
		ReturnNew: true,
	}
	_, err := db.C(models.CollectionAssets).Find(query).Apply(changeInDocument, &result)
	if err != nil {
		println("IncrementAssetTx mistake 1")
	}
	newID := result["txNumber"]
	c.JSON(http.StatusOK, gin.H{
		"txNumber": newID,
	})
}

func GetAssetByAssetIdAndTxNumber(c *gin.Context) []byte {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	var st string
	st = c.Param("txNumber")
	txNumber, err := strconv.ParseInt(st, 10, 64)
	if err != nil {
		println(err)
	}
	asset := models.Asset{}
	err = c.Bind(&asset)
	if err != nil {
		println("GetAssetByAssetIdAndTxNumber mistake 1")
	}
	err = db.C(models.CollectionAssets).Find(query).One(&asset)
	if err != nil {
		println("GetAssetByAssetIdAndTxNumber mistake 1")
	}
	m := asset.Assets
	stringTx := strconv.FormatInt(txNumber, 10)
	return m[stringTx]
}

////////////////////////////////////
//////// Internal functions ////////
////////////////////////////////////

func getTxNumber(c *gin.Context) int64 {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	asset := models.Asset{}
	//err := c.Bind(&asset)
	//if err != nil {
	//	println("tx mistake 1")
	//}
	err := db.C(models.CollectionAssets).Find(query).One(&asset)
	if err != nil {
		println("tx mistake 2")
	}
	return asset.TxNumber
}
