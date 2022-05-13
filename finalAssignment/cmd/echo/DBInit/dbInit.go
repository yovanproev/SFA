package final

import (
	"database/sql"
	"final/cmd/echo/sqlc/db"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {

	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Println("db not open", err)
	}
	if db == nil {
		log.Println("db is not created")
	}
	return db
}

func OpenDBConnection(s string) *db.Queries {
	sqLiteDB := InitDB(s)
	q := db.New(sqLiteDB)

	// create tasks table if not exists
	CreateUsersTable(sqLiteDB)
	// create list table if not eists
	CreateListsTable(sqLiteDB)
	// create tasks table if not exists
	CreateTasksTable(sqLiteDB)

	return q
}

func CreateListsTable(db *sql.DB) {
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
		log.Println(err)
	}
}

func CreateTasksTable(db *sql.DB) {
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
		log.Println(err)
	}
}

func CreateUsersTable(db *sql.DB) {
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
		log.Println(err)
	}
}
