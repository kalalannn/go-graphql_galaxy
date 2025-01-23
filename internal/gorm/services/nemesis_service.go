package services

import (
	"go-graphql_galaxy/internal/gorm/entities"

	"gorm.io/gorm"
)

type NemesisService struct {
	db *gorm.DB
}

func NewNemesisService(db *gorm.DB) *NemesisService {
	return &NemesisService{db: db}
}

func (s *NemesisService) Nemeses() ([]*entities.NemesisEntity, error) {
	var nemeses []*entities.NemesisEntity
	if err := s.db.Preload("Secrets").Find(&nemeses).Error; err != nil {
		return nil, err
	}
	return nemeses, nil
}
