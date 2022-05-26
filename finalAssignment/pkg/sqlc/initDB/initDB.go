package db

import (
	"database/sql"
	handleErrors "final/pkg/app/errors"
	"final/pkg/sqlc/db"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string, e handleErrors.Error) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Println(e.DatabaseInit, err)
	}
	if db == nil {
		log.Println(e.DatabaseInit, err)
	}

	return db
}

func OpenDBConnection(filepath string, e handleErrors.Error) *db.Queries {
	sqLiteDB := InitDB(filepath, e)
	q := db.New(sqLiteDB)

	// create tasks table if not exists
	CreateUsersTable(sqLiteDB, e)
	// create list table if not eists
	CreateListsTable(sqLiteDB, e)
	// create tasks table if not exists
	CreateTasksTable(sqLiteDB, e)

	return q
}

func CreateListsTable(db *sql.DB, e handleErrors.Error) {
	// create able if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS lists(
	id   INTEGER  NOT NULL PRIMARY KEY,
		name TEXT,
		userId INTEGER
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		log.Println(e.DatabaseInit, err)
	}
}

func CreateTasksTable(db *sql.DB, e handleErrors.Error) {
	//create table if not exits
	sql_table := `CREATE TABLE IF NOT EXISTS tasks(
		id  INTEGER NOT NULL PRIMARY KEY,
				text TEXT,
	    listId INTEGER,
		userId INTEGER,
		completed BIT 
		);
		`

	_, err := db.Exec(sql_table)
	if err != nil {
		log.Println(e.DatabaseInit, err)
	}
}

func CreateUsersTable(db *sql.DB, e handleErrors.Error) {
	//create table if not exits
	sql_table := `CREATE TABLE IF NOT EXISTS users(
		id  INTEGER NOT NULL PRIMARY KEY,
		username TEXT,
	    password TEXT,
		datestamp TIMESTAMP NOT NULL
		);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		log.Println(e.DatabaseInit, err)
	}
}
