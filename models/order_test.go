package models

import (
	"strconv"
	"testing"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/stretchr/testify/assert"
)

func TestOrderConvertsFromProductsAPI(t *testing.T) {
	o, err := OrderFromProductsAPI(&apiOrderModel)
	assert.NoError(t, err)

	assert.Equal(t, strconv.Itoa(apiOrderModel.ID), o.ID)
	assert.Len(t, o.Items, 2)

	// Test OrderItems
	assert.Equal(t, strconv.Itoa(apiOrderModel.Items[0].Coffee.ID), o.Items[0].Coffee.ID)
	assert.Equal(t, apiOrderModel.Items[0].Coffee.Name, *o.Items[0].Coffee.Name)
	assert.Equal(t, apiOrderModel.Items[0].Coffee.Teaser, *o.Items[0].Coffee.Teaser)
	assert.Equal(t, apiOrderModel.Items[0].Coffee.Description, *o.Items[0].Coffee.Description)
	assert.Equal(t, float64(apiOrderModel.Items[0].Coffee.Price), *o.Items[0].Coffee.Price)
	// Test OrderItem Ingredients
	assert.Len(t, apiOrderModel.Items[0].Coffee.Ingredient, 2)
	assert.Equal(t, strconv.Itoa(apiOrderModel.Items[0].Coffee.Ingredient[0].ID), o.Items[0].Coffee.Ingredients[0].ID)
	// Test OrderItem Quantity
	assert.Equal(t, apiOrderModel.Items[0].Quantity, *&o.Items[0].Quantity)
}

var apiOrderModel = hashicups.Order{
	ID: 123,
	Items: []hashicups.OrderItem{
		{
			Coffee: hashicups.Coffee{
				ID:          123,
				Name:        "Latte",
				Teaser:      "This is a teaser",
				Collection:  "Origins",
				Origin:      "Summer 2014",
				Color:       "#FFF",
				Description: "This is the description",
				Price:       220,
				Ingredient: []hashicups.CoffeeIngredient{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
				},
			},
			Quantity: 1,
		},
		{
			Coffee: hashicups.Coffee{
				ID:          125,
				Name:        "Espresso",
				Teaser:      "This is a teaser",
				Collection:  "Foundations",
				Origin:      "Fall 2015",
				Color:       "#000",
				Description: "This is the description",
				Price:       200,
				Ingredient: []hashicups.CoffeeIngredient{
					{
						ID: 1,
					},
				},
			},
			Quantity: 10,
		},
	},
}
