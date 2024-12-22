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

	r.GET("/task/:id", func(c *gin.Context) {
		id := c.Param("id")

		var task models.Task
		query := "SELECT id, title, description FROM tasks WHERE id = $1"

		err := db.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"task": task,
		})
	})

	r.DELETE("/task/:id", func(c *gin.Context) {
		id := c.Param("id")

		query := "DELETE FROM tasks WHERE id = $1"

		_, err := db.DB.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Task deleted",
		})
	})

	r.GET("/tasks", func(c *gin.Context) {
		var tasks []models.Task
		query := "SELECT id, title, description FROM tasks"

		rows, err := db.DB.Query(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query tasks: " + err.Error()})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var task models.Task
			if err := rows.Scan(&task.ID, &task.Title, &task.Description); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan task: " + err.Error()})
				return
			}
			tasks = append(tasks, task)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"tasks": tasks,
		})
	})

  r.Run(":8080")
}