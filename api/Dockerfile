FROM golang:1.24.1-alpine AS builder

# Install build dependencies
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy && go mod download

# Copy source code
COPY . .

# Build the API server binary
RUN go build -v -o /solana-faucet-api ./cmd/main.go

# Build the migration tool binary
RUN go build -v -o /migrate-tool ./cmd/migrate/main.go

# Final Image
FROM alpine:latest

WORKDIR /app

# Copy the compiled binaries from build
COPY --from=builder /solana-faucet-api /solana-faucet-api
COPY --from=builder /migrate-tool /migrate-tool

# Copy migration files (IMPORTANT!)
COPY --from=builder /app/cmd/migrate/migrations /app/cmd/migrate/migrations

# Run command
CMD ["/solana-faucet-api"]
