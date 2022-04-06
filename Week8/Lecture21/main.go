package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

type TopStories struct {
	Story []Story `json:"top_stories"`
}

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/top", handleUserJSONResponse)
	http.ListenAndServe(":9000", mux)
}

func handleUserJSONResponse(w http.ResponseWriter, r *http.Request) {
	generatorStoriesToStruct := generatorStoriesToStruct()

	b, err := json.MarshalIndent(generatorStoriesToStruct, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(string(b)))
}

func generatorStoriesToStruct() TopStories {
	ch := make(chan Story)
	var topStories TopStories
	var wg sync.WaitGroup

	firstTenTopStories := fetchTopStories()[0:10]

	go func() {
		for _, story := range firstTenTopStories {
			wg.Add(1)
			go fetchStory(ch, wg, story)
		}
		wg.Wait()
		close(ch)
	}()

	for i := 0; i < 10; i++ {
		topStories.Story = append(topStories.Story, <-ch)
	}

	return topStories
}

func fetchStory(ch chan Story, wg sync.WaitGroup, storyId int) {
	defer wg.Done()

	var story Story

	storyURL := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(storyId) + ".json?print=pretty"

	resp, err := http.Get(storyURL)
	if err != nil {
		fmt.Println("No response from request ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &story); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	ch <- story
}

func fetchTopStories() []int {
	var result []int
	topStoriesURL := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"

	resp, err := http.Get(topStoriesURL)
	if err != nil {
		fmt.Println("No response from request ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

// Output:
// {
// 	"top_stories": [
// 			{
// 					"title": "Debian still having trouble with merged /usr",
// 					"score": 65
// 			},
// 			{
// 					"title": "Design of This Website (2021)",
// 					"score": 139
// 			},
// 			{
// 					"title": "GenieFramework – Web Development with Julia",
// 					"score": 14
// 			},
// 			{
// 					"title": "VisiCalc Executable for the IBM PC (1999)",
// 					"score": 78
// 			},
// 			{
// 					"title": "How Disney Became a Nuclear Power",
// 					"score": 22
// 			},
// 			{
// 					"title": "Ferrari owner Exor wants to build the Italian Y Combinator",
// 					"score": 32
// 			},
// 			{
// 					"title": "Show HN: Warp, a Rust-based terminal",
// 					"score": 819
// 			},
// 			{
// 					"title": "Steam: Half-Life 2 Hardware Survey (2004)",
// 					"score": 67
// 			},
// 			{
// 					"title": "Ask HN: I'm interested in so many disciplines, but what can I do with that?",
// 					"score": 220
// 			},
// 			{
// 					"title": "How to organize yourself as a solo founder",
// 					"score": 76
// 			}
// 	]
// }
