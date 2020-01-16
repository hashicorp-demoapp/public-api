package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/hashicorp-demoapp/public-api/auth"
	"github.com/hashicorp-demoapp/public-api/models"
	"github.com/hashicorp-demoapp/public-api/resolver"
	"github.com/hashicorp-demoapp/public-api/server"
	// "github.com/hashicorp-demoapp/public-api/service"
	"github.com/gorilla/mux"
	"github.com/keratin/authn-go/authn"
)

func main() {
	// Config.
	// config := service.NewConfig()

	// Authentication.
	authn, err := authn.NewClient(authn.Config{
		// The AUTHN_URL of your Keratin AuthN server. This will be used to verify tokens created by
		// AuthN, and will also be used for API calls unless PrivateBaseURL is also set.
		Issuer:         "http://localhost",

		// The domain of your application (no protocol). This domain should be listed in the APP_DOMAINS
		// of your Keratin AuthN server.
		Audience:       "localhost",

		// Credentials for AuthN's private endpoints. These will be used to execute admin actions using
		// the Client provided by this library.
		//
		// TIP: make them extra secure in production!
		Username:       "hello",
		Password:       "world",

		// RECOMMENDED: Send private API calls to AuthN using private network routing. This can be
		// necessary if your environment has a firewall to limit public endpoints.
		PrivateBaseURL: "http://localhost",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Server.
	r := mux.NewRouter()
	r.Use(auth.Middleware(authn))

	// Graphql.
	c := server.Config{
		Resolvers: &resolver.Resolver{},
	}

	// Check if the user is authenticated.
	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		isAuthenticated := auth.IsAuthenticated(ctx)
		if !isAuthenticated {
			return nil, fmt.Errorf("Access denied")
		}

		return next(ctx)
	}

	// Check if the user has a role.
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (interface{}, error) {
		log.Printf("%#v", role)
		return next(ctx)
	}

	// Handlers.
	r.Handle("/", handler.Playground("Playground", "/api"))
	r.Handle("/api", handler.GraphQL(server.NewExecutableSchema(c)))

	log.Println("Listing on port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
