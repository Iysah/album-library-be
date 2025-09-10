package database

import (
	"example/web-service-gin/models"
	"time"
)

var Albums = []models.Album{
	{
		ID:          "1",
		Title:       "Midnight Memories",
		Artist:      "One Direction",
		Price:       19.99,
		ReleaseDate: "2013-11-25",
		Genre:       "Pop Rock",
		TrackCount:  2,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Tracks: []models.Track{
			{
				ID:          "1-1",
				Title:       "Best Song Ever",
				Duration:    "3:22",
				TrackNumber: 1,
				AudioURL:    "https://link-to-audio.com/best-song-ever.mp3",
				AlbumID:     "1",
			},
			{
				ID:          "1-2",
				Title:       "Story of My Life",
				Duration:    "4:05",
				TrackNumber: 2,
				AudioURL:    "https://link-to-audio.com/story-of-my-life.mp3",
				AlbumID:     "1",
			},
		},
	},
	{
		ID:          "2",
		Title:       "Blue Train",
		Artist:      "John Coltrane",
		Price:       56.99,
		ReleaseDate: "1958-01-01",
		Genre:       "Jazz",
		TrackCount:  0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Tracks:      []models.Track{},
	},
}

// Helper functions for tracks
func GetAlbumByID(id string) (*models.Album, int) {
	for i, album := range Albums {
		if album.ID == id {
			return &Albums[i], i
		}
	}
	return nil, -1
}

func GetTrackByID(albumID, trackID string) (*models.Track, int, int) {
	album, albumIndex := GetAlbumByID(albumID)
	if album == nil {
		return nil, -1, -1
	}

	for i, track := range album.Tracks {
		if track.ID == trackID {
			return &Albums[albumIndex].Tracks[i], albumIndex, i
		}
	}
	return nil, albumIndex, -1
}
