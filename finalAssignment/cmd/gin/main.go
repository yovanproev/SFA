package main

import (
	"final/cmd"
	dbFinal "final/cmd/gin/DBInit"
	routesTasks "final/cmd/gin/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	q := dbFinal.OpenDBConnection("storage.db")

	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		// This is a sample demonstration of how to attach middlewares in Gin
		log.Println("Gin middleware was called")
		ctx.Next()
	})

	// Add your handler (API endpoint) registrations here
	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello, World!")
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
