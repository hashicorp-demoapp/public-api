package resolver

import (
	"context"
	"strconv"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp-demoapp/public-api/auth"
	"github.com/hashicorp-demoapp/public-api/models"
)

func (r *QueryResolver) Orders(ctx context.Context) ([]*models.Order, error) {
	authToken := auth.GetAuthHeader(ctx)

	orders, err := r.ProductService.GetAllOrders(authToken)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *QueryResolver) Order(ctx context.Context, orderID string) (*models.Order, error) {
	authToken := auth.GetAuthHeader(ctx)

	order, err := r.ProductService.GetOrder(orderID, authToken)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *MutationResolver) Order(ctx context.Context, items []*models.OrderItemInput) (*models.Order, error) {
	authToken := auth.GetAuthHeader(ctx)

	// Convert models.OrderItems to HashiCups API order items
	ois := make([]hashicups.OrderItem, 0)

	for _, item := range items {
		oId, err := strconv.Atoi(item.Coffee.ID)
		if err != nil {
			return nil, err
		}

		oi := hashicups.OrderItem{
			Coffee: hashicups.Coffee{
				ID: oId,
			},
			Quantity: item.Quantity,
		}

		ois = append(ois, oi)
	}

	order, err := r.ProductService.CreateOrder(ois, authToken)
	if err != nil {
		return nil, err
	}
	return order, nil
}
