include .env
# Define database URL from environment variables
DB_PATH ?= $(DB_PATH)
DB_URL ?= $(DB_URL)

# Create new migration file and remove file .down.sql
create:
	@echo "Creating new migration without .down.sql..."
	@if [ -z "$(NAME)" ]; then \
	    echo "Error: NAME is not specified. Use 'make create NAME=your_migration_name'"; \
	    exit 1; \
	fi; \
	migrate create -ext sql -dir migrations $(NAME); \
	rm -f migrations/*$(NAME).down.sql


# Migration Up
up:
	@echo "Running migration up..."
	migrate -path $(DB_PATH) -database "$(DB_URL)" -verbose up

# Migration Down
down:
	@echo "Running migration down..."
	migrate -path $(DB_PATH) -database "$(DB_URL)" -verbose down

# Rollback Last Migration
last:
	@echo "Rolling back last migration..."
	migrate -path $(DB_PATH) -database "$(DB_URL)" -verbose down 1

# Fix Migration Version
fix:
	@echo "Fixing migration to version $(VERSION)..."
	@if [ -z "$(VERSION)" ]; then \
	    echo "Error: VERSION is not specified. Use 'make fix VERSION=your_version'"; \
	    exit 1; \
	fi
	migrate -path $(DB_PATH) -database "$(DB_URL)" force $(VERSION)

# Default target that runs all migrations
all: up

# Clean target to remove any generated files or perform cleanup
clean:
	@echo "Cleaning up..."
	# Add cleanup commands here

lint:
	go vet ./...
	staticcheck -checks=all ./...