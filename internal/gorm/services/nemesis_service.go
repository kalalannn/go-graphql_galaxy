package services

import (
	"go-graphql_galaxy/internal/gorm/entities"

	"gorm.io/gorm"
)

type NemesisService struct {
	db *gorm.DB
}

func NewNemesisService(db *gorm.DB, preloads [][]string) *NemesisService {
	s := NemesisService{}
	s.db = PreloadDB(db, preloads)
	return &s
}

func (s *NemesisService) Nemeses() ([]*entities.NemesisEntity, error) {
	var nemeses []*entities.NemesisEntity
	if err := s.db.Find(&nemeses).Error; err != nil {
		return nil, err
	}
	return nemeses, nil
}

func (s *NemesisService) Nemesis(id uint) (*entities.NemesisEntity, error) {
	var nemesis entities.NemesisEntity
	if err := s.db.First(&nemesis, id).Error; err != nil {
		return nil, err
	}
	return &nemesis, nil
}
