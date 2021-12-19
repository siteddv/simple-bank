package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5436/simple_bank_db?sslmode=disable"
)

var testQueries = getTestQueries()
var testDB, _ = getTestDB()

func getTestQueries() *Queries {
	testDB, err := getTestDB()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries := New(testDB)

	return testQueries
}

func getTestDB() (*sql.DB, error) {
	testDB, err := sql.Open(dbDriver, dbSource)
	return testDB, err
}
