package assets

import (
	"github.com/gin-gonic/gin"
	"../../db/models"
	"../ethercrypto/hashing"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
	"strconv"
	"encoding/hex"
)

//CheckAndReturn goes to mgo and check if there is an Asset Id
//if there is no assetId it calls InitAsset
//returns (asset.ID, asset.Hash) + error (string type), result of checking (bool)
func CheckAndReturn(c *gin.Context) ([]string, string, bool) {
	_, err := GetAssetId(c)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Answer": "This assetId is already created",
		})
		return nil, "err", false
	}
	return InitAsset(c), "", true
}

//Check
func Check(c *gin.Context) (bool) {
	_, err := GetAssetId(c)
	if err == nil {
		return true
	} else {
		c.JSON(200, gin.H{
			"Answer": "This assetId is not created",
		})
		return false
	}
}

//GetAssetId returns assetId it it exists
func GetAssetId(c *gin.Context) (string, error) {
	db := c.MustGet("test").(*mgo.Database)
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

//InitAsset creates new asset Id with pending asset
func InitAsset(c *gin.Context) []string {
	db := c.MustGet("test").(*mgo.Database)
	timing := time.Now().UnixNano() / int64(time.Millisecond)
	asset := models.Asset{}
	assetId := c.Param("assetId")
	assets := c.Param("hash")
	asset.AssetId = assetId
	asset.CreatedOn = timing
	asset.UpdatedOn = asset.CreatedOn
	asset.Hash = hashing.StringToKeccak(assetId)
	//asset.AssetTimeStamp = time.Now().UnixNano() / int64(time.Millisecond)
	m := make(map[string][]byte)
	timstamp := make(map[string]int64)
	data, _ := hex.DecodeString(assets)
	m["0"] = data
	timstamp["0"] = timing
	//println(m["0"])
	//println(hex.EncodeToString(m["0"]))
	asset.Assets = m
	asset.AssetTimeStamp = timstamp
	err := c.Bind(&asset)
	if err != nil {
		println("InitAsset mistake 1")
		return nil
	}
	err = db.C(models.CollectionAssets).Insert(asset)
	if err != nil {
		println("InitAsset mistake 2")
	}
	var a []string
	a = append(a, asset.AssetId)
	a = append(a, hex.EncodeToString(data))
	a = append(a, strconv.Itoa(int(timing)))
	return a
}

//FindALlAssets returns all assets in mgo
func FindALlAssets(c *gin.Context) []models.Asset {
	db := c.MustGet("test").(*mgo.Database)
	assets := []models.Asset{}
	err := db.C(models.CollectionAssets).Find(nil).All(&assets)
	if err != nil {
		println("FindALlAssets mistake 1")
	}
	return assets
}

//GetAssetsByAssetById returns all assets by id
func GetAssetsByAssetById(c *gin.Context) map[string][]byte {
	db := c.MustGet("test").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	asset := models.Asset{}
	err := db.C(models.CollectionAssets).Find(query).One(&asset)
	if err != nil {
		println("GetAssetsByAssetById mistake 2")
	}
	return asset.Assets
}

//UpdateAssetsByAssetId allow to add new asset to assetId
func UpdateAssetsByAssetId(c *gin.Context) int64 {
	db := c.MustGet("test").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	txNumber := GetTxNumber(c)
	txNumber++
	stringTx := strconv.FormatInt(txNumber, 10)
	m := GetAssetsByAssetById(c)
	data, _ := hex.DecodeString(c.Param("hash"))
	m[stringTx] = data
	timstamp := GetTimestamp(c)
	timing := time.Now().UnixNano() / int64(time.Millisecond)
	timstamp[stringTx] = timing
	var result bson.M
	changeInDocument := mgo.Change{
		Update: bson.M{"$set": bson.M{"updated_on": timing, "assets": m, "assetsTimestamp": timstamp}},
	}
	_, err := db.C(models.CollectionAssets).Find(query).Apply(changeInDocument, &result)
	if err != nil {
		println("UpdateAssetsByAssetId mistake 1")
	}
	return timing
}

// Returns incremented txNumber for assetId and save it to db
func IncrementAssetTx(c *gin.Context, timestamp int64) {
	db := c.MustGet("test").(*mgo.Database)
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
		"timestamp": timestamp,
		"assetId": result["assetId"],
	})
}

//GetAssetByAssetIdAndTxNumber returns asset by it's assetId and txNumber
func GetAssetByAssetIdAndTxNumber(c *gin.Context) []byte {
	db := c.MustGet("test").(*mgo.Database)
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

//GetTxNumber returns last txNumber of assetId
func GetTxNumber(c *gin.Context) int64 {
	db := c.MustGet("test").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	asset := models.Asset{}
	err := db.C(models.CollectionAssets).Find(query).One(&asset)
	if err != nil {
		println("tx mistake 2")
	}
	return asset.TxNumber
}
//GetTxNumber returns last txNumber of assetId
func GetTimestamp(c *gin.Context) map[string]int64 {
	db := c.MustGet("test").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	asset := models.Asset{}
	err := db.C(models.CollectionAssets).Find(query).One(&asset)
	if err != nil {
		println("tx mistake 2")
	}
	return asset.AssetTimeStamp
}
