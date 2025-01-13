package database

import (
    "log"
    "os"

    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

var DB *gorm.DB

func PostgresInnit() {
    host := os.Getenv("POSTGRES_HOST")
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_NAME")
    sslmode := os.Getenv("POSTGRES_SSLMODE")

    dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error initializing Postgres connection: %v", err)
    }

    log.Println("Postgres connection successfully established")
}
