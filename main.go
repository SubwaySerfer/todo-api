package main

import (
    // "todo-api/routes"
    "todo-api/models"
    "github.com/gin-gonic/gin"
		"net/http"
)

func main() {
    r := gin.Default()

    // Подключаем маршруты
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

		c.JSON(http.StatusOK, gin.H{
			"message": "Task created",
			"task": task,
		})
	})

    // Запускаем сервер
    r.Run(":8080") // Сервер слушает порт 8080
}
