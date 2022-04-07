package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rdenadai/gotodo/app/database"
	"github.com/rdenadai/gotodo/app/models"
)

func RenderHomePage(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)

	c.HTML(http.StatusOK, "todo/index.html", todos)
}

func RenderAddPage(c *gin.Context) {
	c.HTML(http.StatusOK, "todo/form.html", nil)
}

func RenderEditPage(c *gin.Context) {
	id := c.Params.ByName("id")

	var todo models.Todo
	database.DB.First(&todo, id)

	c.HTML(http.StatusOK, "todo/form.html", todo)
}

func ExecuteAddPage(c *gin.Context) {
	var bind models.Todo
	if err := c.ShouldBindWith(&bind, binding.FormMultipart); err != nil {
		log.Println("panic occurred:", err)
		c.Redirect(http.StatusBadRequest, "/")
		return
	}

	todo := models.Todo{Title: bind.Title, Description: bind.Description, Done: bind.Done}
	database.DB.Create(&todo)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func ExecuteEditPage(c *gin.Context) {
	var bind models.Todo
	if err := c.ShouldBindWith(&bind, binding.FormMultipart); err != nil {
		log.Println("panic occurred:", err)
		c.Redirect(http.StatusBadRequest, "/")
		return
	}

	var todo models.Todo
	database.DB.First(&todo, bind.ID)

	todo.Title = bind.Title
	todo.Description = bind.Description

	database.DB.Save(&todo)

	c.Redirect(http.StatusMovedPermanently, "/")
}
