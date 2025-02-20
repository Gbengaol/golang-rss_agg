### Getting started

1. Install sqlc (for querying) using cmd `brew install sqlc` on MacOS or cmd `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest` on other environments
2. Confirm that sqlc is installed using cmd `sqlc version`
3. Install goose (for migration) using cmd `brew install goose` on MacOS or cmd `go install github.com/pressly/goose/v3/cmd/goose@latest` on other environments
4. Confirm that goose is installed using cmd `goose -version`
5. cd to sql/schema and run `goose postgres postgres://<username>:<password>@localhost:<port>/<database_name> up`
