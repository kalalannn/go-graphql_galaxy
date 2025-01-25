package resolvers

import (
	"context"
	"fmt"
	"go-graphql_galaxy/internal/gorm/services"
	"go-graphql_galaxy/internal/gqlcontext"

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

func OrderBy(field, direction string) string {
	return fmt.Sprintf("%s %s", field, direction)
}
