package hns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

type TopStories struct {
	Story     []Story `json:"top_stories"`
	PageTitle string
	StoryId   []int
}

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}

func (ts TopStories) HandleUserJSONResponse(w http.ResponseWriter, r *http.Request) {

	b, err := json.Marshal(ts)
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(string(b)))
}

func GeneratorStoriesToStruct(s string) TopStories {
	ch := make(chan Story)
	var topStories TopStories
	var wg sync.WaitGroup

	topStories = TopStories{
		PageTitle: "Top 10 Hacker News Stories",
	}

	allStories := FetchTopStories(topStories)

	go func(URL string) {
		for _, storyId := range allStories.StoryId[:10] {
			wg.Add(1)
			go fetchStory(ch, wg, storyId, URL)
		}
		wg.Wait()
		close(ch)
	}(s)

	for i := 0; i < 10; i++ {
		topStories.Story = append(topStories.Story, <-ch)
	}

	return topStories
}

func fetchStory(ch chan Story, wg sync.WaitGroup, storyURL int, URL string) {
	defer wg.Done()

	var story Story
	var resp *http.Response
	var err error

	if URL == "https://hacker-news.firebaseio.com/v0/item/" {
		resp, err = http.Get(URL + strconv.Itoa(storyURL) + ".json?print=pretty")
	} else {
		resp, err = http.Get(URL)
	}

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

func FetchTopStories(ts TopStories) TopStories {
	topStoriesURL := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"

	resp, err := http.Get(topStoriesURL)
	if err != nil {
		fmt.Println("No response from request ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &ts.StoryId); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return ts
}
