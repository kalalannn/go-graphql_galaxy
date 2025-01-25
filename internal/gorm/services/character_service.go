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

func (s *CharacterService) CharactersCount() int64 {
	var count int64
	s.db.Model(&entities.CharacterEntity{}).Count(&count)
	return count
}

func (s *CharacterService) Characters(orderBy string) ([]*entities.CharacterEntity, error) {
	var characters []*entities.CharacterEntity
	if err := s.db.Order(orderBy).Find(&characters).Error; err != nil {
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

func (s *CharacterService) AverageAge() float64 {
	var averageAge float64
	query := `
		SELECT AVG(EXTRACT(YEAR FROM AGE(born))) AS avg_age
		FROM ` + entities.CharacterTableName
	s.db.Raw(query).Scan(&averageAge)
	return averageAge
}

func (s *CharacterService) AverageWeight() float64 {
	var averageWeight float64
	query := `
		SELECT AVG(weight) AS avg_weight
		FROM ` + entities.CharacterTableName
	s.db.Raw(query).Scan(&averageWeight)
	return averageWeight
}

func (s *CharacterService) AverageBeerConsumption() float64 {
	var averageBeerConsumption float64
	query := `
		SELECT AVG(beer_consumption) AS beer_consumption
		FROM ` + entities.CharacterTableName
	s.db.Raw(query).Scan(&averageBeerConsumption)
	return averageBeerConsumption
}

func (s *CharacterService) Genders() *entities.Genders {
	genders := entities.Genders{}
	query := `
		SELECT
		SUM(CASE 
			WHEN LOWER(gender) IN ('m', 'male') THEN 1 
			ELSE 0 
		END) AS male,
		SUM(CASE 
			WHEN LOWER(gender) IN ('f', 'female') THEN 1 
			ELSE 0 
		END) AS female,
		SUM(CASE 
			WHEN LOWER(gender) NOT IN ('m', 'male', 'f', 'female') OR gender IS NULL THEN 1 
			ELSE 0 
		END) AS other
		FROM ` + entities.CharacterTableName
	s.db.Raw(query).Scan(&genders)
	return &genders
}
