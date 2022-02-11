package models

import (
	"strconv"

	"github.com/hashicorp-demoapp/hashicups-client-go"
)

// OrdersFromProductsAPI is an adaptor function which converts the
// Orders API model into the local model
func OrdersFromProductsAPI(pOrders *[]hashicups.Order) ([]*Order, error) {
	orders := make([]*Order, 0)

	for _, po := range *pOrders {
		order, err := OrderFromProductsAPI(&po)
		if err != nil {
			return orders, nil
		}

		orders = append(orders, order)
	}
	return orders, nil
}

// OrderFromProductsAPI is an adaptor function which converts the
// Order API model into the local model
func OrderFromProductsAPI(po *hashicups.Order) (*Order, error) {
	order := &Order{
		ID: strconv.Itoa(po.ID),
	}

	// add orderItems
	ois := make([]*OrderItem, 0)
	for j, _ := range po.Items {
		oi := po.Items[j]

		price := float64(oi.Coffee.Price)
		c := &Coffee{
			ID:          strconv.Itoa(oi.Coffee.ID),
			Name:        &po.Items[j].Coffee.Name,
			Price:       &price,
			Teaser:      &po.Items[j].Coffee.Teaser,
			Collection:  &po.Items[j].Coffee.Collection,
			Origin:      &po.Items[j].Coffee.Origin,
			Description: &po.Items[j].Coffee.Description,
			Image:       &po.Items[j].Coffee.Image,
		}

		// add the ingredients
		ins := make([]*Ingredient, 0)
		for _, i := range po.Items[j].Coffee.Ingredient {
			ins = append(ins, &Ingredient{
				ID: strconv.Itoa(i.ID),
			})
		}

		c.Ingredients = ins

		ois = append(ois, &OrderItem{
			Coffee:   c,
			Quantity: oi.Quantity,
		})
	}

	order.Items = ois

	return order, nil
}
