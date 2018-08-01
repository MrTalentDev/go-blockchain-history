package assets

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"../../ethercrypto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"encoding/hex"
	"encoding/json"
	"time"
	"log"
)

// New asset
func New(c *gin.Context) {
	old := getAssetId(c)
	db := c.MustGet("db").(*mgo.Database)
	asset := models.Asset{}
	assetId := c.Param("assetId")
	if old == assetId {
		c.JSON(200, gin.H{
			"Answer": "This assetId is already created",
		})
		return
	}
	asset.AssetId = assetId
	asset.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)
	asset.UpdatedOn = asset.CreatedOn
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
		panic(err)
	}
	newID := result["txNumber"]
	c.JSON(http.StatusOK, gin.H{
		"txNumber": newID,
	})
}

func GetProof(c *gin.Context) {
	var m []string
	m = append(m, c.Param("assetID"))
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
	m = append(m, c.Param("assetID"))
	m = append(m, c.Param("txNumber"))

	var d = getTimestampAndDataHash(m)

	c.JSON(200, gin.H{
		"timestamp": d[0],
		"dataHash":  d[1],
	})
}

func Post(c *gin.Context) {
	var m []string
	m = append(m, c.Param("assetID"))
	m = append(m, c.Param("dataHash"))

	c.JSON(200, gin.H{
		"txNumber": getTxNumbers(),
	})
	if len(m) == 0 {
		c.JSON(200, gin.H{
			"txNumber": " ",
		})
	}

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

func getTxNumbers() string {

	return "1111"
}

func getTimestampAndDataHash(m []string) []string {

	var answer []string
	answer = append(answer, m[0])
	answer = append(answer, m[1])
	return answer
}

type Proof struct {
	HashProof string
}

type Proofs []Proof
