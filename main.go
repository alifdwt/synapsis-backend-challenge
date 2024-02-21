package main

import (
	"database/sql"
	"log"

	"github.com/alifdwt/synapsis-backend-challenge/api"
	db "github.com/alifdwt/synapsis-backend-challenge/db/sqlc"
	"github.com/alifdwt/synapsis-backend-challenge/util"
)

// @title Synapsis Backend Challenge
// @version 1.0
// @description This is an API for Synapsis Backend Challenge, and also an assignment for Backend Engineer Position at Synapsis.
// @termsOfService http://swagger.io/terms/

// @contact.name Alif Dewantara
// @contact.url http://github.com/alifdwt
// @contact.email aputradewantara@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @license.name Apache 2.0
// @licence.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
