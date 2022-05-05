package hns

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"google.golang.org/grpc/examples/helloworld/greeter_server/db"
)

func InitializeDB(s string) (*sql.DB, SummedType) {
	sqLiteDB := InitDB(s)
	CreateTable(sqLiteDB)

	q := db.New(sqLiteDB)

	var storiesFromDb SummedType
	listStories, err := q.ListStories(context.Background())
	if err != nil {
		log.Println(err)
	}

	storiesFromDb.Items = append(storiesFromDb.Items, listStories...)

	return sqLiteDB, storiesFromDb
}

func WriteToDBAndPush(ctx context.Context, DB *sql.DB) (SummedType, error) {
	gotFirstTen := GetDataFromServer()

	q := db.New(DB)

	// create N number of stories
	for i := 0; i < len(gotFirstTen.TopStories); i++ {
		_, err := q.CreateStory(ctx, gotFirstTen.TopStories[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	_, resultFromDb := InitializeDB("../data.db")

	return resultFromDb, nil
}

func GetDataFromServer() SummedType {
	storiesURL := "https://hacker-news.firebaseio.com/v0"
	worker := NewWorker(storiesURL)
	topTenStoriesId := FetchTopStories(storiesURL)
	gotFirstTen := worker.GeneratorStoriesToStruct(topTenStoriesId)

	return gotFirstTen
}
