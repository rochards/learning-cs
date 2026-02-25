# Data Access (Go + MySQL)

This project is a small Go application that demonstrates database access patterns with MySQL using `database/sql` and `go-sql-driver/mysql`.

It creates and queries an `album` table in a `recordings` database, including examples for:
- querying multiple rows
- querying a single row by ID
- inserting a new row

Source tutorial: https://go.dev/doc/tutorial/database-access

## Requirements

- Go (version from `go.mod`)
- Docker + Docker Compose
- `make`

## Run

From the project root:

```bash
make run
```

What this does:
- starts MySQL with Docker Compose
- waits until MySQL is healthy
- runs the Go app

Default credentials are set in `Makefile`:
- `DBUSER=appuser`
- `DBPASS=apppass`

You can override them:

```bash
DBUSER=myuser DBPASS=mypassword make run
```

## Stop

```bash
make stop
```

This stops containers and removes volumes (`docker compose down -v`).
