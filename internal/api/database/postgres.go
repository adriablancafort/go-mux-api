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
    DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatalf("Error initializing Postgres connection: %v", err)
    }

    // establish connection
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Error establishing Postgres connection: %v", err)
    }

    log.Println("Postgres connection successfully established")
}

func PostgresClose() {
    // close connection
    err := DB.Close()
    if err != nil {
        log.Fatalf("Error closing Postgres connection: %v", err)
    }

    log.Println("Postgres connection successfully closed")
}
