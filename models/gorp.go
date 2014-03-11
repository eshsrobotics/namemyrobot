package models

import (
	"database/sql"
	"os"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Dbm *gorp.DbMap

func init() {
	Dbm = newDbMap()
}

func newDbMap() *gorp.DbMap {
	dialect, driver := dialectAndDriver()
	return &gorp.DbMap{Db: connect(driver), Dialect: dialect}
}

func dialectAndDriver() (gorp.Dialect, string) {
	switch os.Getenv("ENV") {
	case "PRODUCTION":
		return gorp.PostgresDialect{}, "postgres"
	default:
		return gorp.SqliteDialect{}, "sqlite3"
	}
}

func connect(driver string) *sql.DB {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		panic("DB_DSN env variable not set")
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic("Error connecting to db: " + err.Error())
	}
	return db
}
