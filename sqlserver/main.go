package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// Replace with your own connection parameters
var server = "12.12.12.88"
var port = 1433
var user = "sa"
var password = "S0lm1tr@123"
var database = "SolmitraDB_STAGING"

var db *sql.DB

func ReadEmployees(db *sql.DB) (int, error) {
	tsql := fmt.Sprintf("SELECT TOP(1) ID, Code, IdentityNumber FROM Member;")
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var name, location string
		var id string
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, err
		}
		fmt.Printf("ID: %s, Code: %s, IdentityNumber: %s\n", id, name, location)
		count++
	}
	return count, nil
}

func main() {
	var err error

	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	// SelectVersion()

	count, err := ReadEmployees(db)
	if err != nil {
		log.Fatal("ReadEmployees failed:", err.Error())
	}
	fmt.Printf("Read %d rows successfully.\n", count)

	// Close the database connection pool after program executes
	defer db.Close()
}

// Gets and prints SQL Server version
func SelectVersion() {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	// Run query and scan for result
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
}
