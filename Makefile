compose:
	docker-compose --file docker-compose.yml --env-file docker-compose.env up -d

migrate:
	go run ./database/migrate.go