package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/suhailkassar11/image_processing_system/initializers"
	"github.com/suhailkassar11/image_processing_system/routes"
)

func init() {
	// Load .env file
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	fmt.Println("welcome to the image")
	r := gin.Default()
	routes.SetupUserRoutes(r)
	host := "localhost"
	port := "8080"
	url := fmt.Sprintf("http://%s:%s", host, port)

	// Print the URL to the terminal
	fmt.Printf("Server is running at %s\n", url)

	// Start the server
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
	r.Run()
}
