env "development" {
  url = "postgres://postgres:postgres@localhost:5432/postgres_main?sslmode=disable"
  dev = "postgres://postgres:postgres@localhost:5433/postgres_dev?sslmode=disable"

  migration {
    dir = "file://./internal/database/migrations"
    format = "atlas"
  }

  schema {
    src = ["internal/database/schema.sql"]
  }
}
