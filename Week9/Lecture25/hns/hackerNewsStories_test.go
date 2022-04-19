package hns

import (
	"encoding/json"
	"fmt"
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

func TestHandleUserJSONResponse(t *testing.T) {
	var ts = TopStories{
		Story: []Story{
			{Title: "First Title", Score: 100},
			{Title: "Second Title", Score: 200},
			{Title: "... Title", Score: 300}},
	}

	router := http.NewServeMux()
	mockServer := httptest.NewServer(router)

	r := httptest.NewRequest("", mockServer.URL, nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(ts.HandleUserJSONResponse)

	handler.ServeHTTP(w, r)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	got := w.Body.String()
	want := turnToProperJSONFormat(ts)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %+v, got %+v", want, got)
	}
}

func turnToProperJSONFormat(ts TopStories) string {
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

	var bytes []byte
	var err error
	if bytes, err = json.MarshalIndent(topStoriesMap, "", "   "); err != nil {
		fmt.Printf("Should be able to marshal the results : %v", err)
	}

	return string(bytes)
}

// $ go test . -v -cover
// === RUN   TestInitDB
// --- PASS: TestInitDB (0.00s)
// === RUN   TestCreateAndStoreAndReadTable
// --- PASS: TestCreateAndStoreAndReadTable (0.01s)
// === RUN   TestGeneratorStoriesToStruct
// --- PASS: TestGeneratorStoriesToStruct (0.08s)
// === RUN   TestFetchTopStories
// --- PASS: TestFetchTopStories (0.02s)
// === RUN   TestHandleUserJSONResponse
// --- PASS: TestHandleUserJSONResponse (0.00s)
// === RUN   TestInitializeDB
// --- PASS: TestInitializeDB (0.01s)
// PASS
// coverage: 72.8% of statements
// ok      hns/hns 0.455s  coverage: 72.8% of statements
