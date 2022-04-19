package hns

import (
	"database/sql"
)

func InitializeDB(s string) (*sql.DB, TopStories) {
	ts := TopStories{}

	db := InitDB(s)
	CreateTable(db)
	resultFromDb := TopStories.ReadItem(ts, db)

	return db, resultFromDb
}
