package hns

import (
	"context"
	"database/sql"
	"hns/db"
	"log"
)

func DeleteStory(DB *sql.DB, id int32) {
	q := db.New(DB)

	err := q.DeleteStory(context.Background(), id)
	if err != nil {
		log.Println("delete Story", err)
	}
}
