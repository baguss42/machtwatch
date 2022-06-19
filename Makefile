.PHONY: test

PROJECT_NAME = machtwatch
CI_REGISTRY ?= baguss42
IMAGE = $(CI_REGISTRY)/$(PROJECT_NAME)
ODIR := deploy/_output

export VERSION ?= $(shell git show -q --format=%h)
export LATEST_VERSION ?= latest

run:
	go run ./app/service/main.go

test:
	go test -v -coverprofile cover.out ./...

cover: test
	go tool cover -html cover.out -o cover.html

pretty:
	go fmt ./...

compose:
	docker-compose --file docker-compose.yml --env-file docker-compose.env up -d

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/service/service app/service/main.go

build:
	docker build -t $(IMAGE):$(VERSION) -f Dockerfile .

deploy: compile build compose