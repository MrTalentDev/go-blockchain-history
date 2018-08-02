package main

import (
	"github.com/gin-gonic/gin"
	"./middlewares"
	"./handlers/assets"
	"./db"
)

func main() {

	// mongo init
	db.Connect()

	r := gin.Default()
	r.Use(middlewares.Connect)
	r.Use(middlewares.ErrorHandler)

	// List all assets
	// http://localhost:8080/list
	r.GET("/list", assets.List)

	// return last txNumber of assetId
	// http://localhost:8080/tx/assetOne
	r.GET("/tx/:assetId", assets.ReturnAssetTx)

	// getProof(assetID, txNumber) - достает merkleproof
	// http://localhost:8080/getProof/g/g
	r.GET("/getProof/:assetID/:txNumber", assets.GetProof)

	// get(assetID, txNumber) - достает (timestamp, dataHash)
	// http://localhost:8080/get/1/11
	r.GET("/get/:assetID/:txNumber", assets.Get)

	// post(assetID, dataHash) - добавляет данные для данного assetId, автоинкрементит txNumber. Возвращает txNumber.
	// http://localhost:8080/post/1/11
	r.POST("/post/:assetID/:dataHash", assets.Post)

	// create new asset with assetId
	// http://localhost:8080/new/assetOne
	r.POST("/new/:assetId", assets.New)

	// Return incremented txNumber and saves it
	// http://localhost:8080/atx/assetTwo
	r.POST("/atx/:assetId", assets.IncrementAssetTx)

	r.Run() // listen and serve on 0.0.0.0:8080
}
