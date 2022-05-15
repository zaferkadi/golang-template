DB_SERVER_NAME := postgres
POSTGRES_USER := postgres
POSTGRES_PASSWORD := secret
POSTGRES_DB := simple_books
POSTGRESQL_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DB_SERVER_NAME}:5432/${POSTGRES_DB}?sslmode=disable
POSTGRESQL_MIGRATE_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable

export POSTGRESQL_URL
export POSTGRES_USER
export POSTGRES_PASSWORD
export POSTGRES_DB
export DB_SERVER_NAME

up:
	docker-compose up

down:
	docker-compose down

remove:
	docker rmi template-go-server_api

destroy:
	docker-compose down --remove-orphans

createdb:
	docker exec -it ${DB_SERVER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker exec -it ${DB_SERVER_NAME} dropdb ${POSTGRES_DB}

migrateup:
	migrate -path db/migration -database ${POSTGRESQL_MIGRATE_URL} -verbose up

migrateup1:
	migrate -path db/migration -database ${POSTGRESQL_MIGRATE_URL} -verbose up 1

migratedown:
	migrate -path db/migration -database ${POSTGRESQL_MIGRATE_URL} -verbose down

migratedown1:
	migrate -path db/migration -database ${POSTGRESQL_MIGRATE_URL} -verbose down 1

generate-schema:
	docker exec -it ${DB_SERVER_NAME} pg_dump simple_books -U postgres > schema/schema.sql

sqlc:
	sqlc generate

server:
	go run main.go 

test:
	go test -v -cover ./...

generate-mock:
	go generate -v ./...

.PHONY: generate-mock test server postgres mysql createdb dropdb sqlc migrateup migrateup1 migratedown migratedown1