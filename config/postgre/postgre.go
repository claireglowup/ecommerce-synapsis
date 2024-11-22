package config

import (
	"database/sql"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Postgre struct {
	DB *sql.DB
	dbDriver string
	dbSource string
}

func Database(dbDriver, dbSource string) *Postgre {
	return &Postgre{
		dbDriver: dbDriver,
		dbSource: dbSource,
	}
}

func (p *Postgre) InitPostgre() {

	pg, err := sql.Open(p.dbDriver, p.dbSource)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	err = pg.Ping()
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)

	}
	pg.SetMaxIdleConns(10)
	pg.SetMaxOpenConns(100)
	pg.SetConnMaxIdleTime(5 * time.Minute)
	pg.SetConnMaxLifetime(60 * time.Minute)

	p.DB = pg
}



func (p *Postgre) RunDBMigration(migrationSource string) {
	migration, err := migrate.New(migrationSource, p.dbSource)
	if err != nil {
		log.Fatalf("cannot create new migration instance: %v", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrate up : %v", err)
	}

	log.Print("db migration successfully")
}