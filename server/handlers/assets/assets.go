package assets

import (
		"github.com/gin-gonic/gin"
	    "../../models"
		"gopkg.in/mgo.v2"
	"net/http"
)

// New asset
func New(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	asset := models.Asset{}
	assetId := c.Param("assetId")
	asset.AssetId = assetId
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
		"articles": assets,
	})
}