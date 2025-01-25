package database

import (
	"go-graphql_galaxy/pkg/log"
	"strings"

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

func PreloadDB(db *gorm.DB, preloads [][]string) *gorm.DB {
	for _, preload := range preloads {
		preloadStr := strings.Join(preload, ".")
		db = db.Preload(preloadStr)
	}
	return db
}

func LimitOffsetDB(db *gorm.DB, limit, offset *int32) *gorm.DB {
	if limit != nil {
		db = db.Limit(int(*limit))
		if offset != nil {
			db = db.Offset(int(*offset))
		}
	}
	return db
}
