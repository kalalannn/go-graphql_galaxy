package transformers

import (
	"go-graphql_galaxy/internal/gorm/entities"
	"go-graphql_galaxy/internal/graphql/models"
)

func TransformNemesisEntityToModel(nemesis *entities.NemesisEntity) *models.Nemesis {
	return &models.Nemesis{
		ID:        nemesis.ID,
		IsAlive:   nemesis.IsAlive,
		Years:     nemesis.Years,
		Character: TransformCharacterEntityToModel(&nemesis.Character),
		Secrets:   TransformSecretEntitiesToModels(nemesis.Secrets),
	}
}

func TransformNemesisEntitiesToModels(nemeses []*entities.NemesisEntity) []*models.Nemesis {
	var modelNemeses []*models.Nemesis
	for _, nemesis := range nemeses {
		modelNemeses = append(modelNemeses, TransformNemesisEntityToModel(nemesis))
	}
	return modelNemeses
}
