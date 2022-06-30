DB_SERVER_NAME := postgres
POSTGRES_USER := postgres
POSTGRES_PASSWORD := secret
POSTGRES_DB := simple_books
POSTGRESQL_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DB_SERVER_NAME}:5432/${POSTGRES_DB}?sslmode=disable

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

migrateup:
	docker compose exec api sh -c "./migrate -path /app/migrations -database ${POSTGRESQL_URL} -verbose up"

migrateup1:
	docker compose exec api sh -c "./migrate -path /app/migrations -database ${POSTGRESQL_URL} -verbose up 1"

migratedown:
	docker compose exec api sh -c "./migrate -path /app/migrations -database ${POSTGRESQL_URL} -verbose down"

migratedown1:
	docker compose exec api sh -c "./migrate -path /app/migrations -database ${POSTGRESQL_URL} -verbose down 1"

dump-schema-image:
	docker compose exec schemacrawler \
	 /opt/schemacrawler/bin/schemacrawler.sh  \
	--server=postgresql \
	--host=${DB_SERVER_NAME}  \
	--database=${POSTGRES_DB} --user=${POSTGRES_USER} --password=${POSTGRES_PASSWORD} \
	--info-level=maximum --command=schema \
	--output-format=png \
	--output-file=/schema/schema.png
	
sqlc:
	sqlc generate

server:
	go run cmd/http/main.go 

test:
	go test -v -cover ./...

generate-mock:
	go generate -v ./...

.PHONY: generate-mock test server postgres mysql createdb dropdb sqlc migrateup migrateup1 migratedown migratedown1