package models

import (
	"strconv"

	productsapi "github.com/hashicorp-demoapp/product-api-go/data/model"
)

// CoffeeFromProductsAPI is an adaptor function which converts the products api model
// into the local model
func CoffeeFromProductsAPI(cof []productsapi.Coffee) ([]*Coffee, error) {
	pc := make([]*Coffee, 0)

	for i, co := range cof {
		price := float64(co.Price)
		c := &Coffee{
			ID:    strconv.Itoa(co.ID),
			Name:  &cof[i].Name,
			Price: &price,
			Image: &cof[i].Image,
		}

		// add the ingredients
		ins := make([]*Ingredient, 0)
		for _, i := range co.Ingredients {
			ins = append(ins, &Ingredient{
				ID: strconv.Itoa(i.IngredientID),
			})
		}

		c.Ingredients = ins
		pc = append(pc, c)
	}

	return pc, nil
}
