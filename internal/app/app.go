package app

import (
	"go-graphql_galaxy/internal/server"
	"go-graphql_galaxy/pkg/database"
	"go-graphql_galaxy/pkg/log"
	"go-graphql_galaxy/pkg/utils"
	"sync"

	"gorm.io/gorm"
)

type App struct {
	config *utils.Config
	db     *gorm.DB
}

func NewApp() *App {
	config := utils.MustLoadConfig()

	// Init logger
	log.Init(config.Env)

	// Init and Connect to DB
	db, err := database.ConnectDB(utils.DSN(&config.Database))
	if err != nil {
		log.Fatal("failed to connect database: %v", err)
	}

	return &App{
		config: config,
		db:     db,
	}
}

func (a *App) Run(wg *sync.WaitGroup) {
	// Init Server
	serverService := server.NewServerService(&a.config.Server, a.db)

	// Serve
	serverService.RunServer(wg)
}
