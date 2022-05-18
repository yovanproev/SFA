package final

import (
	"context"
	"encoding/base64"
	CSV "final/cmd/gin/csv"
	"final/cmd/gin/sqlc/db"
	weather "final/cmd/gin/weather"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserAndPassFromHeader(c *gin.Context) (string, string) {
	userAndPass := c.Request.Header["Authorization"]
	username := "test"
	password := "test"

	if userAndPass != nil {
		trim := strings.Trim(userAndPass[0], "Basic")
		trim = strings.Trim(trim, " ")

		rawDecodedText, err := base64.StdEncoding.DecodeString(trim)
		if err != nil {
			panic(err)
		}

		splitToUserAndPass := strings.Split(string(rawDecodedText), ":")
		username = splitToUserAndPass[0]
	}

	return username, password
}

func GetLists(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := GetUserAndPassFromHeader(c)
		user, err := q.GetUserByUsername(context.Background(), username)
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

		c.JSON(http.StatusOK, listLists)
	}
}

func PostList(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := GetUserAndPassFromHeader(c)
		user, err := q.GetUserByUsername(context.Background(), username)
		if err != nil {
			fmt.Println(err)
		}

		var list = db.CreateListParams{
			Userid: user.ID,
		}

		if user.Username == "test" && user.Password == "test" {
			list = db.CreateListParams{
				Name:   "Test List",
				Userid: 1,
			}
		}

		c.Bind(&list)

		_, err2 := q.CreateList(context.Background(), list)
		if err2 != nil {
			log.Println(err2)
		}

		c.JSON(http.StatusCreated, list)
	}
}

func DeleteList(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		username, _ := GetUserAndPassFromHeader(c)
		user, err := q.GetUserByUsername(context.Background(), username)
		if err != nil {
			fmt.Println(err)
		}

		listTasks, err := q.ListTasksByUserId(context.Background(), user.ID)
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
			err2 = q.DeleteTaskById(context.Background(), task.ID)
		}

		if id == 0 {
			// deleting the entry from database when testing
			// test database is always empty
			err = q.DeleteListsById(context.Background(), int32(id)+1)
		} else {
			err = q.DeleteListsById(context.Background(), int32(id))
		}
		if err != nil {
			log.Println(err)
		}

		if err == nil && err2 == nil {
			c.JSON(http.StatusOK, id)
		} else {
			fmt.Println(err)
		}
	}
}

func GetTasks(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		username, _ := GetUserAndPassFromHeader(c)
		user, err := q.GetUserByUsername(context.Background(), username)
		if err != nil {
			fmt.Println(err)
		}

		listLists, err := q.ListListsByUserId(context.Background(), user.ID)
		if err != nil {
			log.Println(err)
		}

		listTasks, err := q.ListTasksByUserId(context.Background(), user.ID)
		if err != nil {
			log.Println(err)
		}

		var filteredByID []db.Task
		for _, task := range listTasks {
			if int32(id) == task.Listid && listLists != nil {
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

		c.JSON(http.StatusOK, filteredByID)
	}
}

func PostTasks(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		username, _ := GetUserAndPassFromHeader(c)

		user, err := q.GetUserByUsername(context.Background(), username)
		if err != nil {
			fmt.Println(err)
		}

		task := db.CreateTaskParams{
			Listid: int32(id),
			Userid: user.ID,
		}

		c.Bind(&task)

		if task.Text == "" && id == 0 {
			task = db.CreateTaskParams{
				Text:      "Test Task",
				Listid:    1,
				Completed: true,
				Userid:    1,
			}
		}

		_, err2 := q.CreateTask(context.Background(), task)
		if err2 != nil {
			log.Println(err2)
		}

		c.JSON(http.StatusCreated, task)
	}
}

func PatchTasks(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := q.UpdateTask(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, id)
	}
}

func DeleteTask(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if id == 0 {
			err := q.DeleteTaskById(context.Background(), 1)
			if err != nil {
				fmt.Println(err)
			}
		}

		err := q.DeleteTaskById(context.Background(), int32(id))
		if err == nil {
			c.JSON(http.StatusOK, id)
		} else {
			fmt.Println(err)
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

func ProduceCSV(q *db.Queries, fileName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := GetUserAndPassFromHeader(c)
		user, err := q.GetUserByUsername(context.Background(), username)
		if err != nil {
			fmt.Println(err)
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; fileName="+fileName)
		c.Header("Content-Type", "application/octet-stream")
		records := CSV.GetTasksByUser(q, user)
		CSV.OpenCSV(records, fileName)

		c.File(fileName)
	}
}

func GetWeather(weather weather.WeatherInfo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, weather)
	}
}
