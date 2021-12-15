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

func GetTestQueries() *Queries {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries := New(conn)

	return testQueries
}
