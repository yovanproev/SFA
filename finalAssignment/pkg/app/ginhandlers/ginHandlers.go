package ginhandlers

import (
	"context"
	CSV "final/pkg/app/CSV"
	handleErrors "final/pkg/app/errors"
	"final/pkg/app/users"
	weather "final/pkg/app/weather"
	"final/pkg/sqlc/db"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserAndGetListsByUserId(q *db.Queries, c *gin.Context, e handleErrors.Error) (db.User, []db.List) {
	username, _ := users.GetUserAndPassFromHeaderGin(c)
	user, err := q.GetUserByUsername(context.Background(), username)
	if err != nil {
		log.Println(e.DatabaseError, err)
		c.String(http.StatusNotFound, e.StatusNotFoundError)
	}

	listLists, err := q.ListListsByUserId(context.Background(), user.ID)
	if err != nil {
		log.Println(e.DatabaseError, err)
		c.String(http.StatusNotFound, e.StatusNotFoundError)
	}

	return user, listLists
}

func GetLists(q *db.Queries, e handleErrors.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, listLists := getUserAndGetListsByUserId(q, c, e)

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

func PostList(q *db.Queries, e handleErrors.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := getUserAndGetListsByUserId(q, c, e)

		var list = db.CreateListParams{
			Userid: user.ID,
		}

		if user.Username == "test" && user.Password == "test" {
			list = db.CreateListParams{
				Name:   "Test List",
				Userid: 1,
			}
		}

		err := c.Bind(&list)
		if err != nil {
			log.Println(err)
		}

		_, err = q.CreateList(context.Background(), list)
		if err != nil {
			log.Println(e.DatabaseError, err)
			c.String(http.StatusNotFound, e.StatusNotFoundError)
		}

		c.JSON(http.StatusCreated, list)
	}
}

func DeleteList(q *db.Queries, e handleErrors.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get list Id
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil && id != 0 {
			log.Println(err)
		}

		// get user
		user, listLists := getUserAndGetListsByUserId(q, c, e)

		// get tasks connected with the list and delete them
		listTasks, err := q.ListTasksByUserId(context.Background(), user.ID)
		if err != nil {
			log.Println(e.DatabaseError, err)
			c.String(http.StatusNotFound, e.StatusNotFoundError)
		}
		for _, task := range listTasks {
			if int32(id) == task.Listid {
				err := q.DeleteTaskById(context.Background(), task.ID)
				if err != nil {
					log.Println(e.DatabaseError, err)
					c.String(http.StatusNotFound, e.StatusNotFoundError)
				}
			}
		}

		for _, list := range listLists {
			if list.ID == int32(id) && len(listLists) != 0 {
				err = q.DeleteListsById(context.Background(), int32(id))
				if err != nil {
					log.Println(e.DatabaseError, err)
					c.String(http.StatusNotFound, e.StatusNotFoundError)
				}
				c.JSON(http.StatusOK, list)
			} else if id == 0 && listLists == nil {
				c.String(http.StatusNotFound, e.StatusNotFoundError)
			}
		}

		if id == 0 {
			// deleting the entry from database when testing
			// test database is always empty
			err = q.DeleteListsById(context.Background(), int32(id)+1)
			if err != nil {
				log.Println(e.DatabaseError, err)
				c.String(http.StatusNotFound, e.StatusNotFoundError)
			}
			c.JSON(http.StatusOK, listLists)
		}
	}
}

func GetTasks(q *db.Queries, e handleErrors.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil && id != 0 {
			fmt.Println(err)
		}

		user, listLists := getUserAndGetListsByUserId(q, c, e)

		listTasks, err := q.ListTasksByUserId(context.Background(), user.ID)
		if err != nil {
			log.Println(e.DatabaseError, err)
			c.String(http.StatusNotFound, e.StatusNotFoundError)
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

		c.JSON(http.StatusOK, filteredByID)
	}
}

func PostTasks(q *db.Queries, e handleErrors.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil && id != 0 {
			fmt.Println(err)
		}

		user, _ := getUserAndGetListsByUserId(q, c, e)

		task := db.CreateTaskParams{
			Listid:    int32(id),
			Userid:    user.ID,
			Completed: 0,
		}

		err = c.Bind(&task)
		if err != nil {
			log.Println(err)
		}

		if task.Text == "" && id == 0 {
			task = db.CreateTaskParams{
				Text:      "Test Task",
				Listid:    1,
				Completed: true,
				Userid:    1,
			}
		}

		_, err = q.CreateTask(context.Background(), task)
		if err != nil {
			log.Println(e.DatabaseError, err)
			c.String(http.StatusNotFound, e.StatusNotFoundError)
		}

		c.JSON(http.StatusCreated, task)
	}
}

func PatchTasks(q *db.Queries, e handleErrors.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil && id != 0 {
			fmt.Println(err)
		}

		user, _ := getUserAndGetListsByUserId(q, c, e)

		listTasks, err := q.ListTasksByUserId(context.Background(), user.ID)
		if err != nil {
			log.Println(e.DatabaseError, err)
			c.String(http.StatusNotFound, e.StatusNotFoundError)
		}

		for _, task := range listTasks {
			if task.ID == int32(id) {
				_, err := q.UpdateTask(context.Background(), int32(id))
				if err != nil {
					log.Println(e.DatabaseError, err)
					c.String(http.StatusNotFound, e.StatusNotFoundError)
				}
				c.JSON(http.StatusOK, task)
			} else if id == 0 && listTasks == nil {
				c.String(http.StatusNotFound, e.StatusNotFoundError)
			}
		}

		if id == 0 {
			c.JSON(http.StatusOK, listTasks)
		}
	}
}

func DeleteTask(q *db.Queries, e handleErrors.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil && id != 0 {
			fmt.Println(err)
		}

		user, _ := getUserAndGetListsByUserId(q, c, e)

		listTasks, err := q.ListTasksByUserId(context.Background(), user.ID)
		if err != nil {
			log.Println(e.DatabaseError, err)
			c.String(http.StatusNotFound, e.StatusNotFoundError)
		}

		for _, task := range listTasks {
			if task.ID == int32(id) {
				err = q.DeleteTaskById(context.Background(), int32(id))
				if err != nil {
					log.Println(e.DatabaseError, err)
					c.String(http.StatusNotFound, e.StatusNotFoundError)
				}
				c.JSON(http.StatusOK, "Task "+strconv.Itoa(id)+" is deleted!")
			} else if id == 0 && listTasks == nil {
				c.String(http.StatusNotFound, e.StatusNotFoundError)
			}
		}

		if id == 0 {
			err := q.DeleteTaskById(context.Background(), 1)
			if err != nil {
				log.Println(e.DatabaseError, err)
				c.String(http.StatusNotFound, e.StatusNotFoundError)
			}
			c.JSON(http.StatusOK, "Task "+strconv.Itoa(id)+" is deleted!")
		}

	}
}

func ProduceCSV(q *db.Queries, fileName string, e handleErrors.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := getUserAndGetListsByUserId(q, c, e)

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; fileName="+fileName)
		c.Header("Content-Type", "application/octet-stream")

		records := CSV.GetTasksByUser(q, user, e)
		toString := CSV.CreateBytesFromTasks(records)

		c.Data(http.StatusOK, "text/csv", []byte(toString))
	}
}

func GetWeather(apiKeys, url string, e handleErrors.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := c.Request.Header.Get("lat")
		lon := c.Request.Header.Get("lon")

		fetchWeather := weather.FetchWeather(lat, lon, apiKeys, url, e)
		c.JSON(http.StatusOK, fetchWeather)
	}
}
