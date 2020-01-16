package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/models"
)

func (r *queryResolver) Coffee(ctx context.Context, coffeeID string) (*models.Coffee, error) {
	return nil, nil
}

func (r *queryResolver) Coffees(ctx context.Context) ([]*models.Coffee, error) {
	coffees, err := r.CoffeeService.FindCoffees()
	if err != nil {
		return nil, err
	}
	return coffees, nil
}
