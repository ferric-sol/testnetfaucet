build:
	@go build -o bin/solana-faucet-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/solana-faucet-api

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up || (echo "Migration failed. Check if the database is dirty."; exit 1)

migrate-force:
	@go run cmd/migrate/main.go force $V || (echo "Failed to force migration. Ensure the version is correct."; exit 1)

migrate-down:
	@go run cmd/migrate/main.go down

migrate-status:
	@go run cmd/migrate/main.go version