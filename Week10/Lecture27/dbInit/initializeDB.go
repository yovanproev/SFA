package hns

import (
	"context"
	"database/sql"
	"hns/db"
	"hns/hns"
	"log"
)

func InitializeDB(s string) (*sql.DB, hns.SummedType) {
	sqLiteDB := InitDB(s)
	CreateTable(sqLiteDB)

	q := db.New(sqLiteDB)

	var storiesFromDb hns.SummedType
	listStories, err := q.ListStories(context.Background())
	if err != nil {
		log.Println(err)
	}

	storiesFromDb.Items = append(storiesFromDb.Items, listStories...)

	return sqLiteDB, storiesFromDb
}
