package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/mrinjamul-admin/handlers"
	"gopkg.in/square/go-jose.v2"
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
	// testing endpoint
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// coming soon
	r.GET("/api/projects", handlers.GetProjectsHandler)

	// unauthorized POST endpoint for recieving messeges from viewers
	r.POST("/api/messages", handlers.AddMessegeHandler)

	// authorized handler
	authorized := r.Group("/")
	// auth0 used here
	authorized.Use(authRequired())
	// protected endpoints
	authorized.GET("/api/messages", handlers.GetMessagesHandler)
	authorized.DELETE("/api/messages/:id", handlers.DeleteMessageHandler)
	authorized.PUT("/api/messages", handlers.MarkAsReadHandler)

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

		var auth0Domain = "https://" + domain + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(c.Request)

		if err != nil {
			log.Println(err)
			terminateWithError(http.StatusUnauthorized, "token is not valid", c)
			return
		}
		c.Next()
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
