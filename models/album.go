package models

import "time"

type Album struct {
	ID          string    `json:"id" validate:"required"`
	Title       string    `json:"title" binding:"required"`
	Artist      string    `json:"artist" binding:"required"`
	Price       float64   `json:"price" binding:"required,min=0,max=1000"`
	ReleaseDate string    `json:"release_date" validate:"required"`
	TrackCount  int       `json:"track_count"`
	Tracks      []Track   `json:"tracks"`
	Genre       string    `json:"genre" validate:"required,min=1,max=30"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Track struct {
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,min=1,max=100"`
	Duration    string `json:"duration" validate:"required,duration"`
	TrackNumber int    `json:"track_number" validate:"required,min=1"`
	AudioURL    string `json:"audio_url" validate:"omitempty,url"`
	AlbumID     string `json:"album_id"`
}
