package main

import (
	"fmt"
	"log"
	"os"
	"sanjay/WedEase/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// TODO: move this into the server package.
func init() {
	// Disable Gin's debug messages
	gin.SetMode(gin.ReleaseMode)

}

// Entery point of our service.
func main() {

	router := gin.Default()

	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Define a handler for the root path
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	// Initialize the server.
	srv := server.ServerInit()

	srv.InjectRoutes(router)

	// Get the port from the environment variable or use a default port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Start the server
	fmt.Println("Server Started and running on port:", port)
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
