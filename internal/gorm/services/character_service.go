package services

import (
	"go-graphql_galaxy/internal/gorm/entities"

	"gorm.io/gorm"
)

type CharacterService struct {
	db *gorm.DB
}

func NewCharacterService(db *gorm.DB, preloads [][]string) *CharacterService {
	s := CharacterService{}
	s.db = PreloadDB(db, preloads)
	return &s
}

func (s *CharacterService) Characters() ([]*entities.CharacterEntity, error) {
	var characters []*entities.CharacterEntity
	if err := s.db.Find(&characters).Error; err != nil {
		return nil, err
	}
	return characters, nil
}

func (s *CharacterService) Character(id uint) (*entities.CharacterEntity, error) {
	var character entities.CharacterEntity
	if err := s.db.First(&character, id).Error; err != nil {
		return nil, err
	}
	return &character, nil
}
