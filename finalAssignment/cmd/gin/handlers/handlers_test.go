package final

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	final "final/cmd/gin/DBInit"
	"final/cmd/gin/sqlc/db"
	weather "final/cmd/gin/weather"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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
	mockTask = []db.CreateTaskParams{{
		Text:      "Test Task",
		Listid:    1,
		Userid:    1,
		Completed: true,
	},
	}
	mockListsCreation = db.CreateListParams{
		Name:   "Test List",
		Userid: 1,
	}
	mockLists = []db.List{{ID: 1, Name: "Test List", Userid: 1}}
)

func mockDB() *db.Queries {
	sqLite, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("could not connect to database: %w", err)
	}

	final.CreateUsersTable(sqLite)
	final.CreateListsTable(sqLite)
	final.CreateTasksTable(sqLite)

	q := db.New(sqLite)

	return q
}

func TestGetUserByUsername(t *testing.T) {
	CreateUser(mockDB(), "test", "test")

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
	// Setup
	e := gin.Default()

	_, err2 := mockDB().UpdateUsersById(context.Background(), 1)
	if err2 != nil {
		fmt.Println(err2)
	}

	handler := PostList(mockDB())

	e.POST("/api/lists", handler)

	r := httptest.NewRequest("POST", "/api/lists", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockListsCreation)
	if err != nil {
		fmt.Println(err)
	}
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, string(j), got)
}

func TestPostTasks(t *testing.T) {
	// Setup
	e := gin.Default()

	handler := PostTasks(mockDB())
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest("POST", "/api/lists/:id/tasks", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockTaskCreation)
	if err != nil {
		fmt.Println(err)
	}
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, string(j), got)
}

func TestGetTasks(t *testing.T) {
	// Setup
	e := gin.Default()
	handler := GetTasks(mockDB())

	e.GET("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest("GET", "/api/lists/:id/tasks", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockTask)
	if err != nil {
		fmt.Println(err)
	}
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(j), got)
}

func TestGetLists(t *testing.T) {
	// Setup
	e := gin.Default()

	_, err2 := mockDB().UpdateUsersById(context.Background(), 1)
	if err2 != nil {
		fmt.Println(err2)
	}

	handler := GetLists(mockDB())

	e.GET("/api/lists", handler)

	r := httptest.NewRequest("GET", "/api/lists", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockLists)
	if err != nil {
		fmt.Println(err)
	}
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(j), got)
}

func CreateTestCSV() {
	csvFile, err := os.Create("want_task.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	write := csv.NewWriter(csvFile)
	defer write.Flush()

	records := [][]string{}
	tasks := []string{}
	tasks = append(tasks, "Tasks", "Test Task")
	records = append(records, tasks)

	for _, record := range records[0] {
		row := []string{record}
		if err := write.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}

func TestProduceCSV(t *testing.T) {
	e := gin.Default()

	filename := "got_task.csv"
	handler := ProduceCSV(mockDB(), filename)
	e.GET("/api/list/export", handler)

	r := httptest.NewRequest("GET", "/api/list/export", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("Should have created a CSV file")
	}

	got, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	CreateTestCSV()
	want, err := ioutil.ReadFile("want_task.csv")
	if err != nil {
		fmt.Println(err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("CSV file should have correct content")
	}
	assert.Equal(t, 200, w.Code)
}

func TestDeleteTasks(t *testing.T) {
	// Setup
	e := gin.Default()

	handler := DeleteTask(mockDB())
	e.DELETE("/api/tasks/:id", handler)

	r := httptest.NewRequest("DELETE", "/api/tasks/:id", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	got := strings.TrimSuffix(w.Body.String(), "\n")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "0", got)
}

func TestDeleteLists(t *testing.T) {
	// Setup
	e := gin.Default()

	handler := DeleteList(mockDB())
	e.DELETE("/api/lists/:id", handler)

	r := httptest.NewRequest("DELETE", "/api/lists/:id", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	got := strings.TrimSuffix(w.Body.String(), "\n")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "0", got)
}

func TestPatchTasks(t *testing.T) {
	// Setup
	e := gin.Default()

	handler := PatchTasks(mockDB())
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest("POST", "/api/lists/:id/tasks", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	got := strings.TrimSuffix(w.Body.String(), "\n")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "0", got)
}

var (
	mockWeatherInfo = weather.WeatherInfo{
		FormattedTemp: "32.190000",
		Description:   "clear sky",
		City:          "Turabah",
	}
)

func TestGetWeather(t *testing.T) {
	e := gin.Default()

	handler := GetWeather(mockWeatherInfo)
	e.GET("/api/weather", handler)

	r := httptest.NewRequest("GET", "/api/weather", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockWeatherInfo)
	if err != nil {
		fmt.Println(err)
	}
	got := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(j), got)
}

func TestEmptyDBAfterTesting(t *testing.T) {
	mockDB().DeleteUserByUsername(context.Background(), "test")
}

// go test . -v -cover
// === RUN   TestGetUserByUsername
// --- PASS: TestGetUserByUsername (0.04s)
// === RUN   TestPostLists
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] POST   /api/lists                --> final/cmd/gin/handlers.PostList.func1 (3 handlers)
// [GIN] 2022/05/17 - 10:05:02 | 201 |      14.775ms |       192.0.2.1 | POST     "/api/lists"
// --- PASS: TestPostLists (0.04s)
// === RUN   TestPostTasks
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] POST   /api/lists/:id/tasks      --> final/cmd/gin/handlers.PostTasks.func1 (3 handlers)
// [GIN] 2022/05/17 - 10:05:02 | 201 |     13.2876ms |       192.0.2.1 | POST     "/api/lists/:id/tasks"
// --- PASS: TestPostTasks (0.02s)
// === RUN   TestGetTasks
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] GET    /api/lists/:id/tasks      --> final/cmd/gin/handlers.GetTasks.func1 (3 handlers)
// [GIN] 2022/05/17 - 10:05:02 | 200 |            0s |       192.0.2.1 | GET      "/api/lists/:id/tasks"
// --- PASS: TestGetTasks (0.01s)
// === RUN   TestGetLists
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] GET    /api/lists                --> final/cmd/gin/handlers.GetLists.func1 (3 handlers)
// [GIN] 2022/05/17 - 10:05:02 | 200 |       520.7µs |       192.0.2.1 | GET      "/api/lists"
// --- PASS: TestGetLists (0.00s)
// === RUN   TestProduceCSV
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] GET    /api/list/export          --> final/cmd/gin/handlers.ProduceCSV.func1 (3 handlers)
// [GIN] 2022/05/17 - 10:05:02 | 200 |     50.9142ms |       192.0.2.1 | GET      "/api/list/export"
// --- PASS: TestProduceCSV (0.06s)
// === RUN   TestDeleteTasks
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] DELETE /api/tasks/:id            --> final/cmd/gin/handlers.DeleteTask.func1 (3 handlers)
// [GIN] 2022/05/17 - 10:05:02 | 200 |     13.3691ms |       192.0.2.1 | DELETE   "/api/tasks/:id"
// --- PASS: TestDeleteTasks (0.01s)
// === RUN   TestDeleteLists
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] DELETE /api/lists/:id            --> final/cmd/gin/handlers.DeleteList.func1 (3 handlers)
// [GIN] 2022/05/17 - 10:05:02 | 200 |     11.5398ms |       192.0.2.1 | DELETE   "/api/lists/:id"
// --- PASS: TestDeleteLists (0.02s)
// === RUN   TestPatchTasks
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] POST   /api/lists/:id/tasks      --> final/cmd/gin/handlers.PatchTasks.func1 (3 handlers)
// [GIN] 2022/05/17 - 10:05:02 | 200 |            0s |       192.0.2.1 | POST     "/api/lists/:id/tasks"
// --- PASS: TestPatchTasks (0.01s)
// === RUN   TestGetWeather
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode)

// [GIN-debug] GET    /api/weather              --> final/cmd/gin/handlers.GetWeather.func1 (3 handlers)
// [GIN] 2022/05/17 - 10:05:02 | 200 |            0s |       192.0.2.1 | GET      "/api/weather"
// --- PASS: TestGetWeather (0.00s)
// === RUN   TestEmptyDBAfterTesting
// --- PASS: TestEmptyDBAfterTesting (0.01s)
// PASS
// coverage: 76.6% of statements
// ok      final/cmd/gin/handlers  0.669s  coverage: 76.6% of statements
