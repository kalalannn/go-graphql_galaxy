package services

import (
	"go-graphql_galaxy/internal/gorm/entities"
	"go-graphql_galaxy/pkg/database"

	"gorm.io/gorm"
)

type SecretService struct {
	db *gorm.DB
}

func NewSecretService(db *gorm.DB) *SecretService {
	return &SecretService{
		db: db,
	}
}

func (s *SecretService) SecretsCount() int64 {
	var count int64
	s.db.Model(&entities.SecretEntity{}).Count(&count)
	return count
}

func (s *SecretService) Secrets(preloads [][]string, orderBy string, limit, offset *int32) ([]*entities.SecretEntity, error) {
	db := database.LimitOffsetDB(database.PreloadDB(s.db, preloads), limit, offset)

	var secrets []*entities.SecretEntity
	if err := db.Order(orderBy).Find(&secrets).Error; err != nil {
		return nil, err
	}
	return secrets, nil
}

func (s *SecretService) Secret(id uint, preloads [][]string) (*entities.SecretEntity, error) {
	db := database.PreloadDB(s.db, preloads)
	var secret entities.SecretEntity
	if err := db.First(&secret, id).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}
