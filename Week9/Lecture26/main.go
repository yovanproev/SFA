package main

import (
	"database/sql"
	"hns/hns"
	"net/http"
	"time"
)

func GetDataFromServer() hns.TopStories {
	storiesURL := "https://hacker-news.firebaseio.com/v0"
	worker := hns.NewWorker(storiesURL)
	topTenStoriesId := worker.FetchTopStories()
	gotFirstTen := worker.GeneratorStoriesToStruct(topTenStoriesId)

	gotFirstTen.PageTitle = "Top 10 Hacker News Stories"
	hns.IndexTemplate(gotFirstTen)

	return gotFirstTen
}

func writeToDBAndPush(db *sql.DB) {
	mux := http.NewServeMux()
	gotFirstTen := GetDataFromServer()
	hns.StoreItem(db, gotFirstTen.Story)

	gotFirstTen.PageTitle = "Top 10 Hacker News Stories"
	hns.IndexTemplate(gotFirstTen)

	mux.HandleFunc("/api/top", gotFirstTen.HandleUserJSONResponse)
	http.ListenAndServe(":9000", mux)
}

func main() {
	mux := http.NewServeMux()

	db, resultFromDb := hns.InitializeDB("data.db")

	if resultFromDb.Story == nil {
		writeToDBAndPush(db)

	} else {
		lastStoredItems := resultFromDb.Story[0].DateStamp
		now := time.Now()
		hourlyCheck := lastStoredItems.Add(1 * time.Hour)

		if now.After(hourlyCheck) {
			writeToDBAndPush(db)

		} else {
			mux.HandleFunc("/api/top", resultFromDb.HandleUserJSONResponse)
			http.ListenAndServe(":9000", mux)
		}
	}

}
