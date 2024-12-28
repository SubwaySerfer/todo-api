package main

import (
	"time"
	"todo-api/auth"
	"todo-api/db"
	"todo-api/middleware"
	"todo-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	defer db.CloseDB()

	db.CreateTasksTable(db.DB)
	r := gin.Default()

	r.POST("/login", auth.LoginHandler)

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5500"},        // Разрешённые домены
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Разрешённые методы
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
			"HX-Request",
			"HX-Target",
			"HX-Trigger",
			"HX-Trigger-Name",
			"HX-Current-URL", // Добавляем hx-current-url
		},
		ExposeHeaders:    []string{"Content-Length"}, // Заголовки, видимые клиенту
		AllowCredentials: true,                       // Разрешение на передачу куков
		MaxAge:           12 * time.Hour,             // Кэширование CORS настроек
	}))

	{
		protected.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the protected area!"})
		})
	}

	routes.SetupRoutes(r)

	r.Run(":8088")
}
