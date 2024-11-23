package db

import "database/sql"

type Database struct {
	DB       *sql.DB
	dbDriver string
	dbSource string
}

func NewSQL(dbDriver, dbSource string) *Database {
	return &Database{
		dbDriver: dbDriver,
		dbSource: dbSource,
	}
}
