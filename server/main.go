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

	//
	//
	r.GET("/list", assets.List)

	//
	//
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

	//
	//
	r.POST("/new/:assetId", assets.New)

	//
	//
	r.POST("/atx/:assetId", assets.IncrementAssetTx)

	r.Run() // listen and serve on 0.0.0.0:8080
}
