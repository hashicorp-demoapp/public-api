VERSION=v0.0.3
REPOSITORY=hashicorpdemoapp/public-api

.PHONY: auth

all: generate run

generate:
	go run scripts/gqlgen.go

run:
	go run main.go

build_linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/public-api main.go

build_docker: build_linux
	docker build -t ${REPOSITORY}:${VERSION} .

build_docker_dev: build_linux
	docker build -t ${REPOSITORY}:dev .

run_functional_tests: build_docker_dev
	# First copy the config to a volume, this is needed for CircleCI
	docker create -v data.volume:/config --name dummy alpine /bin/true
	docker cp ./functional_test/config/config.json dummy:/config                                                                                                                            
	cd functional_test && shipyard test

auth:
	docker run -it --rm \
	--publish 8403:3000 \
	-e AUTHN_URL=localhost:8403 \
	-e APP_DOMAINS=localhost \
	-e DATABASE_URL=sqlite3://:memory:?mode=memory\&cache=shared \
	-e SECRET_KEY_BASE='my-authn-test-secret' \
	-e HTTP_AUTH_USERNAME=hello \
	-e HTTP_AUTH_PASSWORD=world \
	--name authn \
	keratin/authn-server:latest \
	sh -c "./authn migrate && ./authn server"