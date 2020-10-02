package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/models"
)

func (r *QueryResolver) Ingredient(ctx context.Context, ingredientID string) (*models.Ingredient, error) {
	return nil, nil
}

func (r *QueryResolver) Ingredients(ctx context.Context) ([]*models.Ingredient, error) {
	return []*models.Ingredient{}, nil
}
