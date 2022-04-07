HAS_GOCILINT  = $(shell command -v golangci-lint)

postgres:
	docker run --name postgres12 --network books-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_books

dropdb:
	docker exec -it postgres12 dropdb simple_books

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_books?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_books?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_books?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_books?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

server:
	go run ./cmd/http/main.go 

test:
	go test -v -cover ./...

generate-mock:
	go generate -v ./...

lint: ## Perform lint checks
ifndef HAS_GOCILINT
	echo "Please install golang-ci-lint `make install-golang-ci-lint`"
endif
	golangci-lint run --timeout=30m

.PHONY: generate-mock test server postgres mysql createdb dropdb sqlc migrateup migrateup1 migratedown migratedown1 lint