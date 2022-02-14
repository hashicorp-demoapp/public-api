package models

import (
	"strconv"
	"testing"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/stretchr/testify/assert"
)

func TestIngredientConvertsFromProductsAPI(t *testing.T) {
	i, err := IngredientsFromProductAPI(apiIngredientModel)
	assert.NoError(t, err)

	assert.Len(t, i, 2)
	assert.Equal(t, strconv.Itoa(apiIngredientModel[0].ID), i[0].ID)
	assert.Equal(t, apiIngredientModel[0].Name, *i[0].Name)
	assert.Equal(t, apiIngredientModel[0].Quantity, *i[0].Quantity)
	assert.Equal(t, apiIngredientModel[0].Unit, *i[0].Unit)
}

var apiIngredientModel = []hashicups.Ingredient{
	{
		ID:       1,
		Name:     "Espresso",
		Quantity: 40,
		Unit:     "ml",
	},
	{
		ID:       1,
		Name:     "Semi Skimmed Milk",
		Quantity: 20,
		Unit:     "ml",
	},
}
