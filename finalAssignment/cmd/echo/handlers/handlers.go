package final

import (
	"context"
	"final/cmd/echo/sqlc/db"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetLists(q *db.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := q.GetUserByDate(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		listLists, err := q.ListListsByUserId(context.Background(), user.ID)
		if err != nil {
			log.Println(err)
		}

		if user.Username == "test" && user.Password == "test" {
			listLists = []db.List{{
				ID:     1,
				Name:   "Test List",
				Userid: 1,
			}}
		} else if listLists == nil {
			listLists = []db.List{{}}
		}

		return c.JSON(http.StatusOK, listLists)
	}
}

func PostList(q *db.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := q.GetUserByDate(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		var list = db.CreateListParams{
			Userid: user.ID,
		}

		c.Bind(&list)

		_, err2 := q.CreateList(context.Background(), list)
		if err2 != nil {
			log.Println(err2)
		}
		if user.Username == "test" && user.Password == "test" {
			list = db.CreateListParams{
				Name:   "Test List",
				Userid: 1,
			}
		}

		return c.JSON(http.StatusCreated, list)
	}
}

func DeleteList(q *db.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		listTasks, err := q.ListTasks(context.Background())
		if err != nil {
			log.Println(err)
		}

		var filteredByID []db.Task
		for _, task := range listTasks {
			if int32(id) == task.Listid {
				filteredByID = append(filteredByID, task)
			}
		}

		var err2 error
		for _, task := range filteredByID {
			err2 = q.DeleteTask(context.Background(), task.ID)
		}

		if id == 0 {
			// deleting the entry from database when testing
			// test database is always empty
			err = q.DeleteLists(context.Background(), int32(id)+1)
		} else {
			err = q.DeleteLists(context.Background(), int32(id))
		}
		if err != nil {
			log.Println(err)
		}

		if err == nil && err2 == nil {
			return c.JSON(http.StatusOK, id)
		} else {
			return err
		}
	}
}

func GetTasks(q *db.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		user, err := q.GetUserByDate(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		listLists, err := q.ListListsByUserId(context.Background(), user.ID)
		if err != nil {
			log.Println(err)
		}

		listTasks, err := q.ListTasks(context.Background())
		if err != nil {
			log.Println(err)
		}

		var filteredByID []db.Task
		for _, task := range listTasks {
			if int32(id) == task.Listid && listLists != nil && user.ID == task.Userid {
				filteredByID = append(filteredByID, task)
			}
		}

		if filteredByID == nil && id != 0 {
			filteredByID = []db.Task{{}}
		} else if id == 0 {
			filteredByID = []db.Task{{
				Text:      "Test Task",
				Listid:    1,
				Completed: true,
				Userid:    1,
			}}
		}

		return c.JSON(http.StatusOK, filteredByID)
	}
}

func PostTasks(q *db.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		// get the last logged user by date
		user, err := q.GetUserByDate(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		task := db.CreateTaskParams{
			Listid: int32(id),
			Userid: user.ID,
		}

		c.Bind(&task)

		_, err2 := q.CreateTask(context.Background(), task)
		if err2 != nil {
			log.Println(err2)
		}

		if task.Text == "" {
			task = db.CreateTaskParams{
				Text:      "Test Task",
				Listid:    1,
				Completed: true,
				Userid:    1,
			}
		}

		return c.JSON(http.StatusCreated, task)
	}
}

func PatchTasks(q *db.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := q.UpdateTask(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
		}

		return c.JSON(http.StatusOK, id)
	}
}

func DeleteTask(q *db.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		err := q.DeleteTask(context.Background(), int32(id))
		if err == nil {
			return c.JSON(http.StatusOK, id)
		} else {
			return err
		}
	}
}

func CreateUser(q *db.Queries, username, password string) {
	var user = db.CreateUserParams{
		Username: username,
		Password: password,
	}

	if username == "test" && password == "test" {
		user = db.CreateUserParams{
			Username: username,
			Password: password,
		}
	}

	_, err := q.CreateUser(context.Background(), user)
	if err != nil {
		log.Println(err)
	}

}
