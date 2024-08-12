package db

import (
	"database/sql"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var TURSO_DATABASE_URL = os.Getenv("TURSO_DATABASE_URL")
var TURSO_AUTH_TOKEN = os.Getenv("TURSO_AUTH_TOKEN")

var url = TURSO_DATABASE_URL + "?authToken=" + TURSO_AUTH_TOKEN

func GetDB() *sql.DB {
	db, err := sql.Open("libsql", url)
	if err != nil {
		panic(err)
	}
	return db
}
