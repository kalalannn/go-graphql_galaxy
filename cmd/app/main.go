package main

import (
	"go-graphql_galaxy/pkg/database"
	"go-graphql_galaxy/pkg/log"
	"go-graphql_galaxy/pkg/utils"
)

func main() {
	config := utils.MustLoadConfig()

	// Init logger
	log.Init(config.Env)

	// Connect to DB
	dsn := utils.DSN(&config.Database)
	db, err := database.ConnectDB(dsn)
	if err != nil {
		log.Fatal("failed to connect database: %v", err)
	}
	_ = db
}
