
# Makefile for chain-healthCheck

.PHONY: run docker-up

run:
	go run server/main.go

docker-up:
	docker-compose up --build -d

build:
	go build -o chain-healthcheck server/main.go