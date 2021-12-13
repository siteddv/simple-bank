package db

import (
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5436/simple_bank_db?sslmode=disable"
)
