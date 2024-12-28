package main

import (
	"time"
	"todo-api/db"
	"todo-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	defer db.CloseDB()

	db.CreateTasksTable(db.DB)
	r := gin.Default()

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

	routes.SetupRoutes(r)

	r.Run(":8088")
}
