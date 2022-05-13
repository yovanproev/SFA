package main

import (
	"context"
	"crypto/subtle"
	"final/cmd"
	CSV "final/cmd/echo/CSV"
	dbFinal "final/cmd/echo/DBInit"
	routesTasks "final/cmd/echo/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"tawesoft.co.uk/go/dialog"
)

func main() {
	q := dbFinal.OpenDBConnection("storage.db")

	router := echo.New()

	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// This is a sample demonstration of how to attach middlewares in Echo
		return func(ctx echo.Context) error {
			log.Println("Echo middleware was called")
			return next(ctx)
		}
	})

	router.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// check if the user exists in the DB
		user, err := q.GetUserByUsername(context.Background(), username)
		if err != nil {
			fmt.Println(err)
		}

		// create new user if it doesn't exist
		if user.Username != username {
			routesTasks.CreateUser(q, username, password)

			// create a new login data for the user
			_, err2 := q.UpdateUsers(context.Background(), user.ID)
			if err2 != nil {
				fmt.Println(err2)
			}
			return true, nil
		}

		// if username exists and the password is incorrect
		if user.Username == username && user.Password != password {
			dialog.Alert("You password is incorrect, try again!")
			return false, nil
		}

		// if username and password are a match
		if subtle.ConstantTimeCompare([]byte(username), []byte(user.Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(user.Password)) == 1 {

			// update the login
			_, err2 := q.UpdateUsers(context.Background(), user.ID)
			if err2 != nil {
				fmt.Println(err2)
			}

			CSV.OpenCSV(q, user)
			return true, nil
		}

		return false, nil
	}))

	router.Use(middleware.CORS())

	//Add your handler (API endpoint) registrations here
	router.GET("/api", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})

	router.GET("/api/lists", routesTasks.GetLists(q))
	router.POST("/api/lists", routesTasks.PostList(q))
	router.DELETE("/api/lists/:id", routesTasks.DeleteList(q))

	router.GET("/api/lists/:id/tasks", routesTasks.GetTasks(q))
	router.POST("/api/lists/:id/tasks", routesTasks.PostTasks(q))
	router.PATCH("/api/tasks/:id", routesTasks.PatchTasks(q))
	router.DELETE("/api/tasks/:id", routesTasks.DeleteTask(q))

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
