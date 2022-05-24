build:
	go build -o ./.bin/server cmd/server/main.go

run:
	go run cmd/server/main.go

init:
	docker-compose up -d

.PHONY: vendor
vendor:
	go mod vendor
	go mod tidy

test:
	go clean -testcache
	go test -v ./...