#!/usr/bin/env bash
  echo "exporting vars"
	export DB_HOST=localhost
	export DB_NAME=todotest_db
	export DB_PASSWORD=postgres
	export DB_USER=postgres
	export DB_PORT=5432
	export PORT=8000
	export POSTGRESQL_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&search_path=public"

	migrate -database ${POSTGRESQL_URL} -path models/migrations up
	go test -v ./tests/... -race -covermode=atomic -coverprofile=coverage.out
	migrate -database ${POSTGRESQL_URL} -path models/migrations down -all