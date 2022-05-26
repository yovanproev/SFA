package ginhandlers

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

	"github.com/gin-gonic/gin"
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
	e := gin.Default()

	handler := PostList(mockDB(), handleError)

	e.POST("/api/lists", handler)

	r := httptest.NewRequest("POST", "/api/lists", nil)
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
	e := gin.Default()

	handler := PostTasks(mockDB(), handleError)
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest("POST", "/api/lists/:id/tasks", nil)
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
	e := gin.Default()
	handler := GetTasks(mockDB(), handleError)

	e.GET("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest("GET", "/api/lists/:id/tasks", nil)
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
	e := gin.Default()

	handler := GetLists(mockDB(), handleError)

	e.GET("/api/lists", handler)

	r := httptest.NewRequest("GET", "/api/lists", nil)
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
	e := gin.Default()

	handler := ProduceCSV(mockDB(), "", handleError)
	e.GET("/api/list/export", handler)

	r := httptest.NewRequest("GET", "/api/list/export", nil)
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
	e := gin.Default()

	handler := PatchTasks(mockDB(), handleError)
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest("POST", "/api/lists/:id/tasks", nil)
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
	e := gin.Default()

	handler := DeleteTask(mockDB(), handleError)
	e.DELETE("/api/tasks/:id", handler)

	r := httptest.NewRequest("DELETE", "/api/tasks/:id", nil)
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
	e := gin.Default()

	handler := DeleteList(mockDB(), handleError)
	e.DELETE("/api/lists/:id", handler)

	r := httptest.NewRequest("DELETE", "/api/lists/:id", nil)
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

	e := gin.Default()

	apiKeys := config.LoadEnv(configuration.DevelopmentEnv, configuration)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockWeather))
	}))
	defer server.Close()

	handler := GetWeather(apiKeys, server.URL, handleError)
	e.GET("/api/weather", handler)

	r := httptest.NewRequest("GET", "/api/weather", nil)
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

func TestEmptyDBAfterTesting(t *testing.T) {
	mockDB().DeleteUserByUsername(context.Background(), "test")
}

// === RUN   TestGetUserByUsername
// --- PASS: TestGetUserByUsername (0.04s)
// === RUN   TestPostLists
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] POST   /api/lists                --> final/pkg/app/ginhandlers.PostList.func1 (3 handlers)
// [GIN] 2022/05/25 - 18:18:41 | 201 |     21.6611ms |       192.0.2.1 | POST     "/api/lists"
// --- PASS: TestPostLists (0.02s)
// === RUN   TestPostTasks
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] POST   /api/lists/:id/tasks      --> final/pkg/app/ginhandlers.PostTasks.func1 (3 handlers)
// [GIN] 2022/05/25 - 18:18:41 | 201 |     20.4609ms |       192.0.2.1 | POST     "/api/lists/:id/tasks"
// --- PASS: TestPostTasks (0.03s)
// === RUN   TestGetTasks
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] GET    /api/lists/:id/tasks      --> final/pkg/app/ginhandlers.GetTasks.func1 (3 handlers)
// [GIN] 2022/05/25 - 18:18:41 | 200 |      1.6935ms |       192.0.2.1 | GET      "/api/lists/:id/tasks"
// --- PASS: TestGetTasks (0.01s)
// === RUN   TestGetLists
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] GET    /api/lists                --> final/pkg/app/ginhandlers.GetLists.func1 (3 handlers)
// [GIN] 2022/05/25 - 18:18:41 | 200 |            0s |       192.0.2.1 | GET      "/api/lists"
// --- PASS: TestGetLists (0.00s)
// === RUN   TestProduceCSV
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] GET    /api/list/export          --> final/pkg/app/ginhandlers.ProduceCSV.func1 (3 handlers)
// [GIN] 2022/05/25 - 18:18:41 | 200 |         518µs |       192.0.2.1 | GET      "/api/list/export"
// --- PASS: TestProduceCSV (0.00s)
// === RUN   TestPatchTasks
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] POST   /api/lists/:id/tasks      --> final/pkg/app/ginhandlers.PatchTasks.func1 (3 handlers)
// [GIN] 2022/05/25 - 18:18:41 | 200 |       517.8µs |       192.0.2.1 | POST     "/api/lists/:id/tasks"
// --- PASS: TestPatchTasks (0.00s)
// === RUN   TestDeleteTasks
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] DELETE /api/tasks/:id            --> final/pkg/app/ginhandlers.DeleteTask.func1 (3 handlers)
// [GIN] 2022/05/25 - 18:18:41 | 200 |     20.0347ms |       192.0.2.1 | DELETE   "/api/tasks/:id"
// --- PASS: TestDeleteTasks (0.02s)
// === RUN   TestDeleteLists
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] DELETE /api/lists/:id            --> final/pkg/app/ginhandlers.DeleteList.func1 (3 handlers)
// [GIN] 2022/05/25 - 18:18:41 | 200 |     22.1431ms |       192.0.2.1 | DELETE   "/api/lists/:id"
// --- PASS: TestDeleteLists (0.03s)
// === RUN   TestGetWeather
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] GET    /api/weather              --> final/pkg/app/ginhandlers.GetWeather.func1 (3 handlers)
// [GIN] 2022/05/25 - 18:18:41 | 200 |      4.3608ms |       192.0.2.1 | GET      "/api/weather"
// --- PASS: TestGetWeather (0.01s)
// === RUN   TestEmptyDBAfterTesting
// --- PASS: TestEmptyDBAfterTesting (0.02s)
// PASS
// coverage: 64.5% of statements
// ok      final/pkg/app/ginhandlers       0.684s  coverage: 64.5% of statements
