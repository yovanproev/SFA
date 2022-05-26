package main

import (
	"final/cmd"
	handleErrors "final/pkg/app/errors"
	handlers "final/pkg/app/ginhandlers"
	login "final/pkg/app/ginmiddleware"
	"final/pkg/config"
	db "final/pkg/sqlc/initDB"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	configuration := config.Configurations{}
	configuration.SetConfig()
	e := handleErrors.Error{}.SetErrors()

	q := db.OpenDBConnection(configuration.Database.ProductionDBName, e)

	router := gin.Default()

	authorized := router.Group("/")
	authorized.Use(gin.BasicAuth(login.GinAccounts(q)))

	authorized.GET("/api", func(c *gin.Context) {
	})

	authorized.GET("/api/lists", handlers.GetLists(q, e))
	authorized.POST("/api/lists", handlers.PostList(q, e))
	authorized.DELETE("/api/lists/:id", handlers.DeleteList(q, e))

	authorized.GET("/api/lists/:id/tasks", handlers.GetTasks(q, e))
	authorized.POST("/api/lists/:id/tasks", handlers.PostTasks(q, e))
	authorized.PATCH("/api/tasks/:id", handlers.PatchTasks(q, e))
	authorized.DELETE("/api/tasks/:id", handlers.DeleteTask(q, e))

	authorized.GET("/api/list/export", handlers.ProduceCSV(q, configuration.CSVName, e))

	apiKeys := config.LoadEnv(configuration.ProductionEnv, configuration)
	authorized.GET("/api/weather", handlers.GetWeather(apiKeys, "", e))

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(configuration.Server.Port, cmd.CreateCommonMux(router)))
}
