#!make
include .env
export $(shell sed 's/=.*//' .env)

compose:
	docker-compose -f ./docker/docker-compose.yaml up -d

run:
	go run cmd/main.go
	
test:
	go test ./...
