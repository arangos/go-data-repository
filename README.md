```markdown
# mccusa-data-repository-go
Go repository application for MCCUSA

## Getting Started
These instructions will get you a copy of the Go project up and running on your local machine for development and testing purposes.

### Prerequisites
Make sure you have the following installed:

- Go 1.21 or later (https://golang.org/dl/)
- [Migrate CLI](https://github.com/golang-migrate/migrate) (for running Liquibase changelogs)
- Optional: Docker & Docker Compose (for local Postgres testing)

### Installing Go
Follow the official instructions to install Go:

```shell
# On macOS/Linux via Homebrew or package manager
brew install go
# or
sudo apt install golang-go

# On Windows via installer
# Download and run from https://golang.org/dl/
```

### Clone the repository

```shell
git clone https://github.com/CodeBullsCode/mccusa-data-repository-go.git
cd mccusa-data-repository-go
```

## Building the application

### Local Development
Set your active profile and build:

```shell
# Copy .properties for dev environment
cp src/main/resources/application-local.properties .env
# Build the binary
go build -o bin/mccusa-data-repository-go ./cmd/app
```

### Production Build

```shell
# Copy production config
cp src/main/resources/application-prod.properties .env
# Build optimized binary
go build -ldflags="-s -w" -o bin/mccusa-data-repository-go ./cmd/app
```

## Running the application

### Local Environment

```shell
# Load env vars and run
go run ./cmd/app --env=local
```

### Production Environment

```shell
# Ensure .env points to production configs
./bin/mccusa-data-repository-go --env=prod
```

## Configuration
Configuration is managed via Viper reading `.properties` files in `src/main/resources`. See `config/config.go` for variables and environment overrides.

## Database Migrations
This project uses Liquibase changelogs under `src/main/resources/liquibase`.

```shell
# Run migrations before starting the app
migrate -path src/main/resources/liquibase -database "$DATABASE_URL" up
```

## Testing

### Unit Tests

```shell
go test ./...```
```
