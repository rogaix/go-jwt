package main

import (
	"fmt"
	"log"
	"os"

	"jwt/handlers"
	"jwt/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
		return
	}

	err = os.Setenv("JWT_SECRET", "secret-key-here")
	if err != nil {
		fmt.Print("Error setting JWT_SECRET")
		return
	}
}

func main() {
	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	r.Use(cors.New(config))

	// Routes
	r.POST("/api/login", handlers.Login)
	r.POST("/api/register", handlers.Register)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/user", handlers.GetUser)
	}

	log.Fatal(r.Run(":8080"))
}
