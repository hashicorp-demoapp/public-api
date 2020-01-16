package service

import (
	"time"
	
	"github.com/hashicorp-demoapp/public-api/models"
)

// CoffeeService handles interaction with coffees.
type CoffeeService struct {}

// NewCoffeeService creates a new CoffeeService.
func NewCoffeeService() *CoffeeService {
	return &CoffeeService{}
}

// FindCoffees returns a list of tracks.
func (s *CoffeeService) FindCoffees() ([]*models.Coffee, error) {
	coffees := []*models.Coffee{
		&models.Coffee{
			ID: "001",
			Name: "espresso",
			Image: "frappuccino",
			Price: 2.50,
			Ingredients: []models.Ingredient{
				models.Ingredient{
					ID: "001", 
				},
				models.Ingredient{
					ID: "002",
				},
			},
		},
		&models.Coffee{
			ID: "002",
			Name: "latte",
			Image: "frappuccino",
			Price: 3.50,
			Ingredients: []models.Ingredient{
				models.Ingredient{
					ID: "001", 
				},
				models.Ingredient{
					ID: "002",
				},
				models.Ingredient{
					ID: "003",
				},
			},
		},
		&models.Coffee{
			ID: "003",
			Name: "cappuccino",
			Image: "frappuccino",
			Price: 4.50,
			Ingredients: []models.Ingredient{
				models.Ingredient{
					ID: "001", 
				},
				models.Ingredient{
					ID: "002",
				},
				models.Ingredient{
					ID: "003", 
				},
				models.Ingredient{
					ID: "004",
				},
			},
		},
		&models.Coffee{
			ID: "004",
			Name: "espresso",
			Image: "frappuccino",
			Price: 2.50,
			Ingredients: []models.Ingredient{
				models.Ingredient{
					ID: "001", 
				},
				models.Ingredient{
					ID: "002",
				},
			},
		},
		&models.Coffee{
			ID: "005",
			Name: "latte",
			Image: "frappuccino",
			Price: 3.50,
			Ingredients: []models.Ingredient{
				models.Ingredient{
					ID: "001", 
				},
				models.Ingredient{
					ID: "002",
				},
				models.Ingredient{
					ID: "003",
				},
			},
		},
		&models.Coffee{
			ID: "006",
			Name: "cappuccino",
			Image: "frappuccino",
			Price: 4.50,
			Ingredients: []models.Ingredient{
				models.Ingredient{
					ID: "001", 
				},
				models.Ingredient{
					ID: "002",
				},
				models.Ingredient{
					ID: "003", 
				},
				models.Ingredient{
					ID: "004",
				},
			},
		},
	}

	time.Sleep(4 * time.Second)

	return coffees, nil
}