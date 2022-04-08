package main

import (
	"hns/hns"
	"html/template"
	"log"
	"net/http"
)

func main() {

	storyURL := "https://hacker-news.firebaseio.com/v0/item/"
	generatedStories := hns.GeneratorStoriesToStruct(storyURL)

	mux := http.NewServeMux()
	// // Lecture 22 - Task 1
	templates := populateTemplates()
	generatorStoriesToStruct := hns.GeneratorStoriesToStruct(storyURL)

	mux.HandleFunc("/top", func(w http.ResponseWriter, r *http.Request) {
		t := templates.Lookup("index.html")
		t.Execute(w, generatorStoriesToStruct)

		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	mux.HandleFunc("/api/top", generatedStories.HandleUserJSONResponse)
	http.ListenAndServe(":9000", mux)
}

func populateTemplates() *template.Template {
	result := template.New("")
	const basePath = "Week8/Lecture22"
	template.Must(result.ParseGlob(basePath + "/*.html"))

	return result
}

// Output Task 1:
// Top 10 Hacker News Stories
// Andrew Ng: Unbiggen AI: 159
// I thought I’d have accomplished a lot more today and also before I was 35 (2020): 220
// Hand-optimizing the TCC code generator: 82
// Ask HN: Share your personal site: 494
// It’s Time to Launch the Wolfram Institute: 157
// A Square Meal – Foods of the ‘20s and ‘30s: 178
// Dunbar’s number and how speaking is 2.8x better than picking fleas: 38
// Thoughts on the Witness (2016): 17
// News for the Future of BeeWare: 62
// IBM Archives: System/360 Announcement (April 7, 1964): 16
