package hns

import (
	"encoding/json"
	"hns/db"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestGeneratorStoriesToStruct(t *testing.T) {
	want := struct {
		TopStories []db.CreateStoryParams
	}{
		TopStories: []db.CreateStoryParams{
			{Title: "Command Line Programs for the Blind", Score: 63, DateStamp: time.Now()},
			{Title: "“iViewed your API Keys”: Aussie state media publishes env vars", Score: 29, DateStamp: time.Now()},
			{Title: "DOS Nostalgia: On using a modern DOS workstation", Score: 71, DateStamp: time.Now()},
			{Title: "Show HN: AV1 and WebRTC", Score: 71, DateStamp: time.Now()},
			{Title: "An Argument for a Return to Web 1.0", Score: 19, DateStamp: time.Now()},
			{Title: "Licence to Crenellate", Score: 5, DateStamp: time.Now()},
			{Title: "New Covid nasal spray outperforms current antibody treatments in mice", Score: 6, DateStamp: time.Now()},
			{Title: "Clarifying the structure and nature of left-wing authoritarianism", Score: 32, DateStamp: time.Now()},
			{Title: "Heroic Newsboy Funerals", Score: 37, DateStamp: time.Now()},
			{Title: "Building a Soundproof, Dustproof Server Rack", Score: 6, DateStamp: time.Now()},
		},
	}

	ch := make(chan db.CreateStoryParams)

	router := http.NewServeMux()
	router.Handle("/api/top", mockFetchStory(ch, want.TopStories))
	mockServer := httptest.NewServer(router)
	worker := NewWorker(mockServer.URL)

	topTenStories := []int{31023695, 31020229, 31024127, 31019778, 30992719, 31021652, 31014847, 31017098, 31005586, 31023909}
	got := worker.GeneratorStoriesToStruct(topTenStories)

	eqCtr := 0
	for _, got := range got.TopStories {
		for _, want := range want.TopStories {
			if reflect.DeepEqual(got.Title, want.Title) {
				eqCtr++
			}
		}
	}

	if eqCtr != len(want.TopStories) || len(want.TopStories) != len(got.TopStories) {
		t.Errorf(`Got length of slice %d, expected %d`, len(got.TopStories), len(want.TopStories))
		t.Fatalf(`
			Got %+v,
			expected %+v`, got.TopStories, want.TopStories)
	}
}

func TestFetchTopStories(t *testing.T) {
	want := []int{31023695, 31020229, 31024127, 31019778, 30992719, 31021652, 31014847, 31017098, 31005586, 31023909}

	router := http.NewServeMux()
	router.Handle("/", mockFetchTopStories(want))
	mockServer := httptest.NewServer(router)

	got := FetchTopStories(mockServer.URL)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`
			Got %+v, 
			expected %+v`, got, want)
	}

}

func mockFetchStory(dataStream chan db.CreateStoryParams, items []db.CreateStoryParams) http.HandlerFunc {
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

func TestHandleUserJSONResponse(t *testing.T) {
	var ts = SummedType{
		Items: []db.Item{
			{Title: "First Title", Score: 100},
			{Title: "Second Title", Score: 200},
			{Title: "... Title", Score: 300}},
	}

	router := http.NewServeMux()
	mockServer := httptest.NewServer(router)

	r := httptest.NewRequest("", mockServer.URL, nil)
	w := httptest.NewRecorder()
	handler := http.Handler(HandleUserJSONResponse(ts))

	handler.ServeHTTP(w, r)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	got := w.Body.String()
	formattedStories := StoryToMap(ts)
	toJSON, err := json.MarshalIndent(formattedStories, "", "   ")
	if err != nil {
		log.Println(err)
	}
	want := string(toJSON)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %+v, got %+v", want, got)
	}
}

// $ go test . -v -cover
// === RUN   TestGeneratorStoriesToStruct
// --- PASS: TestGeneratorStoriesToStruct (0.03s)
// === RUN   TestFetchTopStories
// --- PASS: TestFetchTopStories (0.01s)
// === RUN   TestHandleUserJSONResponse
// --- PASS: TestHandleUserJSONResponse (0.00s)
// PASS
// coverage: 90.3% of statements
// ok      hns/hns 0.161s  coverage: 90.3% of statements
