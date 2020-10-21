package resolver

import (
	"github.com/hashicorp-demoapp/product-api-go/client"
	"github.com/hashicorp-demoapp/public-api/payments"
	"github.com/hashicorp-demoapp/public-api/server"
	"github.com/hashicorp-demoapp/public-api/service"
	"github.com/hashicorp/go-hclog"
)

// Resolver is the grapqhl root resolver.
// Add services here for convenient access in other resolvers.
type Resolver struct {
	CoffeeService  *service.CoffeeService
	PaymentService *service.PaymentService
	Log            hclog.Logger
	// IngredientService *service.IngredientService
	// UserService *service.UserService
}

func NewResolver(c *client.HTTP, pc *payments.HTTPClient, l hclog.Logger) *Resolver {
	return &Resolver{
		service.NewCoffeeService(c),
		service.NewPaymentService(pc),
		l,
	}
}

// Mutation handles graphql mutations.
func (r *Resolver) Mutation() server.MutationResolver {
	return &MutationResolver{r}
}

// Query handles graphql queries.
func (r *Resolver) Query() server.QueryResolver {
	return &QueryResolver{r}
}
