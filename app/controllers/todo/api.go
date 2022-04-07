package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdenadai/gotodo/app/database"
	"github.com/rdenadai/gotodo/app/models"
)

func ListAllTodo(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var todos []models.Todo
	database.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"content": todos,
	})
}

func ListTodo(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	id := c.Params.ByName("id")

	var todo models.Todo
	database.DB.First(&todo, id)
	if todo.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"content": todo,
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "Resource not found.",
	})
}

func CreateTodo(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var bind models.Todo
	if err := c.ShouldBindJSON(&bind); err != nil {
		log.Println("panic occurred:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Could not create resource.",
		})
		return
	}

	todo := models.Todo{Title: bind.Title, Description: bind.Description, Done: bind.Done}
	database.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"content": todo,
	})
}

func EditTodo(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var bind models.Todo
	if err := c.ShouldBindJSON(&bind); err != nil {
		log.Println("panic occurred:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Could not edit resource.",
		})
		return
	}

	id := c.Params.ByName("id")

	var todo models.Todo
	database.DB.First(&todo, id)

	if bind.Title != "" {
		todo.Title = bind.Title
	}
	if bind.Description != "" {
		todo.Description = bind.Description
	}
	todo.Done = bind.Done

	database.DB.Save(&todo)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"content": todo,
	})
}

func DeleteTodo(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	id := c.Params.ByName("id")

	var todo models.Todo
	database.DB.First(&todo, id)
	if todo.ID > 0 {
		var todoDeleted models.Todo
		database.DB.Delete(&todoDeleted, id)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Resource deleted.",
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "Resource not found.",
	})
}
