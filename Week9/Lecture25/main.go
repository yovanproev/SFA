package main

import (
	"database/sql"
	"hns/hns"
	"html/template"
	"log"
	"net/http"
	"time"
)

func getDataFromServer() hns.TopStories {
	storiesURL := "https://hacker-news.firebaseio.com/v0"
	worker := hns.NewWorker(storiesURL)
	topTenStoriesId := worker.FetchTopStories()
	gotFirstTen := worker.GeneratorStoriesToStruct(topTenStoriesId)

	return gotFirstTen
}

func writeToDBAndPush(db *sql.DB) {
	mux := http.NewServeMux()
	gotFirstTen := getDataFromServer()
	hns.StoreItem(db, gotFirstTen.Story)

	mux.HandleFunc("/api/top", gotFirstTen.HandleUserJSONResponse)
	http.ListenAndServe(":9000", mux)
}

func indexTemplate(gotFirstTen hns.TopStories) {
	mux := http.NewServeMux()

	templates := populateTemplates()

	mux.HandleFunc("/top", func(w http.ResponseWriter, r *http.Request) {
		t := templates.Lookup("index.html")
		t.Execute(w, gotFirstTen)

		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	})
	http.ListenAndServe(":9000", mux)
}

func main() {
	mux := http.NewServeMux()

	ts := hns.TopStories{}

	db := hns.InitDB("data.db")
	hns.CreateTable(db)
	resultFromDb := hns.TopStories.ReadItem(ts, db)

	if resultFromDb.Story == nil {
		writeToDBAndPush(db)
		ts := getDataFromServer()
		ts.PageTitle = "Top 10 Hacker News Stories"
		indexTemplate(ts)

	} else {
		lastStoredItems := resultFromDb.Story[0].DateStamp
		now := time.Now()
		hourlyCheck := lastStoredItems.Add(1 * time.Hour)

		ts := resultFromDb
		ts.PageTitle = "Top 10 Hacker News Stories"
		indexTemplate(ts)

		if now.After(hourlyCheck) {
			writeToDBAndPush(db)
		} else {
			mux.HandleFunc("/api/top", resultFromDb.HandleUserJSONResponse)
			http.ListenAndServe(":9000", mux)
		}
	}

}

func populateTemplates() *template.Template {
	result := template.New("")
	const basePath = "Week9/Lecture25"
	template.Must(result.ParseGlob(basePath + "/*.html"))

	return result
}
