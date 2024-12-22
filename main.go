package main

import (
    // "todo-api/routes"
		"todo-api/db"
    "todo-api/models"
    "github.com/gin-gonic/gin"
		"net/http"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	db.CreateTasksTable(db.DB)
  r := gin.Default()

// routes.SetupRoutes(r)

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.POST("/task", func(c *gin.Context) {
		var task models.Task

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id"

		var id int
		err := db.DB.QueryRow(query, task.Title, task.Description).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Task created",
			"task": gin.H{
				"id": id,
				"title": task.Title,
				"description": task.Description,
			},
		})
	})

  r.Run(":8080")
}