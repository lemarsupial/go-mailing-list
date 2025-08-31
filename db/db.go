package db

import (
	"database/sql"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// var TURSO_DATABASE_URL = os.Getenv("TURSO_DATABASE_URL")
// var TURSO_AUTH_TOKEN = os.Getenv("TURSO_AUTH_TOKEN")

// var url = TURSO_DATABASE_URL + "?authToken=" + TURSO_AUTH_TOKEN

var DbHandle *sql.DB

func init() {
	var err error
	DbHandle, err = sql.Open("sqlite3", "toni.db")
	if err != nil {
		panic(err)
	}
	DbHandle.SetMaxOpenConns(1000) // Maximum number of open connections
	DbHandle.SetMaxIdleConns(50)   // Maximum number of idle connections
	DbHandle.SetConnMaxLifetime(0)
}
