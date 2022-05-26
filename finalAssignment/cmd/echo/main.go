package main

import (
	"final/cmd"
	handlers "final/pkg/app/echohandlers"
	login "final/pkg/app/echomiddleware"
	handleErrors "final/pkg/app/errors"
	"final/pkg/config"
	db "final/pkg/sqlc/initDB"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configuration := config.Configurations{}
	configuration.SetConfig()
	e := handleErrors.Error{}.SetErrors()

	q := db.OpenDBConnection(configuration.Database.ProductionDBName, e)

	router := echo.New()
	router.Use(middleware.Recover())

	router.Use(middleware.CORS())
	router.Use(middleware.BasicAuth(login.Authenticate(q, e)))

	//Add your handler (API endpoint) registrations here
	router.GET("/api/", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})

	router.GET("/api/lists", handlers.GetLists(q, e))
	router.POST("/api/lists", handlers.PostList(q, e))
	router.DELETE("/api/lists/:id", handlers.DeleteList(q, e))

	router.GET("/api/lists/:id/tasks", handlers.GetTasks(q, e))
	router.POST("/api/lists/:id/tasks", handlers.PostTasks(q, e))
	router.PATCH("/api/tasks/:id", handlers.PatchTasks(q, e))
	router.DELETE("/api/tasks/:id", handlers.DeleteTask(q, e))

	router.GET("/api/list/export", handlers.ProduceCSV(q, configuration.CSVName, e))

	apiKeys := config.LoadEnv(configuration.ProductionEnv, configuration)
	router.GET("/api/weather", handlers.GetWeather(apiKeys, "", e))

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(configuration.Server.Port, cmd.CreateCommonMux(router)))
}
