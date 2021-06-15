package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// port := "3000"
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "3000"
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Welcome to My Admin",
		})

	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	r.Run(":" + port) // Run gin

}
