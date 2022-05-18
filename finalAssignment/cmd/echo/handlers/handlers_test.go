package final

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	final "final/cmd/echo/DBInit"
	login "final/cmd/echo/login"
	"final/cmd/echo/sqlc/db"
	weather "final/cmd/echo/weather"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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
	mockTask = []db.Task{{
		ID:        0,
		Text:      "Test Task",
		Listid:    1,
		Userid:    1,
		Completed: true,
	},
	}
	mockListsCreation = db.List{
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
	login.CreateUser(mockDB(), "test", "test")

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
	e := echo.New()

	handler := PostList(mockDB())

	e.POST("/api/lists", handler)

	r := httptest.NewRequest(echo.POST, "/api/lists", nil)
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
	e := echo.New()

	handler := PostTasks(mockDB())
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest(echo.POST, "/api/lists/:id/tasks", nil)
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
	e := echo.New()
	handler := GetTasks(mockDB())

	e.GET("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest(echo.GET, "/api/lists/:id/tasks", nil)
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
	e := echo.New()

	handler := GetLists(mockDB())

	e.GET("/api/lists", handler)

	r := httptest.NewRequest(echo.GET, "/api/lists", nil)
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
	e := echo.New()

	filename := "got_task.csv"
	handler := ProduceCSV(mockDB(), filename)
	e.GET("/api/list/export", handler)

	r := httptest.NewRequest(echo.GET, "/api/list/export", nil)
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
	e := echo.New()

	handler := DeleteTask(mockDB())
	e.DELETE("/api/tasks/:id", handler)

	r := httptest.NewRequest(echo.DELETE, "/api/tasks/:id", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	got := strings.TrimSuffix(w.Body.String(), "\n")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "0", got)
}

func TestDeleteLists(t *testing.T) {
	// Setup
	e := echo.New()

	handler := DeleteList(mockDB())
	e.DELETE("/api/lists/:id", handler)

	r := httptest.NewRequest(echo.DELETE, "/api/lists/:id", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	got := strings.TrimSuffix(w.Body.String(), "\n")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "0", got)
}

func TestPatchTasks(t *testing.T) {
	// Setup
	e := echo.New()

	handler := PatchTasks(mockDB())
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest(echo.POST, "/api/lists/:id/tasks", nil)
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
	e := echo.New()

	handler := GetWeather(mockWeatherInfo)
	e.GET("/api/weather", handler)

	r := httptest.NewRequest(echo.GET, "/api/weather", nil)
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

func TestEmptyTheDBAfterTests(t *testing.T) {
	mockDB().DeleteUserByUsername(context.Background(), "test")
}

// $ go test . -v -cover
// === RUN   TestGetUserByUsername
// --- PASS: TestGetUserByUsername (0.04s)
// === RUN   TestPostLists
// --- PASS: TestPostLists (0.04s)
// === RUN   TestPostTasks
// --- PASS: TestPostTasks (0.02s)
// === RUN   TestGetTasks
// --- PASS: TestGetTasks (0.01s)
// === RUN   TestGetLists
// --- PASS: TestGetLists (0.00s)
// === RUN   TestProduceCSV
// --- PASS: TestProduceCSV (0.05s)
// === RUN   TestDeleteTasks
// --- PASS: TestDeleteTasks (0.02s)
// === RUN   TestDeleteLists
// --- PASS: TestDeleteLists (0.03s)
// === RUN   TestPatchTasks
// --- PASS: TestPatchTasks (0.01s)
// === RUN   TestGetWeather
// --- PASS: TestGetWeather (0.00s)
// === RUN   TestEmptyTheDBAfterTests
// --- PASS: TestEmptyTheDBAfterTests (0.02s)
// PASS
// coverage: 73.4% of statements
// ok      final/cmd/echo/handlers 0.625s  coverage: 76.5% of statements
