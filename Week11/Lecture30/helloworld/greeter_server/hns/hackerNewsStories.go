package hns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc/examples/helloworld/greeter_server/db"
)

type SummedType struct {
	PageTitle  string
	TopStories []db.CreateStoryParams
	Items      []db.Item
	serverURL  string
}

func NewWorker(serverUrl string) *SummedType {
	var result SummedType
	result.serverURL = serverUrl

	return &result
}

func (s SummedType) GeneratorStoriesToStruct(topTenStoriesId []int) SummedType {
	ch := make(chan db.CreateStoryParams)

	var wg sync.WaitGroup

	go func(URL string) {
		for _, storyId := range topTenStoriesId {
			wg.Add(1)
			go fetchStory(ch, wg, storyId, URL)
		}
		wg.Wait()
		close(ch)
	}(s.serverURL)

	for i := 0; i < 10; i++ {
		s.TopStories = append(s.TopStories, <-ch)
		s.TopStories[i].DateStamp = time.Now()
	}

	return s
}

func fetchStory(ch chan db.CreateStoryParams, wg sync.WaitGroup, storyId int, URL string) {
	defer wg.Done()

	var story db.CreateStoryParams
	var resp *http.Response
	var err error

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
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(body, &story); err != nil {
		log.Printf("%s fetchStory for a channel", err)
	}

	ch <- story
}

func FetchTopStories(serverURL string) []int {
	var result []int

	var resp *http.Response
	var err error

	resp, err = http.Get(serverURL + "/topstories.json?print=pretty")

	if err != nil {
		log.Println("No response from request ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Can not unmarshal JSON %+v", err)
	}

	return result[:10]
}
