package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp-demoapp/public-api/models"
)

// ProductService handles interaction with the product API.
type ProductService struct {
	c *hashicups.Client
}

// NewProductService creates a new ProductService.
func NewProductService(c *hashicups.Client) *ProductService {
	return &ProductService{c}
}

func (s *ProductService) HealthCheck() bool {
	resp, err := s.c.HTTPClient.Get(fmt.Sprintf("%s/health/readyz", s.c.HostURL))
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

// SignUp - Create new user, return user token upon successful creation
func (s *ProductService) SignUp(auth models.UserAuth) (models.AuthResponse, error) {
	pUserAuth := hashicups.AuthStruct{
		Username: auth.Username,
		Password: auth.Password,
	}

	pAuthResp, err := s.c.SignUp(pUserAuth)
	if err != nil {
		return models.AuthResponse{}, err
	}

	authResp := models.AuthResponse{
		UserID:   pAuthResp.UserID,
		Username: pAuthResp.Username,
		Token:    pAuthResp.Token,
	}

	return authResp, nil
}

// SignIn Get a new token for user.
func (s *ProductService) SignIn(auth models.UserAuth) (models.AuthResponse, error) {
	pUserAuth := hashicups.AuthStruct{
		Username: auth.Username,
		Password: auth.Password,
	}

	pAuthResp, err := s.c.GetUserTokenSignIn(pUserAuth)
	if err != nil {
		return models.AuthResponse{}, err
	}

	authResp := models.AuthResponse{
		UserID:   pAuthResp.UserID,
		Username: pAuthResp.Username,
		Token:    pAuthResp.Token,
	}

	return authResp, nil
}

// SignOut - Revoke the token for a user
func (s *ProductService) SignOut(authToken *string) error {
	return s.c.SignOut(authToken)
}

// GetCoffees returns a list of coffees.
func (s *ProductService) GetCoffees() ([]*models.Coffee, error) {
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

// GetCoffee returns a specific coffee (list).
func (s *ProductService) GetCoffee(coffeeID string) (*models.Coffee, error) {
	cofs, err := s.c.GetCoffee(coffeeID)
	if err != nil {
		return nil, err
	}

	coffees, err := models.CoffeeFromProductsAPI(cofs)
	if err != nil {
		return nil, err
	}

	if len(coffees) < 1 {
		return nil, errors.New("No coffee with that ID")
	}

	return coffees[0], nil
}

// GetCoffeeIngredients - Returns list of coffee ingredients (no auth required)
func (s *ProductService) GetCoffeeIngredients(coffeeID string) ([]*models.Ingredient, error) {
	pIngs, err := s.c.GetCoffeeIngredients(coffeeID)
	if err != nil {
		return nil, err
	}

	ings, err := models.IngredientsFromProductAPI(pIngs)
	if err != nil {
		return nil, err
	}

	return ings, nil
}

// GetAllOrders - Get all user's order
func (s *ProductService) GetAllOrders(authToken string) ([]*models.Order, error) {
	pOrders, err := s.c.GetAllOrders(&authToken)
	if err != nil {
		return nil, err
	}

	orders, err := models.OrdersFromProductsAPI(pOrders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrder - Get specific user order
func (s *ProductService) GetOrder(orderID string, authToken string) (*models.Order, error) {
	pOrder, err := s.c.GetOrder(orderID, &authToken)
	if err != nil {
		return nil, err
	}

	order, err := models.OrderFromProductsAPI(pOrder)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// CreateOrder - Create new order
func (s *ProductService) CreateOrder(ois []hashicups.OrderItem, authToken string) (*models.Order, error) {
	pOrder, err := s.c.CreateOrder(ois, &authToken)
	if err != nil {
		return nil, err
	}

	order, err := models.OrderFromProductsAPI(pOrder)
	if err != nil {
		return nil, err
	}

	return order, nil
}
