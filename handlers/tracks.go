package handlers

import (
	"fmt"
	"net/http"

	"example/web-service-gin/database"
	"example/web-service-gin/models"

	"github.com/go-playground/validator/v10"

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

	var trackData struct {
		Title       string `json:"title" validate:"required,min=1,max=100"`
		Duration    string `json:"duration" validate:"required"`
		TrackNumber int    `json:"track_number" validate:"required,min=1"`
		AudioURL    string `json:"audio_url" validate:"omitempty,url"`
	}

	// Bind and validate JSON
	if err := c.ShouldBindJSON(&trackData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON format",
			"details": err.Error(),
		})
		return
	}

	// Validate the struct
	if err := validator.New().Struct(trackData); err != nil {
		errors := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, getValidationErrorMessage(err))
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": errors,
		})
		return
	}

	// Validate duration format
	if !isValidDuration(trackData.Duration) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Duration must be in MM:SS format",
		})
		return
	}

	for _, track := range album.Tracks {
		if track.TrackNumber == trackData.TrackNumber {
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
		Title:       trackData.Title,
		Duration:    trackData.Duration,
		TrackNumber: trackData.TrackNumber,
		AudioURL:    trackData.AudioURL,
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

// Helper functions
func getValidationErrorMessage(err validator.FieldError) string {
	field := err.Field()
	tag := err.Tag()

	switch tag {
	case "required":
		return field + " is required"
	case "min":
		if field == "TrackNumber" {
			return field + " must be greater than or equal to " + err.Param()
		}
		return field + " must be at least " + err.Param() + " characters long"
	case "max":
		return field + " must be at most " + err.Param() + " characters long"
	case "url":
		return field + " must be a valid URL"
	default:
		return field + " is invalid"
	}
}

func isValidDuration(duration string) bool {
	// Simple regex for MM:SS format
	matched := true
	if len(duration) < 4 || len(duration) > 5 {
		matched = false
	}

	// Check for : in the right position
	if matched && duration[len(duration)-3] != ':' {
		matched = false
	}

	return matched
}
