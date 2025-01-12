package db

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
    connectionString := "host= user= password= dbname= sslmode=disable"
    
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