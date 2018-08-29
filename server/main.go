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

	r.GET("/proof/:assetId/:txNumber/:hash/:timestamp", handlers.GetTotalProof)
	r.GET("/list", handlers.List)
	r.Static("/assets/main", "./assets/main")
	r.Static("/assets/upload", "./assets/upload")
	r.Static("/assets/download", "./assets/download")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "http://ec2-18-210-150-89.compute-1.amazonaws.com:80/assets/main/")
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
