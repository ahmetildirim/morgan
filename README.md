# morgan-go

Morgan is a clean architecture example written in Go. It is a side project to learn Go and clean architecture.

## How to run

### Start Postgres

```bash
docker-compose up -d
```

### Install Goose

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Run migrations

```bash
goose -dir migrations postgres "user=postgres password=postgres dbname=morgan sslmode=disable host=localhost" up
```

## Run the app

```bash
go run cmd/main.go
```
