#!/bin/sh

migrate -path /app/migrations -database ${POSTGRESQL_MIGRATE} -verbose up