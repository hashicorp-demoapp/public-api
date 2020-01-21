package models

import (
	"strconv"
	"testing"

	productsapi "github.com/hashicorp-demoapp/product-api-go/data/model"
	"github.com/stretchr/testify/assert"
)

func TestConvertsFromProductsAPI(t *testing.T) {
	c, err := CoffeeFromProductsAPI(apiModel)
	assert.NoError(t, err)

	assert.Len(t, c, 2)
	assert.Equal(t, strconv.Itoa(apiModel[0].ID), c[0].ID)
	assert.Equal(t, apiModel[0].Name, c[0].Name)
	assert.Equal(t, apiModel[0].Teaser, c[0].Teaser)
	assert.Equal(t, apiModel[0].Description, c[0].Description)
	assert.Equal(t, float64(apiModel[0].Price), c[0].Price)

	assert.Len(t, c[0].Ingredients, 2)
	assert.Equal(t, strconv.Itoa(apiModel[0].Ingredients[0].IngredientID), c[0].Ingredients[0].ID)
}

var apiModel = []productsapi.Coffee{
	productsapi.Coffee{
		ID:          123,
		Name:        "Latte",
		Teaser:      "This is a teaser",
		Description: "This is the description",
		Price:       220,
		Ingredients: []productsapi.CoffeeIngredients{
			productsapi.CoffeeIngredients{
				IngredientID: 1,
			},
			productsapi.CoffeeIngredients{
				IngredientID: 2,
			},
		},
	},
	productsapi.Coffee{
		ID:          125,
		Name:        "Espresso",
		Teaser:      "This is a teaser",
		Description: "This is the description",
		Price:       200,
		Ingredients: []productsapi.CoffeeIngredients{
			productsapi.CoffeeIngredients{
				IngredientID: 1,
			},
		},
	},
}
