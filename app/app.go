package app

import (
	"github.com/gin-gonic/gin"
	todo "github.com/rdenadai/gotodo/app/controllers/todo"
)

func StartServer() {
	r := gin.Default()
	r.LoadHTMLGlob("app/templates/**/*.html")

	// API
	r.GET("/api/v1/todos", todo.ListAllTodo)
	r.GET("/api/v1/todos/:id", todo.ListTodo)
	r.POST("/api/v1/todos", todo.CreateTodo)
	r.PATCH("/api/v1/todos/:id", todo.EditTodo)
	r.DELETE("/api/v1/todos/:id", todo.DeleteTodo)

	// UI
	r.GET("/", todo.RenderHomePage)
	r.GET("/add", todo.RenderAddPage)
	r.POST("/add", todo.ExecuteAddPage)
	r.GET("/edit/:id", todo.RenderEditPage)
	r.POST("/edit", todo.ExecuteEditPage)

	// Execute
	r.Run()
}
