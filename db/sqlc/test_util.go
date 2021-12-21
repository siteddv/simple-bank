package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/siteddv/simple-bank/util"
	"log"
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
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load configs", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	return testDB, err
}
