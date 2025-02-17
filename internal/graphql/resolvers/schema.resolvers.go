package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"go-graphql_galaxy/internal/graphql/generated"
	"go-graphql_galaxy/internal/transformers"
	"time"
)

// ServerTime is the resolver for the server_time field.
func (r *queryResolver) ServerTime(ctx context.Context) (string, error) {
	return *transformers.TransformTimeToString(time.Now()), nil
}

// HealthCheck is the resolver for the health_check field.
func (r *queryResolver) HealthCheck(ctx context.Context) (bool, error) {
	return true, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
