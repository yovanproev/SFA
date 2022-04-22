package main

import (
	"context"
	"database/sql"
	"fmt"
	"hns/db"
	database "hns/dbInit"
	"hns/hns"
	templates "hns/templates"
	"log"
	"net/http"
	"time"
)

func GetDataFromServer() hns.SummedType {
	storiesURL := "https://hacker-news.firebaseio.com/v0"
	worker := hns.NewWorker(storiesURL)
	topTenStoriesId := hns.FetchTopStories(storiesURL)
	gotFirstTen := worker.GeneratorStoriesToStruct(topTenStoriesId)

	return gotFirstTen
}

func writeToDBAndPush(ctx context.Context, DB *sql.DB) {
	mux := http.NewServeMux()
	gotFirstTen := GetDataFromServer()

	q := db.New(DB)

	// create N number of stories
	for i := 0; i < len(gotFirstTen.TopStories); i++ {
		_, err := q.CreateStory(ctx, gotFirstTen.TopStories[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	_, resultFromDb := database.InitializeDB("../../data.db")
	resultFromDb.PageTitle = "Top 10 Hacker News Stories"
	templates.IndexTemplate(resultFromDb)

	mux.Handle("/api/top", hns.HandleUserJSONResponse(resultFromDb))
	http.ListenAndServe(":9000", mux)
}

func main() {
	mux := http.NewServeMux()

	sqLite, resultFromDb := database.InitializeDB("../../data.db")

	if resultFromDb.Items == nil {
		writeToDBAndPush(context.Background(), sqLite)
	} else {
		q := db.New(sqLite)
		lastStoredItems, err := q.GetLastStory(context.Background())
		if err != nil {
			log.Println(err)
		}
		now := time.Now()
		hourlyCheck := lastStoredItems.DateStamp.Add(1 * time.Hour)

		if now.After(hourlyCheck) {
			writeToDBAndPush(context.Background(), sqLite)
		} else {
			resultFromDb.PageTitle = "Top 10 Hacker News Stories"
			templates.IndexTemplate(resultFromDb)

			mux.Handle("/api/top", hns.HandleUserJSONResponse(resultFromDb))
			http.ListenAndServe(":9000", mux)
		}
	}

}
