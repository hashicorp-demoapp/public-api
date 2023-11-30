package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/hashicorp-demoapp/go-hckit"
	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp-demoapp/public-api/auth"
	"github.com/hashicorp-demoapp/public-api/handlers"
	"github.com/hashicorp-demoapp/public-api/models"
	"github.com/hashicorp-demoapp/public-api/payments"
	"github.com/hashicorp-demoapp/public-api/resolver"
	"github.com/hashicorp-demoapp/public-api/server"
	"github.com/hashicorp-demoapp/public-api/service"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"github.com/rs/cors"

	// "github.com/hashicorp-demoapp/public-api/service"
	"github.com/gorilla/mux"
	"github.com/keratin/authn-go/authn"
)

var logger hclog.Logger

var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind address for the server")
var metricsAddress = env.String("METRICS_ADDRESS", false, ":9102", "Metrics address for the server")
var productAddress = env.String("PRODUCT_API_URI", false, "http://localhost:9090", "Address for the product api")
var paymentAddress = env.String("PAYMENT_API_URI", false, "http://localhost:18000", "Address for the payment api")
var errorRate = env.Int("ERROR_RATE", false, 0, "Integer between 0 and 100 that represents how often the service should return 500 error code")

func main() {
	err := env.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// Config.
	// config := service.NewConfig()

	logger = hclog.New(&hclog.LoggerOptions{
		Name:  "public-api",
		Level: hclog.Debug,
	})

	closer, err := hckit.InitGlobalTracer("public-api")
	if err != nil {
		logger.Error("Unable to initialize Tracer", "error", err)
		os.Exit(1)
	}
	defer closer.Close()

	// Authentication.
	authn, err := authn.NewClient(authn.Config{
		// The AUTHN_URL of your Keratin AuthN server. This will be used to verify tokens created by
		// AuthN, and will also be used for API calls unless PrivateBaseURL is also set.
		Issuer: "http://localhost",

		// The domain of your application (no protocol). This domain should be listed in the APP_DOMAINS
		// of your Keratin AuthN server.
		Audience: "localhost",

		// Credentials for AuthN's private endpoints. These will be used to execute admin actions using
		// the Client provided by this library.
		//
		// TIP: make them extra secure in production!
		Username: "hello",
		Password: "world",

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
	errorsMiddleware := handlers.NewErrorMiddleware(*errorRate, logger)
	r.Use(errorsMiddleware.Middleware)
	r.Use(hckit.TracingMiddleware)

	// Enable CORS for all hosts
	r.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}).Handler)

	// create the client to the products-api
	productsClient, err := hashicups.NewClient(productAddress, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	productService := service.NewProductService(productsClient)

	// create the client for the payments-api
	paymentClient := payments.NewHTTP(*paymentAddress)

	// Graphql
	c := server.Config{
		Resolvers: resolver.NewResolver(productsClient, paymentClient, logger),
	}

	// Check if the user is authenticated.
	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		authHeader := auth.GetAuthHeader(ctx)
		if authHeader == "" {
			return nil, fmt.Errorf("Authorization header missing")
		}

		return next(ctx)
	}

	// Check if the user has a role.
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (interface{}, error) {
		logger.Debug("Auth has role", "role", role)
		return next(ctx)
	}

	// Handlers.
	r.Handle("/", handler.Playground("Playground", "/api"))
	r.Handle("/api", handler.GraphQL(server.NewExecutableSchema(c)))
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		var hcFailedServices []string
		if !productService.HealthCheck() {
			hcFailedServices = append(hcFailedServices, "product-api")
		}
		if !paymentClient.HealthCheck() {
			hcFailedServices = append(hcFailedServices, "payments")
		}
		if len(hcFailedServices) > 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error. the following services are down: %s", strings.Join(hcFailedServices, ", "))
			return
		}
		fmt.Fprintf(w, "%s", "ok")
	})

	logger.Info("Starting server", "bind", *bindAddress, "metrics", *metricsAddress)

	err = http.ListenAndServe(*bindAddress, r)
	if err != nil {
		logger.Error("Unable to start server", "error", err)
		os.Exit(1)
	}
}
