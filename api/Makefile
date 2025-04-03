build:
	@go build -o bin/solana-faucet-api cmd/main.go

test:
	@go test -v ./...

run: build
	@GIN_MODE=release ./bin/solana-faucet-api

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run db/migrate.go up || (echo "Migration failed. Check if the database is dirty."; exit 1)

migrate-force:
	@go run db/migrate.go force $V || (echo "Failed to force migration. Ensure the version is correct."; exit 1)

migrate-down:
	@go run db/migrate.go down

migrate-status:
	@go run db/migrate.go version
