package hns

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {

	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS items(
		id   INTEGER  NOT NULL PRIMARY KEY,
		Title TEXT,
		Score INTEGER,
		DateStamp DATETIME 
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}
