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

func (s *NemesisService) NemesesCount() int64 {
	var count int64
	s.db.Model(&entities.NemesisEntity{}).Count(&count)
	return count
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

func (s *NemesisService) AverageNemesesYears() float64 {
	var averageYears float64
	query := `
		SELECT AVG(years)
		FROM ` + entities.NemesisTableName
	s.db.Raw(query).Scan(&averageYears)
	return averageYears
}

func (s *NemesisService) AliveNemeses() *entities.AliveNemeses {
	aliveNemeses := entities.AliveNemeses{}
	query := `
		SELECT 
			COUNT(CASE WHEN is_alive = 't' THEN 1 END) AS alive,
			COUNT(CASE WHEN is_alive = 'f' THEN 1 END) AS dead
		FROM ` + entities.NemesisTableName
	s.db.Raw(query).Scan(&aliveNemeses)
	return &aliveNemeses
}
