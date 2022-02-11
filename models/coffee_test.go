package models

import (
	"strconv"
	"testing"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/stretchr/testify/assert"
)

func TestCoffeeConvertsFromProductsAPI(t *testing.T) {
	c, err := CoffeeFromProductsAPI(apiCoffeeModel)
	assert.NoError(t, err)

	assert.Len(t, c, 2)
	assert.Equal(t, strconv.Itoa(apiCoffeeModel[0].ID), c[0].ID)
	assert.Equal(t, apiCoffeeModel[0].Name, *c[0].Name)
	assert.Equal(t, apiCoffeeModel[0].Teaser, *c[0].Teaser)
	assert.Equal(t, apiCoffeeModel[0].Description, *c[0].Description)
	assert.Equal(t, float64(apiCoffeeModel[0].Price), *c[0].Price)

	assert.Len(t, c[0].Ingredients, 2)
	assert.Equal(t, strconv.Itoa(apiCoffeeModel[0].Ingredient[0].ID), c[0].Ingredients[0].ID)
}

var apiCoffeeModel = []hashicups.Coffee{
	{
		ID:          123,
		Name:        "Latte",
		Teaser:      "This is a teaser",
		Collection:  "Origins",
		Origin:      "Summer 2014",
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
	{
		ID:          125,
		Name:        "Espresso",
		Teaser:      "This is a teaser",
		Collection:  "Foundations",
		Origin:      "Fall 2015",
		Description: "This is the description",
		Price:       200,
		Ingredient: []hashicups.CoffeeIngredient{
			{
				ID: 1,
			},
		},
	},
}
