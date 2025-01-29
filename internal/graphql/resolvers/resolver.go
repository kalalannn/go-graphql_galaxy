package resolvers

import (
	"context"
	"fmt"
	"go-graphql_galaxy/internal/gorm/services"
	"go-graphql_galaxy/internal/gqlcontext"
	"go-graphql_galaxy/internal/graphql/models"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CharacterService *services.CharacterService
	NemesisService   *services.NemesisService
	SecretService    *services.SecretService
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		CharacterService: services.NewCharacterService(db),
		NemesisService:   services.NewNemesisService(db),
		SecretService:    services.NewSecretService(db),
	}
}

func EmptyPreloads() [][]string {
	return [][]string{}
}

func GetPreloads(ctx context.Context) [][]string {
	return ctx.Value(gqlcontext.PreloadContextKey).([][]string)
}

func GetOrderBy(field, direction string) string {
	return fmt.Sprintf("%s %s", field, direction)
}

// func GetPagination(limit, offset *int32) (*int32, *int32, error) {
func GetPagination(pagination *models.PaginationInput) (*int32, *int32, error) {
	if pagination == nil {
		return nil, nil, nil
	}
	limit, offset := pagination.Limit, pagination.Offset
	if offset != nil && limit == nil {
		return nil, nil, fmt.Errorf("limit is required when offset is provided")
	}
	var zero int32 = 0
	if limit != nil && offset == nil {
		offset = &zero
	}
	return limit, offset, nil
}
