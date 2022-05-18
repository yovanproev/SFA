package main

import (
	"final/cmd"
	dbFinal "final/cmd/echo/DBInit"
	handlers "final/cmd/echo/handlers"
	login "final/cmd/echo/login"
	weather "final/cmd/echo/weather"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	q := dbFinal.OpenDBConnection("storage.db")

	router := echo.New()
	router.Use(middleware.Recover())

	router.Use(middleware.CORS())
	router.Use(middleware.BasicAuth(login.Login(q)))

	//Add your handler (API endpoint) registrations here
	router.GET("/api", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})

	router.GET("/api/lists", handlers.GetLists(q))
	router.POST("/api/lists", handlers.PostList(q))
	router.DELETE("/api/lists/:id", handlers.DeleteList(q))

	router.GET("/api/lists/:id/tasks", handlers.GetTasks(q))
	router.POST("/api/lists/:id/tasks", handlers.PostTasks(q))
	router.PATCH("/api/tasks/:id", handlers.PatchTasks(q))
	router.DELETE("/api/tasks/:id", handlers.DeleteTask(q))

	router.GET("/api/list/export", handlers.ProduceCSV(q, "tasks.csv"))

	apiKeys := weather.LoadEnv("production.env")
	fetchWeather := weather.FetchWeather(41.99646, 21.43141, apiKeys, "")
	router.GET("/api/weather", handlers.GetWeather(fetchWeather))

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
