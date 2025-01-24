package services

import (
	"strings"

	"gorm.io/gorm"
)

func PreloadDB(db *gorm.DB, preloads [][]string) *gorm.DB {
	for _, preload := range preloads {
		preloadStr := strings.Join(preload, ".")
		db = db.Preload(preloadStr)
	}
	return db
}
