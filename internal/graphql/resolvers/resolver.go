package resolvers

import (
	"context"
	"go-graphql_galaxy/internal/gqlcontext"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

func GetPreloads(ctx context.Context) [][]string {
	return ctx.Value(gqlcontext.PreloadContextKey).([][]string)
}
