package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rdenadai/gotodo/app/controllers"
)

func StartServer() {
	r := gin.Default()
	r.LoadHTMLGlob("app/templates/**/*.html")

	// API
	r.GET("/api/v1/todos", controllers.ListAllTodo)
	r.GET("/api/v1/todos/:id", controllers.ListTodo)
	r.POST("/api/v1/todos", controllers.CreateTodo)
	r.PATCH("/api/v1/todos/:id", controllers.EditTodo)
	r.DELETE("/api/v1/todos/:id", controllers.DeleteTodo)

	// Templates
	r.GET("/", controllers.RenderHomePage)
	r.GET("/add", controllers.RenderAddPage)
	r.POST("/add", controllers.ExecuteAddPage)
	r.GET("/edit/:id", controllers.RenderEditPage)
	r.POST("/edit", controllers.ExecuteEditPage)

	// Execute
	r.Run()
}
