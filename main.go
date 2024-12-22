package main

import (
    "todo-api/routes"
		"todo-api/db"
    "github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	db.CreateTasksTable(db.DB)
  r := gin.Default()

	routes.SetupRoutes(r)

  r.Run(":8080")
}