package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"go-graphql_galaxy/internal/graphql/models"
	"go-graphql_galaxy/internal/transformers"
)

// SecretsCount is the resolver for the secrets_count field.
func (r *queryResolver) SecretsCount(ctx context.Context) (int64, error) {
	return r.SecretService.SecretsCount(), nil
}

// Secrets is the resolver for the secrets field.
func (r *queryResolver) Secrets(ctx context.Context, orderBy *models.SecretOrderBy, pagination *models.PaginationInput) ([]*models.Secret, error) {
	limit, offset, err := GetPagination(pagination)
	if err != nil {
		return nil, err
	}

	secretEntities, err := r.SecretService.Secrets(
		GetPreloads(ctx),
		GetOrderBy(orderBy.Field.String(), orderBy.Direction.String()),
		limit, offset,
	)
	if err != nil {
		return nil, err //! check business error
	}

	return transformers.TransformSecretEntitiesToModels(secretEntities), nil
}

// Secret is the resolver for the secret field.
func (r *queryResolver) Secret(ctx context.Context, id uint) (*models.Secret, error) {
	secretEntity, err := r.SecretService.Secret(id, GetPreloads(ctx))
	if err != nil {
		return nil, err //! check business error
	}
	return transformers.TransformSecretEntityToModel(secretEntity), nil
}
