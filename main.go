package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	audience string
	domain   string
)

func main() {
	// set auth0variables
	setAuth0Variables()

	// set port from env
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "3000"
	}
	r := gin.Default()

	r.Use(CORSMiddleware())

	// This will ensure that the angular files are served correctly
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./static/index.html")
		} else {
			c.File("./static/" + path.Join(dir, file))
		}
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	// TODO: write authorized handler

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}

}

func setAuth0Variables() {
	audience = os.Getenv("AUTH0_API_IDENTIFIER")
	domain = os.Getenv("AUTH0_DOMAIN")
}

// ValidateRequest will verify that a token received from an http request
// is valid and signyed by Auth0
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

func terminateWithError(statusCode int, message string, c *gin.Context) {
	c.JSON(statusCode, gin.H{"error": message})
	c.Abort()
}

// CORSMiddleware : cross origin resource sharing
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
