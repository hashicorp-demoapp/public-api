VERSION=v0.0.10
REPOSITORY=hashicorpdemoapp/public-api

.PHONY: auth

all: generate run

generate:
	go run scripts/gqlgen.go

run:
	go run main.go

build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/amd64/public-api main.go

build_linux_arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/arm64/public-api main.go

build_darwin_arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./bin/arm64-darwin/public-api main.go

build_docker: build_linux_amd64 build_linux_arm64
	docker build -t ${REPOSITORY}:${VERSION} .

build_docker_dev: build_linux_amd64 build_linux_arm64
	docker build -t ${REPOSITORY}:dev .

buildx_docker: build_linux_amd64 build_linux_arm64
	docker buildx create --use --name multi_arch_build || true
	docker buildx build --platform linux/amd64,linux/arm64 \
		-t ${REPOSITORY}:${VERSION} \
		--push \
		.

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
