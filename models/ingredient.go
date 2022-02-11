package models

import (
	"log"
	"strconv"

	"github.com/hashicorp-demoapp/hashicups-client-go"
)

// IngredientsFromProductAPI is an adaptor function which converts the
// Ingredients API model into the local model
func IngredientsFromProductAPI(pIngs []hashicups.Ingredient) ([]*Ingredient, error) {
	ings := make([]*Ingredient, 0)

	for i, pIng := range pIngs {
		ing := &Ingredient{
			ID:       strconv.Itoa(pIng.ID),
			Name:     &pIngs[i].Name,
			Quantity: &pIngs[i].Quantity,
			Unit:     &pIngs[i].Unit,
		}

		log.Printf("%+v: %+v", ing, pIng)

		ings = append(ings, ing)
	}

	return ings, nil
}
