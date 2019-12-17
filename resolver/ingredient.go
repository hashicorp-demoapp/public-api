package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/models"
)

func (r *queryResolver) Ingredient(ctx context.Context, ingredientID string) (*models.Ingredient, error) {
	return nil, nil
}

func (r *queryResolver) Ingredients(ctx context.Context) ([]*models.Ingredient, error) {
	return []*models.Ingredient{}, nil
}
