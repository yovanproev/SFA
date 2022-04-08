package hns

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	server *httptest.Server
)

func TestFetchStory(t *testing.T) {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// switch strings.TrimSpace(r.URL.Path) {
		// case "/":
		// 	mockFetchDataEndpoint(w, r)
		// default:
		// 	http.NotFoundHandler().ServeHTTP(w, r)
		// }
	}))

	handler := GeneratorStoriesToStruct(server.URL).HandleUserJSONResponse

	mockRequest := httptest.NewRequest(http.MethodGet, server.URL, nil)
	responseRecorder := httptest.NewRecorder()
	handler(responseRecorder, mockRequest)

	actualResponse := responseRecorder.Body.String()
	expectedResult := `{"top_stories":[{"title":"","score":0},{"title":"","score":0},{"title":"","score":0},{"title":"","score":0},{"title":"","score":0},{"title":"","score":0},{"title":"","score":0},{"title":"","score":0},{"title":"","score":0},{"title":"","score":0}],"PageTitle":"Top 10 Hacker News Stories","StoryId":null}`

	if !reflect.DeepEqual(actualResponse, expectedResult) {
		t.Errorf("Expected %s, got %s", expectedResult, actualResponse)
	}

}
