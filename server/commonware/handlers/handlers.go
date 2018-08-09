package handlers

import (
	"github.com/gin-gonic/gin"
	"../assets"
	"../tree"
	"net/http"
	"encoding/hex"
)

// Add new asset to assetId and change merkle tree
func UpdateAssetId(c *gin.Context) {
	assets.UpdateAssetsByAssetId(c)
	defer assets.IncrementAssetTx(c)
}

// Create new assetId with asset
func CreateAssetId(c *gin.Context) {
	id, er := assets.CheckAndReturn(c)
	tree.RebuildOrCreateTree(c)
	if er == "err" {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"assetId": id,
	})

}

// Lists all assets in DB
func List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"assets": assets.FindALlAssets(c),
	})
}

// Get Merkle proof of specified asset
func GetSpecifiedProof(c *gin.Context) {

}

// Get total Merkle proof
func GetTotalProof(c *gin.Context) {

}

// Get timestamp and hash of specified asset in assetId
func GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"assets": hex.EncodeToString(assets.GetAssetByAssetIdAndTxNumber(c)),
	})
}
