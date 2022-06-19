.PHONY: test

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
