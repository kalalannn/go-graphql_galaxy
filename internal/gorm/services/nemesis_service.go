package services

import (
	"go-graphql_galaxy/internal/gorm/entities"
	"go-graphql_galaxy/pkg/database"

	"gorm.io/gorm"
)

type NemesisService struct {
	db *gorm.DB
}

func NewNemesisService(db *gorm.DB) *NemesisService {
	return &NemesisService{
		db: db,
	}
}

func (s *NemesisService) NemesesCount() int64 {
	var count int64
	s.db.Model(&entities.NemesisEntity{}).Count(&count)
	return count
}

func (s *NemesisService) Nemeses(preloads [][]string) ([]*entities.NemesisEntity, error) {
	db := database.PreloadDB(s.db, preloads)
	var nemeses []*entities.NemesisEntity
	if err := db.Find(&nemeses).Error; err != nil {
		return nil, err
	}
	return nemeses, nil
}

func (s *NemesisService) Nemesis(id uint, preloads [][]string) (*entities.NemesisEntity, error) {
	db := database.PreloadDB(s.db, preloads)
	var nemesis entities.NemesisEntity
	if err := db.First(&nemesis, id).Error; err != nil {
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
