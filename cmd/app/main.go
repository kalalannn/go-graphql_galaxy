package main

import (
	"fmt"
	"go-graphql_galaxy/internal/models"
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

	var characters []models.Character
	err = db.Preload("Nemeses.Secrets").Find(&characters).Error
	if err != nil {
		log.Fatal("failed to load characters:", err)
	}

	for _, character := range characters {
		fmt.Printf("Character: %s\n", character.Name)
		for _, nemesis := range character.Nemeses {
			fmt.Printf("\tNemesis: %v\n", nemesis)
			for _, secret := range nemesis.Secrets {
				fmt.Printf("\t\tSecret Code: %d\n", secret.SecretCode)
			}
		}
	}
}
