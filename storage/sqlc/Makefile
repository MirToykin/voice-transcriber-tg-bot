include ../../.env

LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

migrate:
	bin/goose -dir "migrations" sqlite3 $(DB_DSN) up -v

sqlc-gen:
	docker run --rm -v ./:/src -w /src sqlc/sqlc generate