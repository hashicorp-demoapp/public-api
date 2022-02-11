package resolver

import (
	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp-demoapp/public-api/payments"
	"github.com/hashicorp-demoapp/public-api/server"
	"github.com/hashicorp-demoapp/public-api/service"
	"github.com/hashicorp/go-hclog"
)

// Resolver is the grapqhl root resolver.
// Add services here for convenient access in other resolvers.
type Resolver struct {
	ProductService *service.ProductService
	PaymentService *service.PaymentService
	Log            hclog.Logger
	// IngredientService *service.IngredientService
	// UserService *service.UserService
}

func NewResolver(c *hashicups.Client, pc *payments.HTTPClient, l hclog.Logger) *Resolver {
	return &Resolver{
		service.NewProductService(c),
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
