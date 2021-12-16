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
var testDB, _ = GetTestDB()

func getTestQueries() *Queries {
	testDB, err := GetTestDB()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries := New(testDB)

	return testQueries
}

func GetTestDB() (*sql.DB, error) {
	testDB, err := sql.Open(dbDriver, dbSource)
	return testDB, err
}
