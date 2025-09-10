package main

import (
	"example/web-service-gin/routes"
)

func main() {
	router := routes.SetupRoutes()

	// Start the server on port 8080
	router.Run("localhost:8080")
}
