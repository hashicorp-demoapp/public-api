package models

import (
	"strconv"

	"github.com/hashicorp-demoapp/hashicups-client-go"
)

// CoffeeFromProductsAPI is an adaptor function which converts the products api model
// into the local model
func CoffeeFromProductsAPI(cof []hashicups.Coffee) ([]*Coffee, error) {
	pc := make([]*Coffee, 0)

	for i, co := range cof {
		price := float64(co.Price)
		c := &Coffee{
			ID:          strconv.Itoa(co.ID),
			Name:        &cof[i].Name,
			Price:       &price,
			Teaser:      &cof[i].Teaser,
			Collection:  &cof[i].Collection,
			Origin:      &cof[i].Origin,
			Description: &cof[i].Description,
			Image:       &cof[i].Image,
		}

		// add the ingredients
		ins := make([]*Ingredient, 0)
		for _, i := range co.Ingredient {
			ins = append(ins, &Ingredient{
				ID: strconv.Itoa(i.ID),
			})
		}

		c.Ingredients = ins
		pc = append(pc, c)
	}

	return pc, nil
}
