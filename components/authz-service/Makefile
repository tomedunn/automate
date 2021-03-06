include ../../Makefile.common_go

PACKAGE_PATH = github.com/chef/automate/components/authz-service

BINS = ${PACKAGE_PATH}/cmd/authz-service

VERSION ?= $(shell git rev-pasrse HEAD)
BUILD_TIME ?= $(shell date -u '+%Y%m%d%H%M%S')
GO_LDFLAGS = --ldflags "-X ${LIBRARY_PATH}/version.Version=${BUILD_TIME} -X ${LIBRARY_PATH}/version.GitSHA=${GIT_SHA}"

packages:=${PACKAGE_PATH}/...
ifdef CI
    verbose?="-v"
endif

all: lint build test

static: lint
unit: build test

${BINS}: bin
	@echo "GO $@"
	@cd bin; go build --race ${GO_LDFLAGS} $@

bin:
	mkdir -p bin

build: ${BINS}

start: build
	./bin/authz-service examples/config.yml

test:
	@go test -i $(packages)
	@go test $(verbose) -cover -count=1 -parallel=1 -p 1 $(packages)

test_with_db:
	@docker ps | grep authz-postgres || (echo "Docker postgres not up. Run make setup_docker_pg and try this command again."; exit 1)
	go test -i $(packages)
	@PG_URL="postgresql://postgres@127.0.0.1:5432/authz_test?sslmode=disable" go test -cover -count=1 -parallel=1 -p 1 $(packages)
	@echo "Docker containers still up, run 'make kill_docker_pg' to bring them down or test again with make test_with_db."

setup_docker_pg:
	docker run --name authz-postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=authz_test -p 5432:5432 -d postgres:9

kill_docker_pg:
	docker rm -f authz-postgres || true

# Regenerates OPA data from rego files
HAVE_GO_BINDATA := $(shell command -v go-bindata 2> /dev/null)
generate:
ifndef HAVE_GO_BINDATA
	@echo "requires 'go-bindata' (go get -u github.com/twitter/go-bindata/...)"
	@exit 1 # fail
else
	go generate ./...
endif

# Regenerate all *.pb.go files
proto:
	cd ../../ && hab studio run 'source .studiorc; compile_go_protobuf_component authz-service'

.PHONY: all static unit build compile test lint
