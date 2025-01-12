package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

    connectionString := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode

    var err error
    // initialize connection
    DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatalf("Error initializing database connection: %v", err)
    }

    // establish connection
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Error establishing database connection: %v", err)
    }

    log.Println("Database connection successfully established")
}

func Close() {
    // close connection
    err := DB.Close()
    if err != nil {
        log.Fatalf("Error closing database connection: %v", err)
    }

    log.Println("Database connection successfully closed")
}
