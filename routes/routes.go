package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	RegisterTaskRoutes(router)
}
