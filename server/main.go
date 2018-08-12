package main

import (
	"./db"
	"./middlewares"
	"./commonware/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	// mongo init
	db.Connect()

	r := gin.Default()
	r.Use(middlewares.Connect)
	r.Use(middlewares.ErrorHandler)
	a := r.Group("/a", gin.BasicAuth(gin.Accounts{
		"bankex": "ObshiDostup1",
	}))

	a.POST("/new/:assetId/:assets", handlers.CreateAssetId)
	a.POST("/update/:assetId/:assets", handlers.UpdateAssetId)
	r.GET("/get/:assetId/:txNumber", handlers.GetData)
	r.GET("/proof/:assets", handlers.GetSpecifiedProof)
	r.GET("/proof", handlers.GetTotalProof)
	r.GET("/list", handlers.List)
	r.Run() // listen and serve on 0.0.0.0:8080
}
