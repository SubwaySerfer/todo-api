package main

import (
    // "todo-api/routes"
    "todo-api/models"
    "github.com/gin-gonic/gin"
		"net/http"
		"log"
		"fmt"
		"github.com/joho/godotenv"
		"os"

		"database/sql"
		_ "github.com/lib/pq"
)

func main() {
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

		c.JSON(http.StatusOK, gin.H{
			"message": "Task created",
			"task": task,
		})
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	// Чтение переменных окружения
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, sslMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных!")

  r.Run(":8080")
}