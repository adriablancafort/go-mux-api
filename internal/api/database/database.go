package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

    connectionString := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode

	// initialize connection pool
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }

	// establish connection
	err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
	
	log.Println("Database connection successfully established")
}

func Close() {
	// close connection
    err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Database connection successfully closed")
}