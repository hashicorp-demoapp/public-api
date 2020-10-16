package service

import (
	"github.com/hashicorp-demoapp/product-api-go/client"
	"github.com/hashicorp-demoapp/public-api/models"
)

// CoffeeService handles interaction with coffees.
type CoffeeService struct {
	c *client.HTTP
}

// NewCoffeeService creates a new CoffeeService.
func NewCoffeeService(c *client.HTTP) *CoffeeService {
	return &CoffeeService{c}
}

// FindCoffees returns a list of coffees.
func (s *CoffeeService) FindCoffees() ([]*models.Coffee, error) {
	cofs, err := s.c.GetCoffees()
	if err != nil {
		return nil, err
	}

	coffees, err := models.CoffeeFromProductsAPI(cofs)
	if err != nil {
		return nil, err
	}

	return coffees, nil
}
