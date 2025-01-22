package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"
	"go-graphql_galaxy/graph/model"
	"go-graphql_galaxy/internal/models"
	"time"
)

// Characters is the resolver for the characters field.
func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	var characters []*models.Character
	if err := r.DB.Preload("Nemeses.Secrets").Find(&characters).Error; err != nil {
		return nil, err
	}
	var modelCharacters []*model.Character
	for _, character := range characters {
		modelCharacters = append(modelCharacters, &model.Character{
			ID:              character.ID,
			Name:            character.Name,
			Gender:          character.Gender,
			Ability:         character.Ability,
			MinimalDistance: character.MinimalDistance,
			Weight:          character.Weight,
			Born:            character.Born.Format(time.RFC3339),
			InSpaceSince: func() *string {
				if character.InSpaceSince == nil {
					return nil
				}
				inSpaceSince := character.InSpaceSince.Format(time.RFC3339)
				return &inSpaceSince
			}(),
			BeerConsumption: int32(character.BeerConsumption),
			KnowsTheAnswer:  character.KnowsTheAnswer,
		})
	}
	return modelCharacters, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id uint) (*model.Character, error) {
	panic(fmt.Errorf("not implemented: Character - character"))
}

// Nemeses is the resolver for the nemeses field.
func (r *queryResolver) Nemeses(ctx context.Context) ([]*model.Nemesis, error) {
	panic(fmt.Errorf("not implemented: Nemeses - nemeses"))
}

// Secret is the resolver for the secret field.
func (r *queryResolver) Secret(ctx context.Context, id uint) (*model.Secret, error) {
	panic(fmt.Errorf("not implemented: Secret - secret"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
