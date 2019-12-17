package resolver

import (
	"context"

	"github.com/hashicorp-demoapp/public-api/server"
	// "github.com/hashicorp-demoapp/public-api/service"
)

// Resolver is the grapqhl root resolver.
// Add services here for convenient access in other resolvers.
type Resolver struct {
	// CoffeeService     *service.CoffeeService
	// IngredientService *service.IngredientService
	// UserService *service.UserService
}

// Mutation handles graphql mutations.
func (r *Resolver) Mutation() server.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Login(ctx context.Context) (bool, error) {
	return true, nil
}

// Query handles graphql queries.
func (r *Resolver) Query() server.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return "1", nil
}
