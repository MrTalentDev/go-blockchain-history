package main

import (
	"./db"
	"./middlewares"
	"./commonware/handlers"
	"github.com/gin-gonic/gin"
	"os"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "./docs"
	"github.com/gin-gonic/contrib/cors"
)

var LOGIN = os.Getenv("LOGIN")
var PASSWORD = os.Getenv("PASSWORD")

func main() {

	// @title Swagger History API
	// @version 1.0
	// @description This is POA History swagger documentation

	// @contact.name API Support
	// @contact.email nk@bankexfoundation.org

	// @license.name MIT
	// @license.url https://opensource.org/licenses/MIT

	// @host localhost:8080
	// @BasePath /

	db.Connect()

	r := gin.Default()

	r.Use(cors.Default())

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
