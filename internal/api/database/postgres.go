package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func PostgresConnect() {
    host := os.Getenv("POSTGRES_HOST")
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_NAME")
    sslmode := os.Getenv("POSTGRES_SSLMODE")

    connectionString := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode

    var err error
    // initialize connection
    if DB, err = sql.Open("postgres", connectionString); err != nil {
        log.Fatalf("Error initializing Postgres connection: %v", err)
    }

    // establish connection
    if err := DB.Ping(); err != nil {
        log.Fatalf("Error establishing Postgres connection: %v", err)
    }

    log.Println("Postgres connection successfully established")
}

func PostgresClose() {
    // close connection
    if err := DB.Close(); err != nil {
        log.Fatalf("Error closing Postgres connection: %v", err)
    }

    log.Println("Postgres connection successfully closed")
}
