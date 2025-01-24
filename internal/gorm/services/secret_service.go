package services

import (
	"go-graphql_galaxy/internal/gorm/entities"

	"gorm.io/gorm"
)

type SecretService struct {
	db *gorm.DB
}

func NewSecretService(db *gorm.DB, preloads [][]string) *SecretService {
	s := SecretService{}
	s.db = PreloadDB(db, preloads)
	return &s
}

func (s *SecretService) Secrets() ([]*entities.SecretEntity, error) {
	var secrets []*entities.SecretEntity
	if err := s.db.Find(&secrets).Error; err != nil {
		return nil, err
	}
	return secrets, nil
}

func (s *SecretService) Secret(id uint) (*entities.SecretEntity, error) {
	var secret entities.SecretEntity
	if err := s.db.Find(&secret, id).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}
