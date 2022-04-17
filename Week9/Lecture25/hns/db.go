package hns

import (
	"database/sql"
	"fmt"

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

func StoreItem(db *sql.DB, items []Story) {
	sql_additem := `
	INSERT OR REPLACE INTO items(
		Title,
		Score,
		DateStamp
	) values(?, ?, ?)
	`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(item.Title, item.Score, item.DateStamp)
		if err2 != nil {
			panic(err2)
		}
	}
}

func (ts TopStories) ReadItem(db *sql.DB) TopStories {
	sql_readall := `
	SELECT Title, Score, DateStamp FROM items
	ORDER BY datetime(DateStamp) DESC
	LIMIT 10
	`
	rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result TopStories
	for rows.Next() {
		item := Story{}
		err2 := rows.Scan(&item.Title, &item.Score, &item.DateStamp)
		if err2 != nil {
			panic(err2)
		}
		result.Story = append(result.Story, item)
	}

	return result
}

func checkErr(err error, args ...string) {
	if err != nil {
		fmt.Println("Error")
		fmt.Printf("%q: %s", err, args)
	}
}
