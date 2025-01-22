package database

import (
	"go-graphql_galaxy/pkg/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect database: %v", err)
		return nil, err
	}
	log.Info("connected to database")
	return db, nil
}
