package services

import (
	"go-graphql_galaxy/internal/gorm/entities"

	"gorm.io/gorm"
)

type SecretService struct {
	db *gorm.DB
}

func NewSecretService(db *gorm.DB) *SecretService {
	return &SecretService{db: db}
}

func (s *SecretService) Secrets() ([]*entities.SecretEntity, error) {
	var secrets []*entities.SecretEntity
	if err := s.db.Find(&secrets).Error; err != nil {
		return nil, err
	}
	return secrets, nil
}
