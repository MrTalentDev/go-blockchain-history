package main

import (
	"./db"
	"./middlewares"
	"./commonware/handlers"
	"github.com/gin-gonic/gin"
	"os"
)

var LOGIN = os.Getenv("LOGIN")
var PASSWORD = os.Getenv("PASSWORD")

func main() {

	// mongo init
	db.Connect()

	r := gin.Default()
	r.Use(middlewares.Connect)
	r.Use(middlewares.ErrorHandler)

	a := r.Group("/a", gin.BasicAuth(gin.Accounts{
		LOGIN: PASSWORD,
	}))
	
	a.POST("/new/:assetId/:hash", handlers.CreateAssetId)
	a.POST("/update/:assetId/:hash", handlers.UpdateAssetId)

	r.GET("/get/:assetId/:txNumber", handlers.GetData)

	r.GET("/proof", handlers.GetTotalProof)
	r.GET("/list", handlers.List)
	r.Run() // listen and serve on 0.0.0.0:8080
}
