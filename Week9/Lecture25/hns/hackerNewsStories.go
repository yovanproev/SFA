package hns

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type TopStories struct {
	Story     []Story `json:"top_stories,omitempty"`
	PageTitle string  `json:"page_title,omitempty"`
	serverURL string
}

type Story struct {
	Title     string    `json:"title,omitempty"`
	Score     int       `json:"score,omitempty"`
	DateStamp time.Time `json:"date_stamp,omitempty"`
}

func NewWorker(serverUrl string) *TopStories {
	var result TopStories
	result.serverURL = serverUrl

	return &result
}

func (ts TopStories) HandleUserJSONResponse(w http.ResponseWriter, r *http.Request) {
	story := []Story{}

	for _, v := range ts.Story {
		story = append(story, Story{
			Title: v.Title,
			Score: v.Score,
		})
	}

	var topStoriesMap = make(map[string][]map[string]interface{})
	var sliceOfMaps = make([]map[string]interface{}, 0)

	for _, v := range story {
		elem := reflect.ValueOf(&v).Elem()
		relType := elem.Type()

		var myMap = make(map[string]interface{})

		for i := 0; i < relType.NumField(); i++ {
			myMap[relType.Field(i).Name] = elem.Field(i).Interface()
		}
		delete(myMap, "DateStamp")

		sliceOfMaps = append(sliceOfMaps, myMap)
		topStoriesMap["top_stories"] = sliceOfMaps
	}

	b, err := json.MarshalIndent(topStoriesMap, "", "   ")
	if err != nil {
		log.Println(err)
	}

	w.Write([]byte(string(b)))
}

func (ts TopStories) GeneratorStoriesToStruct(topTenStoriesId []int) TopStories {
	ch := make(chan Story)
	var wg sync.WaitGroup

	ts.PageTitle = "Top 10 Hacker News Stories"

	go func(URL string) {
		for _, storyId := range topTenStoriesId {
			wg.Add(1)
			go fetchStory(ch, wg, storyId, URL)
		}
		wg.Wait()
		close(ch)
	}(ts.serverURL)

	for i := 0; i < 10; i++ {
		ts.Story = append(ts.Story, <-ch)
	}

	return ts
}

func fetchStory(ch chan Story, wg sync.WaitGroup, storyId int, URL string) {
	defer wg.Done()

	var story Story
	var resp *http.Response
	var err error
	story.DateStamp = time.Now()

	if URL == "https://hacker-news.firebaseio.com/v0" {
		resp, err = http.Get(URL + "/item/" + strconv.Itoa(storyId) + ".json?print=pretty")
	} else {
		resp, err = http.Get(URL + "/api/top")
	}

	if err != nil {
		log.Println("No response from request ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &story); err != nil {
		log.Printf("%s fetchStory for a channel", err)
	}

	ch <- story
}

func (ts TopStories) FetchTopStories() []int {
	var result []int

	var resp *http.Response
	var err error

	resp, err = http.Get(ts.serverURL + "/topstories.json?print=pretty")

	if err != nil {
		log.Println("No response from request ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result[:10]
}
