package main

import (
	"go-graphql_galaxy/internal/server"
	"go-graphql_galaxy/pkg/database"
	"go-graphql_galaxy/pkg/log"
	"go-graphql_galaxy/pkg/utils"
)

func main() {
	config := utils.MustLoadConfig()

	// Init logger
	log.Init(config.Env)

	// Init and Connect to DB
	db, err := database.ConnectDB(utils.DSN(&config.Database))
	if err != nil {
		log.Fatal("failed to connect database: %v", err)
	}

	// Init Server
	serverService := server.NewServerService(&config.Server, db)

	// Serve
	serverService.RunServer()
}
