package handlers

import (
	"net/http"

	"example/web-service-gin/database"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"data":    database.Albums,
		"count":   len(database.Albums),
		"message": "Albums retrieved successfully",
		"status":  http.StatusOK,
	})
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client,
// then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range database.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{
				"data":    database.Albums[i],
				"message": "Album found",
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "album not found",
	})
}

// createAlbum adds an album from JSON received in the request body.
func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, a := range database.Albums {
		if a.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "album with this ID already exists",
			})
			return
		}
	}

	// Add the new album to the slice.
	database.Albums = append(database.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, gin.H{
		"data":    newAlbum,
		"message": "Album created successfully",
	})
}

// updateAlbum updates an existing album identified by ID with JSON data from the request body.
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON((http.StatusBadRequest), gin.H{
			"error": err.Error(),
		})
		return
	}
	for i, a := range database.Albums {
		if a.ID == id {
			database.Albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, gin.H{
				"data":    updatedAlbum,
				"message": "Album updated successfully",
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "album not found",
	})
}

// deleteAlbum removes an album identified by ID from the collection.
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range database.Albums {
		if a.ID == id {
			// Remove the album from the slice
			database.Albums = append(database.Albums[:i], database.Albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Album deleted successfully",
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "album not found",
	})
}

// GetAlbumsCount returns the total count of albums
func GetAlbumsCount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"count": len(database.Albums),
	})
}
