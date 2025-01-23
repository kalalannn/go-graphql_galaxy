package transformers

import (
	"go-graphql_galaxy/internal/gorm/entities"
	"go-graphql_galaxy/internal/graphql/models"
)

func TransformCharacterEntityToModel(character *entities.CharacterEntity) *models.Character {
	return &models.Character{
		ID:              character.ID,
		Name:            character.Name,
		Gender:          character.Gender,
		Ability:         character.Ability,
		MinimalDistance: character.MinimalDistance,
		Weight:          character.Weight,
		Born:            *TransformTimeToString(character.Born),
		InSpaceSince:    TransformTimeToString(character.InSpaceSince),
		BeerConsumption: int32(character.BeerConsumption),
		KnowsTheAnswer:  character.KnowsTheAnswer,
		Nemeses:         TransformNemesisEntitiesToModels(character.Nemeses),
	}
}

func TransformCharacterEntitiesToModels(characters []*entities.CharacterEntity) []*models.Character {
	var modelCharacters []*models.Character
	for _, character := range characters {
		modelCharacters = append(modelCharacters, TransformCharacterEntityToModel(character))
	}
	return modelCharacters
}
