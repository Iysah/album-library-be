package routes

import (
	"example/web-service-gin/handlers"
	"example/web-service-gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Recovery())

	// album routes
	albumRoutes := router.Group("/albums")
	{
		albumRoutes.GET("", handlers.GetAlbums)
		albumRoutes.POST("", handlers.CreateAlbum)
		albumRoutes.GET("/:id", handlers.GetAlbumByID)
		albumRoutes.PUT("/:id", handlers.UpdateAlbum)
		albumRoutes.DELETE("/:id", handlers.DeleteAlbum)

		// tracks
		albumRoutes.GET("/:id/tracks", handlers.GetAlbumTracks)
		albumRoutes.POST("/:id/tracks", handlers.CreateTrack)
		albumRoutes.GET("/:id/tracks/:trackId", handlers.GetTrackByID)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Server is running!",
		})
	})

	return router
}
