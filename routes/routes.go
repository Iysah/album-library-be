package routes

import (
	"example/web-service-gin/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// album routes
	albumRoutes := router.Group("/albums")
	{
		albumRoutes.GET("", handlers.GetAlbums)
		albumRoutes.POST("", handlers.CreateAlbum)
		albumRoutes.GET("/:id", handlers.GetAlbumByID)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Server is running!",
		})
	})

	return router
}
