package assets

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"../../ethercrypto/hashing"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"encoding/json"
	"time"
	"log"
	"../../ethercrypto/tree"
	"encoding/hex"
)

// New asset
func New(c *gin.Context) {
	old := getAssetId(c)
	db := c.MustGet("db").(*mgo.Database)
	asset := models.Asset{}
	assetId := c.Param("assetId")
	assets := c.Param("assets")
	if old == assetId {
		c.JSON(200, gin.H{
			"Answer": "This assetId is already created",
		})
		return
	}
	asset.AssetId = assetId
	asset.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)
	asset.UpdatedOn = asset.CreatedOn
	asset.Hash = hashing.StringToKeccak(assetId)
	asset.Assets = assetId + hex.EncodeToString(hashing.StringToKeccak(assets))
	err := c.Bind(&asset)
	if err != nil {
		c.Error(err)
		return
	}
	err = db.C(models.CollectionAssets).Insert(asset)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"assetId": assetId,
	})
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
func incrementAssetTx(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	var result bson.M
	changeInDocument := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"txNumber": 1, "updated_on": time.Now().UnixNano() / int64(time.Millisecond),}},
		ReturnNew: true,
	}
	_, err := db.C(models.CollectionAssets).Find(query).Apply(changeInDocument, &result)
	if err != nil {
		panic(err)
	}
	newID := result["txNumber"]
	c.JSON(http.StatusOK, gin.H{
		"txNumber": newID,
	})
}

func GetProof(c *gin.Context) {
	var m []string
	m = append(m, c.Param("assetId"))
	m = append(m, c.Param("txNumber"))

	var d = getMerkleProof()

	var proofs = Proofs{}

	for i := 0; i < len(d); i++ {
		var a = d[i]
		proofs = append(proofs, Proof{
			a,
		})
	}

	myJson, err := json.Marshal(proofs)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	c.Data(200, "JSON", myJson)
}

func Get(c *gin.Context) {
	var m []string
	m = append(m, c.Param("assetId"))
	m = append(m, c.Param("txNumber"))

	var d = getTimestampAndDataHash(m)

	c.JSON(200, gin.H{
		"timestamp": d[0],
		"dataHash":  d[1],
	})
}

func Post(c *gin.Context) {
	oldAssets := getsAssetsByAssetById(c)
	updateAssetsByAssetId(c, oldAssets)
	defer incrementAssetTx(c)
}

func Test(c *gin.Context) {
	var m []string
	m = append(m, c.Param("asset"))
	c.JSON(200, gin.H{
		"txNumber": tree.Tree(m),
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

func getAssetId(c *gin.Context) string {
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
	return asset.AssetId
}

func getMerkleProof() []string {

	var m []string
	m = append(m, "1")
	m = append(m, "1")
	m = append(m, "1")
	m = append(m, "1")
	m = append(m, "1")
	return m
}

func updateAssetsByAssetId(c *gin.Context, oldAssets string) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"assetId": c.Param("assetId")}
	newAssets := oldAssets + c.Param("assetId") + c.Param("assets")
	println(newAssets)
	var result bson.M
	changeInDocument := mgo.Change{
		Update: bson.M{"$set": bson.M{"updated_on": time.Now().UnixNano() / int64(time.Millisecond), "assets": newAssets}},
	}
	_, err := db.C(models.CollectionAssets).Find(query).Apply(changeInDocument, &result)
	if err != nil {
		panic(err)
	}
}

func getTimestampAndDataHash(m []string) []string {

	var answer []string
	answer = append(answer, m[0])
	answer = append(answer, m[1])
	return answer
}

func getsAssetsByAssetById(c *gin.Context) string {
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
	return asset.Assets
}

type Proof struct {
	HashProof string
}

type Proofs []Proof
