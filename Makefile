
run:
	go run cmd/main.go

install:
	go mod tidy
	go mod vendor

env:
	cp .env.example .env