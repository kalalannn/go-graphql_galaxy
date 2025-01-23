package services

import (
	"go-graphql_galaxy/internal/gorm/entities"

	"gorm.io/gorm"
)

type CharacterService struct {
	db *gorm.DB
}

func NewCharacterService(db *gorm.DB) *CharacterService {
	return &CharacterService{db: db}
}

func (s *CharacterService) Characters() ([]*entities.CharacterEntity, error) {
	var characters []*entities.CharacterEntity
	if err := s.db.Preload("Nemeses.Secrets").Find(&characters).Error; err != nil {
		return nil, err
	}
	return characters, nil
}

func (s *CharacterService) Character(id uint) (*entities.CharacterEntity, error) {
	var character entities.CharacterEntity
	// if err := s.db.Preload("Nemeses.Secrets").Preload("Nemeses.Character").First(&character, id).Error; err != nil {
	if err := s.db.Preload("Nemeses.Secrets").First(&character, id).Error; err != nil {
		return nil, err
	}
	return &character, nil
}
