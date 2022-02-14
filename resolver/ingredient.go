package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/models"
)

func (r *QueryResolver) CoffeeIngredients(ctx context.Context, coffeeID string) ([]*models.Ingredient, error) {
	ings, err := r.ProductService.GetCoffeeIngredients(coffeeID)
	if err != nil {
		return nil, err
	}
	return ings, nil
}
