package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"go-graphql_galaxy/internal/graphql/models"
	"go-graphql_galaxy/internal/transformers"
)

// NemesesCount is the resolver for the nemeses_count field.
func (r *queryResolver) NemesesCount(ctx context.Context) (int64, error) {
	return r.NemesisService.NemesesCount(), nil
}

// Nemeses is the resolver for the nemeses field.
func (r *queryResolver) Nemeses(ctx context.Context, orderBy *models.NemesisOrderBy, pagination *models.PaginationInput) ([]*models.Nemesis, error) {
	limit, offset, err := GetPagination(pagination)
	if err != nil {
		return nil, err
	}

	nemesisEntities, err := r.NemesisService.Nemeses(
		GetPreloads(ctx),
		GetOrderBy(orderBy.Field.String(), orderBy.Direction.String()),
		limit, offset,
	)
	if err != nil {
		return nil, err //! check business error
	}

	return transformers.TransformNemesisEntitiesToModels(nemesisEntities), nil
}

// Nemesis is the resolver for the nemesis field.
func (r *queryResolver) Nemesis(ctx context.Context, id uint) (*models.Nemesis, error) {
	nemesisEntity, err := r.NemesisService.Nemesis(id, GetPreloads(ctx))
	if err != nil {
		return nil, err //! check business error
	}
	return transformers.TransformNemesisEntityToModel(nemesisEntity), nil
}

// AverageNemesesYears is the resolver for the average_nemeses_years field.
func (r *queryResolver) AverageNemesesYears(ctx context.Context) (float64, error) {
	return transformers.RoundFloat(r.NemesisService.AverageNemesesYears(), 2), nil
}

// AliveNemeses is the resolver for the alive_nemeses field.
func (r *queryResolver) AliveNemeses(ctx context.Context) (*models.AliveNemeses, error) {
	aliveNemeses := r.NemesisService.AliveNemeses()
	return &models.AliveNemeses{
		Alive: aliveNemeses.Alive,
		Dead:  aliveNemeses.Dead,
	}, nil
}
