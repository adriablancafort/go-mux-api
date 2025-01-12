package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
    connectionString := "host=188.245.187.207 user=postgres password=40acf5dee87f1cc2 dbname=postgres sslmode=disable"
    
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