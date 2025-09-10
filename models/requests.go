package models

type CreateAlbumRequest struct {
	ID          string  `json:"id" validate:"required,albumid"`
	Title       string  `json:"title" validate:"required,min=1,max=100"`
	Artist      string  `json:"artist" validate:"required,min=1,max=50"`
	Price       float64 `json:"price" validate:"required,min=0,max=1000"`
	ReleaseDate string  `json:"release_date" validate:"required"`
	Genre       string  `json:"genre" validate:"required,min=1,max=30"`
}

type CreateTrackRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=100"`
	Duration    string `json:"duration" validate:"required,duration"`
	TrackNumber int    `json:"track_number" validate:"required,min=1"`
	AudioURL    string `json:"audio_url" validate:"omitempty,url"`
}
