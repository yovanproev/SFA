package final

import (
	"context"
	"database/sql"
	"encoding/json"
	final "final/cmd/echo/DBInit"
	"final/cmd/echo/sqlc/db"
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
	mockTask = []db.Task{{
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

func fakeDB() *db.Queries {
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
	CreateUser(fakeDB(), "test", "test")

	want := db.User{
		ID:       1,
		Username: "test",
		Password: "test",
	}

	got, err := fakeDB().GetUserByUsername(context.Background(), want.Username)
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

	_, err2 := fakeDB().UpdateUsers(context.Background(), 1)
	if err2 != nil {
		fmt.Println(err2)
	}

	handler := PostList(fakeDB())

	e.POST("/api/lists", handler)

	r := httptest.NewRequest(echo.POST, "/api/lists", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockListsCreation)
	if err != nil {
		fmt.Println(err)
	}
	cutN := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, string(j), cutN)
}

func TestPostTasks(t *testing.T) {
	// Setup
	e := echo.New()

	handler := PostTasks(fakeDB())
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest(echo.POST, "/api/lists/:id/tasks", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockTaskCreation)
	if err != nil {
		fmt.Println(err)
	}
	cutN := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, string(j), cutN)
}

func TestGetTasks(t *testing.T) {
	// Setup
	e := echo.New()
	handler := GetTasks(fakeDB())

	e.GET("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest(echo.GET, "/api/lists/:id/tasks", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockTask)
	if err != nil {
		fmt.Println(err)
	}
	cutN := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(j), cutN)
}

func TestGetLists(t *testing.T) {
	// Setup
	e := echo.New()

	_, err2 := fakeDB().UpdateUsers(context.Background(), 1)
	if err2 != nil {
		fmt.Println(err2)
	}

	handler := GetLists(fakeDB())

	e.GET("/api/lists", handler)

	r := httptest.NewRequest(echo.GET, "/api/lists", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	j, err := json.Marshal(mockLists)
	if err != nil {
		fmt.Println(err)
	}
	cutN := strings.TrimSuffix(w.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(j), cutN)
}

func TestDeleteLists(t *testing.T) {
	// Setup
	e := echo.New()

	handler := DeleteList(fakeDB())
	e.DELETE("/api/lists/:id", handler)

	r := httptest.NewRequest(echo.DELETE, "/api/lists/:id", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	cutN := strings.TrimSuffix(w.Body.String(), "\n")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "0", cutN)
}

func TestPatchTasks(t *testing.T) {
	// Setup
	e := echo.New()

	handler := PatchTasks(fakeDB())
	e.POST("/api/lists/:id/tasks", handler)

	r := httptest.NewRequest(echo.POST, "/api/lists/:id/tasks", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)

	cutN := strings.TrimSuffix(w.Body.String(), "\n")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "0", cutN)
}

func TestEmptyDBAfterTesting(t *testing.T) {
	fakeDB().DeleteUsers(context.Background(), "test")
}

// go test . -v -cover
// === RUN   TestGetUserByUsername
// --- PASS: TestGetUserByUsername (0.02s)
// === RUN   TestPostLists
// --- PASS: TestPostLists (0.04s)
// === RUN   TestPostTasks
// --- PASS: TestPostTasks (0.02s)
// === RUN   TestGetTasks
// --- PASS: TestGetTasks (0.01s)
// === RUN   TestGetLists
// --- PASS: TestGetLists (0.00s)
// === RUN   TestDeleteLists
// --- PASS: TestDeleteLists (0.04s)
// === RUN   TestPatchTasks
// --- PASS: TestPatchTasks (0.01s)
// === RUN   TestEmptyDBAfterTesting
// --- PASS: TestEmptyDBAfterTesting (0.02s)
// PASS
// coverage: 74.7% of statements
// ok      final/cmd/echo/handlers 0.569s  coverage: 74.7% of statements
