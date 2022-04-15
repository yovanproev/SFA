package hns

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
)

func TestGeneratorStoriesToStruct(t *testing.T) {
	want := struct {
		Story []Story
	}{
		Story: []Story{
			{Title: "DOS Nostalgia: On using a modern DOS workstation", Score: 71},
			{Title: "“iViewed your API Keys”: Aussie state media publishes env vars", Score: 29},
			{Title: "An Argument for a Return to Web 1.0", Score: 19},
			{Title: "Show HN: AV1 and WebRTC", Score: 71},
			{Title: "Licence to Crenellate", Score: 5},
			{Title: "New Covid nasal spray outperforms current antibody treatments in mice", Score: 6},
			{Title: "Clarifying the structure and nature of left-wing authoritarianism", Score: 32},
			{Title: "Heroic Newsboy Funerals", Score: 37},
			{Title: "Building a Soundproof, Dustproof Server Rack", Score: 6},
			{Title: "Command Line Programs for the Blind", Score: 63},
		},
	}

	ch := make(chan Story)

	router := http.NewServeMux()
	router.Handle("/api/top", mockFetchStory(ch, want.Story))
	mockServer := httptest.NewServer(router)
	worker := NewWorker(mockServer.URL)

	topTenStories := []int{31023695, 31020229, 31024127, 31019778, 30992719, 31021652, 31014847, 31017098, 31005586, 31023909}
	got := worker.GeneratorStoriesToStruct(topTenStories)

	eqCtr := 0
	for _, got := range got.Story {
		for _, want := range want.Story {
			if reflect.DeepEqual(got, want) {
				eqCtr++
			}
		}
	}

	if eqCtr != len(want.Story) || len(want.Story) != len(got.Story) {
		t.Errorf(`Got length of slice %d, expected %d`, len(got.Story), len(want.Story))
		t.Fatalf(`
			Got %+v,
			expected %+v`, got.Story, want.Story)
	}
}

func TestFetchTopStories(t *testing.T) {
	want := []int{31023695, 31020229, 31024127, 31019778, 30992719, 31021652, 31014847, 31017098, 31005586, 31023909}

	router := http.NewServeMux()
	router.Handle("/", mockFetchTopStories(want))
	mockServer := httptest.NewServer(router)
	worker := NewWorker(mockServer.URL)

	got := worker.FetchTopStories()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`
			Got %+v, 
			expected %+v`, got, want)
	}

}

func mockFetchStory(dataStream chan Story, items []Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var wg sync.WaitGroup
		go func() {
			for i := 0; i < len(items); i++ {
				wg.Add(1)
				go func(idx int) {
					defer wg.Done()
					dataStream <- items[idx]
				}(i)
			}
			wg.Wait()
		}()

		json.NewEncoder(w).Encode(<-dataStream)
	}
}

func mockFetchTopStories(items []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(items)
	}
}

// $ go test . -v -cover
// === RUN   TestGeneratorStoriesToStruct
// --- PASS: TestGeneratorStoriesToStruct (0.04s)
// === RUN   TestFetchTopStories
// --- PASS: TestFetchTopStories (0.00s)
// PASS
// coverage: 79.5% of statements
// ok      hns/hns 0.137s  coverage: 79.5% of statements
