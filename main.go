package main

import (
	"github.com/BANKEX/poa-history/config"
	"github.com/BANKEX/poa-history/db"
	"github.com/BANKEX/poa-history/server"
	"log"
)

func main() {

	// @title Swagger History API
	// @version 1.0
	// @description This is POA History swagger documentation

	// @contact.name API Support
	// @contact.email nk@bankexfoundation.org

	// @license.name MIT
	// @license.url https://opensource.org/licenses/MIT

	// @host history.bankex.team:8080
	// @BasePath /

	cfg := config.GetConfig()

	err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	server.Serve(s)

}
