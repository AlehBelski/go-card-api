.PHONY: dependency unit-test integration-test docker-up docker-down clear

dependency:
	@go mod download

integration-test:
	make dependency docker-up
	@go test -v -run Integration ./...
	make docker-down

unit-test: dependency
	@go test -v -short ./...

docker-up:
	@docker-compose up -d
	@goose -dir migrations postgres "user=postgres password=postgres dbname=postgres sslmode=disable" up

docker-down:
	@docker-compose down
