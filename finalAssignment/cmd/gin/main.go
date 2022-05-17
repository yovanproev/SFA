package main

import (
	"final/cmd"
	dbFinal "final/cmd/gin/DBInit"
	handlers "final/cmd/gin/handlers"
	login "final/cmd/gin/login"
	weather "final/cmd/gin/weather"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	q := dbFinal.OpenDBConnection("storage.db")

	router := gin.Default()

	authorized := router.Group("/")
	authorized.Use(gin.BasicAuth(login.GinAccounts(q)))

	authorized.Use(login.Login(q))

	authorized.GET("/api", func(c *gin.Context) {
	})

	authorized.GET("/api/lists", handlers.GetLists(q))
	authorized.POST("/api/lists", handlers.PostList(q))
	authorized.DELETE("/api/lists/:id", handlers.DeleteList(q))

	authorized.GET("/api/lists/:id/tasks", handlers.GetTasks(q))
	authorized.POST("/api/lists/:id/tasks", handlers.PostTasks(q))
	authorized.PATCH("/api/tasks/:id", handlers.PatchTasks(q))
	authorized.DELETE("/api/tasks/:id", handlers.DeleteTask(q))

	authorized.GET("/api/list/export", handlers.ProduceCSV(q, "tasks.csv"))

	apiKeys := weather.LoadEnv("production.env")
	fetchWeather := weather.FetchWeather(41.99646, 21.43141, apiKeys, "")
	authorized.GET("/api/weather", handlers.GetWeather(fetchWeather))

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
