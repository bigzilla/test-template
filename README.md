# Test Template

## Prerequisites

- [Go](https://go.dev)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Usage

Export necessary variable

```bash
export PORT=:8080
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
export POSTGRES_DB=testdb
```

Check docker compose config

```bash
docker-compose config
```

Run docker compose to setup DB

```bash
docker-compose up -d
```

Run the server

```bash
go run main.go
```

Run unit tests

```bash
go test ./... -race -coverprofile cover.out
go tool cover -func cover.out
```

Coverage

```bash
go tool cover -func cover.out | grep total | awk '{print $3}'
```
