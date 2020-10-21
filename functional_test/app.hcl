network "local" {
  subnet = "10.0.0.0/16"
}

container "products" {
    network {
        name = "network.local"
    }

    image {
        name = "hashicorpdemoapp/product-api:v4280cf7"
    }

    env_var = {
      CONFIG_FILE: "/config/config.json"
    }
    
    port {
      local = "9090"
      remote = "9090"
    }

    volume {
      source = "data.volume"
      destination = "/config"
      type = "volume"
    }
}

container "products_db" {
    network {
        name = "network.local"
    }

    image {
        name = "hashicorpdemoapp/product-api-db:v4280cf7"
    }

    env_var = {
      POSTGRES_DB: "products"
      POSTGRES_USER: "postgres"
      POSTGRES_PASSWORD: "password"
    }
}

container "payments" {
    network {
        name = "network.local"
    }

    image {
        name = "hashicorpdemoapp/payments:v0.0.3"
    }
    
    port {
      local = "8080"
      remote = "8080"
    }
}

container "public_api" {
    network {
        name = "network.local"
    }

    image {
        name = "hashicorpdemoapp/public-api:dev"
    }
    
    env_var = {
      PRODUCT_API_URI: "http://products.container.shipyard.run:9090"
      PAYMENT_API_URI: "http://payments.container.shipyard.run:8080"
    }

    port {
      local = "8080"
      remote = "8080"
      host = "8080"
    }
}

