package handlers

import (
	"fmt"
	"net/http"

	"example/web-service-gin/database"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// GetAlbumTracks responds with the list of tracks for a specific album as JSON.
func GetAlbumTracks(c *gin.Context) {
	albumId := c.Param("id")
	album, _ := database.GetAlbumByID(albumId)

	if album == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Album not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    album.Tracks,
		"count":   len(album.Tracks),
		"message": fmt.Sprintf("Tracks for album %s retrieved successfully", albumId),
		"status":  http.StatusOK,
	})
}

// GetTrackByID responds with a specific track from a specific album as JSON.
func GetTrackByID(c *gin.Context) {
	albumId := c.Param("id")
	trackId := c.Param("trackId")

	track, _, _ := database.GetTrackByID(albumId, trackId)
	if track == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Track not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    track,
		"message": "Track found successfully",
	})
}

func CreateTrack(c *gin.Context) {
	albumID := c.Param("id")
	album, albumIndex := database.GetAlbumByID(albumID)

	if album == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Album not found!",
		})
		return
	}

	validatedData, exists := c.Get("validatedTrack")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Validation data not found",
		})
		return
	}

	data := validatedData.(struct {
		Title       string `json:"title" validate:"required,min=1,max=100"`
		Duration    string `json:"duration" validate:"required,duration"`
		TrackNumber int    `json:"track_number" validate:"required,min=1"`
		AudioURL    string `json:"audio_url" validate:"omitempty,url"`
	})

	for _, track := range album.Tracks {
		if track.TrackNumber == data.TrackNumber {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Track with this track number already exists in the album",
			})
			return
		}
	}

	// Generate a new track ID
	trackID := fmt.Sprintf("%s-%d", albumID, len(album.Tracks)+1)
	newTrack := models.Track{
		ID:          trackID,
		Title:       data.Title,
		Duration:    data.Duration,
		TrackNumber: data.TrackNumber,
		AudioURL:    data.AudioURL,
		AlbumID:     albumID,
	}

	database.Albums[albumIndex].Tracks = append(database.Albums[albumIndex].Tracks, newTrack)
	// Update the track count
	database.Albums[albumIndex].TrackCount = len(database.Albums[albumIndex].Tracks)

	c.JSON(http.StatusCreated, gin.H{
		"data":    newTrack,
		"message": "Track created successfully",
	})
}
