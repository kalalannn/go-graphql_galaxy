package transformers

import (
	"go-graphql_galaxy/internal/gorm/entities"
	"go-graphql_galaxy/internal/graphql/models"
)

func TransformSecretEntityToModel(secret *entities.SecretEntity) *models.Secret {
	return &models.Secret{
		ID:         secret.ID,
		SecretCode: secret.SecretCode,
		Nemesis:    TransformNemesisEntityToModel(&secret.Nemesis),
	}
}

func TransformSecretEntitiesToModels(secrets []*entities.SecretEntity) []*models.Secret {
	var modelSecrets []*models.Secret
	for _, secret := range secrets {
		modelSecrets = append(modelSecrets, TransformSecretEntityToModel(secret))
	}
	return modelSecrets
}
