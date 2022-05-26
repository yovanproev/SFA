package echohandlers

import (
	"context"
	"database/sql"
	"encoding/json"
	CSV "final/pkg/app/CSV"
	handleErrors "final/pkg/app/errors"
	"final/pkg/app/users"
	weather "final/pkg/app/weather"
	"final/pkg/config"
	"final/pkg/sqlc/db"
	initDB "final/pkg/sqlc/initDB"

	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var (
	mockTaskCreation = db.CreateTaskParams{
		Text:      "Test Task",
		Listid:    1,
		Completed: true,
		Userid:    1,
	}
	mockPatchTask = []db.Task{{
		ID:        1,
		Text:      "Test Task",
		Listid:    1,
		Userid:    1,
		Completed: 1,
	}}
	mockListsCreation = db.List{
		Name:   "Test List",
		Userid: 1,
	}
	mockLists       = []db.List{{ID: 1, Name: "Test List", Userid: 1}}
	mockTaskRecords = [][]string{{"Tasks"}, {"Test Task"}}
)

func mockDB() *db.Queries {
	configuration := config.Configurations{}
	configuration.SetConfig()
	handleError := handleErrors.Error{}.SetErrors()

	sqLite, err := sql.Open("sqlite3", configuration.Database.DevelopmentDBName)
	if err != nil {
		fmt.Println("could not connect to database: %w", err)
	}

	initDB.CreateUsersTable(sqLite, handleError)
	initDB.CreateListsTable(sqLite, handleError)
	initDB.CreateTasksTable(sqLite, handleError)

	q := db.New(sqLite)

	return q
}

func TestGetUserByUsername(t *testing.T) {
	users.CreateUser(mockDB(), "test", "test")

	want := db.User{
		ID:       1,
		Username: "test",
		Password: "test",
	}

	got, err := mockDB().GetUserByUsername(context.Background(), want.Username)
	if err != nil {
		fmt.Println(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %+v, got %+v", want, got)
	}
}

func TestPostLists(t *testing.T) {
	handleError := handleErrors.Error{}.SetErrors()
	// Setup
	e := echo.New()

	handler := PostList(mockDB(), handleError)

	e.POST("/api/lists", handler)

	r := httptest.NewRequest(echo.POST, "/api/lists", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockListsCreation)
	if err != nil {
		fmt.Println(err)
	}
	want := string(j)
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, want, got)
}

func TestPostTasks(t *testing.T) {
	handleError := handleErrors.Error{}.SetErrors()
	// Setup
	e := echo.New()

	handler := PostTasks(mockDB(), handleError)
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest(echo.POST, "/api/lists/:id/tasks", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockTaskCreation)
	if err != nil {
		fmt.Println(err)
	}
	want := string(j)
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, want, got)
}

func TestGetTasks(t *testing.T) {
	handleError := handleErrors.Error{}.SetErrors()
	// Setup
	e := echo.New()
	handler := GetTasks(mockDB(), handleError)

	e.GET("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest(echo.GET, "/api/lists/:id/tasks", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	var slice = []db.CreateTaskParams{mockTaskCreation}
	j, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}
	want := string(j)
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, got)
}

func TestGetLists(t *testing.T) {
	handleError := handleErrors.Error{}.SetErrors()
	// Setup
	e := echo.New()

	handler := GetLists(mockDB(), handleError)

	e.GET("/api/lists", handler)

	r := httptest.NewRequest(echo.GET, "/api/lists", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockLists)
	if err != nil {
		fmt.Println(err)
	}
	want := string(j)
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, got)
}

func TestProduceCSV(t *testing.T) {
	handleError := handleErrors.Error{}.SetErrors()
	e := echo.New()

	handler := ProduceCSV(mockDB(), "", handleError)
	e.GET("/api/list/export", handler)

	r := httptest.NewRequest(echo.GET, "/api/list/export", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	toString := CSV.CreateBytesFromTasks(mockTaskRecords)
	want := toString

	got := w.Body.String()

	assert.Equal(t, want, got)
	assert.Equal(t, 200, w.Code)
}

func TestPatchTasks(t *testing.T) {
	handleError := handleErrors.Error{}.SetErrors()
	// Setup
	e := echo.New()

	handler := PatchTasks(mockDB(), handleError)
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest(echo.POST, "/api/lists/:id/tasks", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockPatchTask)
	if err != nil {
		fmt.Println(err)
	}
	want := string(j)
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, want, got)
}

func TestDeleteTasks(t *testing.T) {
	handleError := handleErrors.Error{}.SetErrors()
	// Setup
	e := echo.New()

	handler := DeleteTask(mockDB(), handleError)
	e.DELETE("/api/tasks/:id", handler)

	r := httptest.NewRequest(echo.DELETE, "/api/tasks/:id", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal("Task 0 is deleted!")
	if err != nil {
		fmt.Println(err)
	}
	want := string(j)
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, got)
}

func TestDeleteLists(t *testing.T) {
	handleError := handleErrors.Error{}.SetErrors()
	// Setup
	e := echo.New()

	handler := DeleteList(mockDB(), handleError)
	e.DELETE("/api/lists/:id", handler)

	r := httptest.NewRequest(echo.DELETE, "/api/lists/:id", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)
	j, err := json.Marshal(mockLists)
	if err != nil {
		fmt.Println(err)
	}
	want := string(j)
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, got)
}

var (
	mockWeatherInfo = weather.WeatherInfo{
		FormattedTemp: "32.190000",
		Description:   "clear sky",
		City:          "Turabah",
	}
	mockWeather = `{"coord":{"lon":12,"lat":15},"weather":[{"id":804,"main":"Clouds","description":"clear sky","icon":"04n"}],"base":"stations","main":{"temp":32.19,"feels_like":26.8,"temp_min":28.26,"temp_max":28.26,"pressure":1008,"humidity":9,"sea_level":1008,"grnd_level":969},"visibility":10000,"wind":{"speed":4.92,"deg":56,"gust":6.31},"clouds":{"all":93},"dt":1653359063,"sys":{"country":"NE","sunrise":1653367318,"sunset":1653413759},"timezone":3600,"id":2440495,"name":"Turabah","cod":200}`
)

func TestGetWeather(t *testing.T) {
	configuration := config.Configurations{}
	configuration.SetConfig()
	handleError := handleErrors.Error{}.SetErrors()

	e := echo.New()

	apiKeys := config.LoadEnv(configuration.DevelopmentEnv, configuration)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockWeather))
	}))
	defer server.Close()

	handler := GetWeather(apiKeys, server.URL, handleError)
	e.GET("/api/weather", handler)

	r := httptest.NewRequest(echo.GET, "/api/weather", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockWeatherInfo)
	if err != nil {
		fmt.Println(err)
	}
	got := strings.TrimSuffix(w.Body.String(), "\n")

	want := string(j)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, want, got)
}

func TestEmptyTheDBAfterTests(t *testing.T) {
	mockDB().DeleteUserByUsername(context.Background(), "test")
}

// $ go test . -v -cover
// === RUN   TestGetUserByUsername
// --- PASS: TestGetUserByUsername (0.25s)
// === RUN   TestPostLists
// --- PASS: TestPostLists (0.06s)
// === RUN   TestPostTasks
// --- PASS: TestPostTasks (0.13s)
// === RUN   TestGetTasks
// --- PASS: TestGetTasks (0.12s)
// === RUN   TestGetLists
// --- PASS: TestGetLists (0.02s)
// === RUN   TestProduceCSV
// --- PASS: TestProduceCSV (0.59s)
// === RUN   TestPatchTasks
// --- PASS: TestPatchTasks (0.02s)
// === RUN   TestDeleteTasks
// --- PASS: TestDeleteTasks (0.06s)
// === RUN   TestDeleteLists
// --- PASS: TestDeleteLists (0.59s)
// === RUN   TestGetWeather
// --- PASS: TestGetWeather (0.03s)
// === RUN   TestEmptyTheDBAfterTests
// --- PASS: TestEmptyTheDBAfterTests (0.16s)
// PASS
// coverage: 62.3% of statements
// ok      final/pkg/app/echohandlers      0.699s  coverage: 62.3% of statements
