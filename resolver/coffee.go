package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/models"
)

func (r *QueryResolver) Coffees(ctx context.Context) ([]*models.Coffee, error) {
	coffees, err := r.ProductService.GetCoffees()
	if err != nil {
		return nil, err
	}
	return coffees, nil
}
